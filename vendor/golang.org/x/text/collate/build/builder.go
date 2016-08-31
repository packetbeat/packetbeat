// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package build // import "golang.org/x/text/collate/build"

import (
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/collate/colltab"
	"golang.org/x/text/language"
	"golang.org/x/text/unicode/norm"
)

// TODO: optimizations:
// - expandElem is currently 20K. By putting unique colElems in a separate
//   table and having a byte array of indexes into this table, we can reduce
//   the total size to about 7K. By also factoring out the length bytes, we
//   can reduce this to about 6K.
// - trie valueBlocks are currently 100K. There are a lot of sparse blocks
//   and many consecutive values with the same stride. This can be further
//   compacted.
// - Compress secondary weights into 8 bits.
// - Some LDML specs specify a context element. Currently we simply concatenate
//   those.  Context can be implemented using the contraction trie. If Builder
//   could analyze and detect when using a context makes sense, there is no
//   need to expose this construct in the API.

// A Builder builds a root collation table.  The user must specify the
// collation elements for each entry.  A common use will be to base the weights
// on those specified in the allkeys* file as provided by the UCA or CLDR.
type Builder struct {
	index  *trieBuilder
	root   ordering
	locale []*Tailoring
	t      *table
	err    error
	built  bool

	minNonVar int // lowest primary recorded for a variable
	varTop    int // highest primary recorded for a non-variable

	// indexes used for reusing expansions and contractions
	expIndex map[string]int      // positions of expansions keyed by their string representation
	ctHandle map[string]ctHandle // contraction handles keyed by a concatenation of the suffixes
	ctElem   map[string]int      // contraction elements keyed by their string representation
}

// A Tailoring builds a collation table based on another collation table.
// The table is defined by specifying tailorings to the underlying table.
// See http://unicode.org/reports/tr35/ for an overview of tailoring
// collation tables.  The CLDR contains pre-defined tailorings for a variety
// of languages (See http://www.unicode.org/Public/cldr/<version>/core.zip.)
type Tailoring struct {
	id      string
	builder *Builder
	index   *ordering

	anchor *entry
	before bool
}

// NewBuilder returns a new Builder.
func NewBuilder() *Builder {
	return &Builder{
		index:    newTrieBuilder(),
		root:     makeRootOrdering(),
		expIndex: make(map[string]int),
		ctHandle: make(map[string]ctHandle),
		ctElem:   make(map[string]int),
	}
}

// Tailoring returns a Tailoring for the given locale.  One should
// have completed all calls to Add before calling Tailoring.
func (b *Builder) Tailoring(loc language.Tag) *Tailoring {
	t := &Tailoring{
		id:      loc.String(),
		builder: b,
		index:   b.root.clone(),
	}
	t.index.id = t.id
	b.locale = append(b.locale, t)
	return t
}

// Add adds an entry to the collation element table, mapping
// a slice of runes to a sequence of collation elements.
// A collation element is specified as list of weights: []int{primary, secondary, ...}.
// The entries are typically obtained from a collation element table
// as defined in http://www.unicode.org/reports/tr10/#Data_Table_Format.
// Note that the collation elements specified by colelems are only used
// as a guide.  The actual weights generated by Builder may differ.
// The argument variables is a list of indices into colelems that should contain
// a value for each colelem that is a variable. (See the reference above.)
func (b *Builder) Add(runes []rune, colelems [][]int, variables []int) error {
	str := string(runes)
	elems := make([]rawCE, len(colelems))
	for i, ce := range colelems {
		if len(ce) == 0 {
			break
		}
		elems[i] = makeRawCE(ce, 0)
		if len(ce) == 1 {
			elems[i].w[1] = defaultSecondary
		}
		if len(ce) <= 2 {
			elems[i].w[2] = defaultTertiary
		}
		if len(ce) <= 3 {
			elems[i].w[3] = ce[0]
		}
	}
	for i, ce := range elems {
		p := ce.w[0]
		isvar := false
		for _, j := range variables {
			if i == j {
				isvar = true
			}
		}
		if isvar {
			if p >= b.minNonVar && b.minNonVar > 0 {
				return fmt.Errorf("primary value %X of variable is larger than the smallest non-variable %X", p, b.minNonVar)
			}
			if p > b.varTop {
				b.varTop = p
			}
		} else if p > 1 { // 1 is a special primary value reserved for FFFE
			if p <= b.varTop {
				return fmt.Errorf("primary value %X of non-variable is smaller than the highest variable %X", p, b.varTop)
			}
			if b.minNonVar == 0 || p < b.minNonVar {
				b.minNonVar = p
			}
		}
	}
	elems, err := convertLargeWeights(elems)
	if err != nil {
		return err
	}
	cccs := []uint8{}
	nfd := norm.NFD.String(str)
	for i := range nfd {
		cccs = append(cccs, norm.NFD.PropertiesString(nfd[i:]).CCC())
	}
	if len(cccs) < len(elems) {
		if len(cccs) > 2 {
			return fmt.Errorf("number of decomposed characters should be greater or equal to the number of collation elements for len(colelems) > 3 (%d < %d)", len(cccs), len(elems))
		}
		p := len(elems) - 1
		for ; p > 0 && elems[p].w[0] == 0; p-- {
			elems[p].ccc = cccs[len(cccs)-1]
		}
		for ; p >= 0; p-- {
			elems[p].ccc = cccs[0]
		}
	} else {
		for i := range elems {
			elems[i].ccc = cccs[i]
		}
	}
	// doNorm in collate.go assumes that the following conditions hold.
	if len(elems) > 1 && len(cccs) > 1 && cccs[0] != 0 && cccs[0] != cccs[len(cccs)-1] {
		return fmt.Errorf("incompatible CCC values for expansion %X (%d)", runes, cccs)
	}
	b.root.newEntry(str, elems)
	return nil
}

func (t *Tailoring) setAnchor(anchor string) error {
	anchor = norm.NFC.String(anchor)
	a := t.index.find(anchor)
	if a == nil {
		a = t.index.newEntry(anchor, nil)
		a.implicit = true
		a.modified = true
		for _, r := range []rune(anchor) {
			e := t.index.find(string(r))
			e.lock = true
		}
	}
	t.anchor = a
	return nil
}

// SetAnchor sets the point after which elements passed in subsequent calls to
// Insert will be inserted.  It is equivalent to the reset directive in an LDML
// specification.  See Insert for an example.
// SetAnchor supports the following logical reset positions:
// <first_tertiary_ignorable/>, <last_teriary_ignorable/>, <first_primary_ignorable/>,
// and <last_non_ignorable/>.
func (t *Tailoring) SetAnchor(anchor string) error {
	if err := t.setAnchor(anchor); err != nil {
		return err
	}
	t.before = false
	return nil
}

// SetAnchorBefore is similar to SetAnchor, except that subsequent calls to
// Insert will insert entries before the anchor.
func (t *Tailoring) SetAnchorBefore(anchor string) error {
	if err := t.setAnchor(anchor); err != nil {
		return err
	}
	t.before = true
	return nil
}

// Insert sets the ordering of str relative to the entry set by the previous
// call to SetAnchor or Insert.  The argument extend corresponds
// to the extend elements as defined in LDML.  A non-empty value for extend
// will cause the collation elements corresponding to extend to be appended
// to the collation elements generated for the entry added by Insert.
// This has the same net effect as sorting str after the string anchor+extend.
// See http://www.unicode.org/reports/tr10/#Tailoring_Example for details
// on parametric tailoring and http://unicode.org/reports/tr35/#Collation_Elements
// for full details on LDML.
//
// Examples: create a tailoring for Swedish, where "ä" is ordered after "z"
// at the primary sorting level:
//      t := b.Tailoring("se")
// 		t.SetAnchor("z")
// 		t.Insert(colltab.Primary, "ä", "")
// Order "ü" after "ue" at the secondary sorting level:
//		t.SetAnchor("ue")
//		t.Insert(colltab.Secondary, "ü","")
// or
//		t.SetAnchor("u")
//		t.Insert(colltab.Secondary, "ü", "e")
// Order "q" afer "ab" at the secondary level and "Q" after "q"
// at the tertiary level:
// 		t.SetAnchor("ab")
// 		t.Insert(colltab.Secondary, "q", "")
// 		t.Insert(colltab.Tertiary, "Q", "")
// Order "b" before "a":
//      t.SetAnchorBefore("a")
//      t.Insert(colltab.Primary, "b", "")
// Order "0" after the last primary ignorable:
//      t.SetAnchor("<last_primary_ignorable/>")
//      t.Insert(colltab.Primary, "0", "")
func (t *Tailoring) Insert(level colltab.Level, str, extend string) error {
	if t.anchor == nil {
		return fmt.Errorf("%s:Insert: no anchor point set for tailoring of %s", t.id, str)
	}
	str = norm.NFC.String(str)
	e := t.index.find(str)
	if e == nil {
		e = t.index.newEntry(str, nil)
	} else if e.logical != noAnchor {
		return fmt.Errorf("%s:Insert: cannot reinsert logical reset position %q", t.id, e.str)
	}
	if e.lock {
		return fmt.Errorf("%s:Insert: cannot reinsert element %q", t.id, e.str)
	}
	a := t.anchor
	// Find the first element after the anchor which differs at a level smaller or
	// equal to the given level.  Then insert at this position.
	// See http://unicode.org/reports/tr35/#Collation_Elements, Section 5.14.5 for details.
	e.before = t.before
	if t.before {
		t.before = false
		if a.prev == nil {
			a.insertBefore(e)
		} else {
			for a = a.prev; a.level > level; a = a.prev {
			}
			a.insertAfter(e)
		}
		e.level = level
	} else {
		for ; a.level > level; a = a.next {
		}
		e.level = a.level
		if a != e {
			a.insertAfter(e)
			a.level = level
		} else {
			// We don't set a to prev itself. This has the effect of the entry
			// getting new collation elements that are an increment of itself.
			// This is intentional.
			a.prev.level = level
		}
	}
	e.extend = norm.NFD.String(extend)
	e.exclude = false
	e.modified = true
	e.elems = nil
	t.anchor = e
	return nil
}

func (o *ordering) getWeight(e *entry) []rawCE {
	if len(e.elems) == 0 && e.logical == noAnchor {
		if e.implicit {
			for _, r := range e.runes {
				e.elems = append(e.elems, o.getWeight(o.find(string(r)))...)
			}
		} else if e.before {
			count := [colltab.Identity + 1]int{}
			a := e
			for ; a.elems == nil && !a.implicit; a = a.next {
				count[a.level]++
			}
			e.elems = []rawCE{makeRawCE(a.elems[0].w, a.elems[0].ccc)}
			for i := colltab.Primary; i < colltab.Quaternary; i++ {
				if count[i] != 0 {
					e.elems[0].w[i] -= count[i]
					break
				}
			}
			if e.prev != nil {
				o.verifyWeights(e.prev, e, e.prev.level)
			}
		} else {
			prev := e.prev
			e.elems = nextWeight(prev.level, o.getWeight(prev))
			o.verifyWeights(e, e.next, e.level)
		}
	}
	return e.elems
}

func (o *ordering) addExtension(e *entry) {
	if ex := o.find(e.extend); ex != nil {
		e.elems = append(e.elems, ex.elems...)
	} else {
		for _, r := range []rune(e.extend) {
			e.elems = append(e.elems, o.find(string(r)).elems...)
		}
	}
	e.extend = ""
}

func (o *ordering) verifyWeights(a, b *entry, level colltab.Level) error {
	if level == colltab.Identity || b == nil || b.elems == nil || a.elems == nil {
		return nil
	}
	for i := colltab.Primary; i < level; i++ {
		if a.elems[0].w[i] < b.elems[0].w[i] {
			return nil
		}
	}
	if a.elems[0].w[level] >= b.elems[0].w[level] {
		err := fmt.Errorf("%s:overflow: collation elements of %q (%X) overflows those of %q (%X) at level %d (%X >= %X)", o.id, a.str, a.runes, b.str, b.runes, level, a.elems, b.elems)
		log.Println(err)
		// TODO: return the error instead, or better, fix the conflicting entry by making room.
	}
	return nil
}

func (b *Builder) error(e error) {
	if e != nil {
		b.err = e
	}
}

func (b *Builder) errorID(locale string, e error) {
	if e != nil {
		b.err = fmt.Errorf("%s:%v", locale, e)
	}
}

// patchNorm ensures that NFC and NFD counterparts are consistent.
func (o *ordering) patchNorm() {
	// Insert the NFD counterparts, if necessary.
	for _, e := range o.ordered {
		nfd := norm.NFD.String(e.str)
		if nfd != e.str {
			if e0 := o.find(nfd); e0 != nil && !e0.modified {
				e0.elems = e.elems
			} else if e.modified && !equalCEArrays(o.genColElems(nfd), e.elems) {
				e := o.newEntry(nfd, e.elems)
				e.modified = true
			}
		}
	}
	// Update unchanged composed forms if one of their parts changed.
	for _, e := range o.ordered {
		nfd := norm.NFD.String(e.str)
		if e.modified || nfd == e.str {
			continue
		}
		if e0 := o.find(nfd); e0 != nil {
			e.elems = e0.elems
		} else {
			e.elems = o.genColElems(nfd)
			if norm.NFD.LastBoundary([]byte(nfd)) == 0 {
				r := []rune(nfd)
				head := string(r[0])
				tail := ""
				for i := 1; i < len(r); i++ {
					s := norm.NFC.String(head + string(r[i]))
					if e0 := o.find(s); e0 != nil && e0.modified {
						head = s
					} else {
						tail += string(r[i])
					}
				}
				e.elems = append(o.genColElems(head), o.genColElems(tail)...)
			}
		}
	}
	// Exclude entries for which the individual runes generate the same collation elements.
	for _, e := range o.ordered {
		if len(e.runes) > 1 && equalCEArrays(o.genColElems(e.str), e.elems) {
			e.exclude = true
		}
	}
}

func (b *Builder) buildOrdering(o *ordering) {
	for _, e := range o.ordered {
		o.getWeight(e)
	}
	for _, e := range o.ordered {
		o.addExtension(e)
	}
	o.patchNorm()
	o.sort()
	simplify(o)
	b.processExpansions(o)   // requires simplify
	b.processContractions(o) // requires simplify

	t := newNode()
	for e := o.front(); e != nil; e, _ = e.nextIndexed() {
		if !e.skip() {
			ce, err := e.encode()
			b.errorID(o.id, err)
			t.insert(e.runes[0], ce)
		}
	}
	o.handle = b.index.addTrie(t)
}

func (b *Builder) build() (*table, error) {
	if b.built {
		return b.t, b.err
	}
	b.built = true
	b.t = &table{
		maxContractLen: utf8.UTFMax,
		variableTop:    uint32(b.varTop),
	}

	b.buildOrdering(&b.root)
	b.t.root = b.root.handle
	for _, t := range b.locale {
		b.buildOrdering(t.index)
		if b.err != nil {
			break
		}
	}
	i, err := b.index.generate()
	b.t.index = *i
	b.error(err)
	return b.t, b.err
}

// Build builds the root Collator.
// TODO: return Weighter instead
func (b *Builder) Build() (colltab.Weighter, error) {
	t, err := b.build()
	if err != nil {
		return nil, err
	}
	table := colltab.Init(t)
	if table == nil {
		panic("generated table of incompatible type")
	}
	return table, nil
}

// Build builds a Collator for Tailoring t.
func (t *Tailoring) Build() (colltab.Weighter, error) {
	// TODO: implement.
	return nil, nil
}

// Print prints the tables for b and all its Tailorings as a Go file
// that can be included in the Collate package.
func (b *Builder) Print(w io.Writer) (n int, err error) {
	p := func(nn int, e error) {
		n += nn
		if err == nil {
			err = e
		}
	}
	t, err := b.build()
	if err != nil {
		return 0, err
	}
	p(fmt.Fprintf(w, `var availableLocales = "und`))
	for _, loc := range b.locale {
		if loc.id != "und" {
			p(fmt.Fprintf(w, ",%s", loc.id))
		}
	}
	p(fmt.Fprint(w, "\"\n\n"))
	p(fmt.Fprintf(w, "const varTop = 0x%x\n\n", b.varTop))
	p(fmt.Fprintln(w, "var locales = [...]tableIndex{"))
	for _, loc := range b.locale {
		if loc.id == "und" {
			p(t.fprintIndex(w, loc.index.handle, loc.id))
		}
	}
	for _, loc := range b.locale {
		if loc.id != "und" {
			p(t.fprintIndex(w, loc.index.handle, loc.id))
		}
	}
	p(fmt.Fprint(w, "}\n\n"))
	n, _, err = t.fprint(w, "main")
	return
}

// reproducibleFromNFKD checks whether the given expansion could be generated
// from an NFKD expansion.
func reproducibleFromNFKD(e *entry, exp, nfkd []rawCE) bool {
	// Length must be equal.
	if len(exp) != len(nfkd) {
		return false
	}
	for i, ce := range exp {
		// Primary and secondary values should be equal.
		if ce.w[0] != nfkd[i].w[0] || ce.w[1] != nfkd[i].w[1] {
			return false
		}
		// Tertiary values should be equal to maxTertiary for third element onwards.
		// TODO: there seem to be a lot of cases in CLDR (e.g. ㏭ in zh.xml) that can
		// simply be dropped.  Try this out by dropping the following code.
		if i >= 2 && ce.w[2] != maxTertiary {
			return false
		}
		if _, err := makeCE(ce); err != nil {
			// Simply return false. The error will be caught elsewhere.
			return false
		}
	}
	return true
}

func simplify(o *ordering) {
	// Runes that are a starter of a contraction should not be removed.
	// (To date, there is only Kannada character 0CCA.)
	keep := make(map[rune]bool)
	for e := o.front(); e != nil; e, _ = e.nextIndexed() {
		if len(e.runes) > 1 {
			keep[e.runes[0]] = true
		}
	}
	// Tag entries for which the runes NFKD decompose to identical values.
	for e := o.front(); e != nil; e, _ = e.nextIndexed() {
		s := e.str
		nfkd := norm.NFKD.String(s)
		nfd := norm.NFD.String(s)
		if e.decompose || len(e.runes) > 1 || len(e.elems) == 1 || keep[e.runes[0]] || nfkd == nfd {
			continue
		}
		if reproducibleFromNFKD(e, e.elems, o.genColElems(nfkd)) {
			e.decompose = true
		}
	}
}

// appendExpansion converts the given collation sequence to
// collation elements and adds them to the expansion table.
// It returns an index to the expansion table.
func (b *Builder) appendExpansion(e *entry) int {
	t := b.t
	i := len(t.expandElem)
	ce := uint32(len(e.elems))
	t.expandElem = append(t.expandElem, ce)
	for _, w := range e.elems {
		ce, err := makeCE(w)
		if err != nil {
			b.error(err)
			return -1
		}
		t.expandElem = append(t.expandElem, ce)
	}
	return i
}

// processExpansions extracts data necessary to generate
// the extraction tables.
func (b *Builder) processExpansions(o *ordering) {
	for e := o.front(); e != nil; e, _ = e.nextIndexed() {
		if !e.expansion() {
			continue
		}
		key := fmt.Sprintf("%v", e.elems)
		i, ok := b.expIndex[key]
		if !ok {
			i = b.appendExpansion(e)
			b.expIndex[key] = i
		}
		e.expansionIndex = i
	}
}

func (b *Builder) processContractions(o *ordering) {
	// Collate contractions per starter rune.
	starters := []rune{}
	cm := make(map[rune][]*entry)
	for e := o.front(); e != nil; e, _ = e.nextIndexed() {
		if e.contraction() {
			if len(e.str) > b.t.maxContractLen {
				b.t.maxContractLen = len(e.str)
			}
			r := e.runes[0]
			if _, ok := cm[r]; !ok {
				starters = append(starters, r)
			}
			cm[r] = append(cm[r], e)
		}
	}
	// Add entries of single runes that are at a start of a contraction.
	for e := o.front(); e != nil; e, _ = e.nextIndexed() {
		if !e.contraction() {
			r := e.runes[0]
			if _, ok := cm[r]; ok {
				cm[r] = append(cm[r], e)
			}
		}
	}
	// Build the tries for the contractions.
	t := b.t
	for _, r := range starters {
		l := cm[r]
		// Compute suffix strings. There are 31 different contraction suffix
		// sets for 715 contractions and 82 contraction starter runes as of
		// version 6.0.0.
		sufx := []string{}
		hasSingle := false
		for _, e := range l {
			if len(e.runes) > 1 {
				sufx = append(sufx, string(e.runes[1:]))
			} else {
				hasSingle = true
			}
		}
		if !hasSingle {
			b.error(fmt.Errorf("no single entry for starter rune %U found", r))
			continue
		}
		// Unique the suffix set.
		sort.Strings(sufx)
		key := strings.Join(sufx, "\n")
		handle, ok := b.ctHandle[key]
		if !ok {
			var err error
			handle, err = t.contractTries.appendTrie(sufx)
			if err != nil {
				b.error(err)
			}
			b.ctHandle[key] = handle
		}
		// Bucket sort entries in index order.
		es := make([]*entry, len(l))
		for _, e := range l {
			var p, sn int
			if len(e.runes) > 1 {
				str := []byte(string(e.runes[1:]))
				p, sn = t.contractTries.lookup(handle, str)
				if sn != len(str) {
					log.Fatalf("%s: processContractions: unexpected length for '%X'; len=%d; want %d", o.id, e.runes, sn, len(str))
				}
			}
			if es[p] != nil {
				log.Fatalf("%s: multiple contractions for position %d for rune %U", o.id, p, e.runes[0])
			}
			es[p] = e
		}
		// Create collation elements for contractions.
		elems := []uint32{}
		for _, e := range es {
			ce, err := e.encodeBase()
			b.errorID(o.id, err)
			elems = append(elems, ce)
		}
		key = fmt.Sprintf("%v", elems)
		i, ok := b.ctElem[key]
		if !ok {
			i = len(t.contractElem)
			b.ctElem[key] = i
			t.contractElem = append(t.contractElem, elems...)
		}
		// Store info in entry for starter rune.
		es[0].contractionIndex = i
		es[0].contractionHandle = handle
	}
}
