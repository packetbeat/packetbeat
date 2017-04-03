package match

import "regexp/syntax"

type trans func(*syntax.Regexp) (bool, *syntax.Regexp)

var transformations = []trans{
	simplify,
	uncapture,
	trimLeft,
	trimRight,
	unconcat,
	concatRepetition,
	flattenRepetition,
}

// optimize runs minimal regular expression optimizations
// until fix-point.
func optimize(r *syntax.Regexp) *syntax.Regexp {
	for {
		changed := false
		for _, t := range transformations {
			var upd bool
			upd, r = t(r)
			changed = changed || upd
		}

		if changed == false {
			return r
		}
	}
}

// Simplify regular expression by stdlib.
func simplify(r *syntax.Regexp) (bool, *syntax.Regexp) {
	return false, r.Simplify()
}

// uncapture optimizes regular expression by removing capture groups from
// regular expression potentially allocating memory when executed.
func uncapture(r *syntax.Regexp) (bool, *syntax.Regexp) {
	if r.Op == syntax.OpCapture {
		// try to uncapture
		if len(r.Sub) == 1 {
			_, sub := uncapture(r.Sub[0])
			return true, sub
		}

		tmp := *r
		tmp.Op = syntax.OpConcat
		r = &tmp
	}

	sub := make([]*syntax.Regexp, len(r.Sub))
	modified := false
	for i := range r.Sub {
		var m bool
		m, sub[i] = uncapture(r.Sub[i])
		modified = modified || m
	}

	if !modified {
		return false, r
	}

	tmp := *r
	tmp.Sub = sub
	return true, &tmp
}

// trimLeft removes not required '.*' from beginning of regular expressions.
func trimLeft(r *syntax.Regexp) (bool, *syntax.Regexp) {
	if eqPrefixAnyRegex(r, patDotStar, patNullBeginDotStar) {
		tmp := *r
		tmp.Sub = tmp.Sub[1:]
		return true, &tmp
	}

	return false, r
}

// trimRight removes not required '.*' from end of regular expressions.
func trimRight(r *syntax.Regexp) (bool, *syntax.Regexp) {
	if eqSuffixAnyRegex(r, patDotStar, patNullEndDotStar) {
		i := len(r.Sub) - 1
		tmp := *r
		tmp.Sub = tmp.Sub[0:i]
		return true, &tmp
	}

	return false, r
}

// unconcat removes intermediate regular expression concatenations generated by
// parser if concatenation contains only 1 element. Removal of object from
// parse-tree can enable other optimization to fire.
func unconcat(r *syntax.Regexp) (bool, *syntax.Regexp) {
	switch {
	case r.Op == syntax.OpConcat && len(r.Sub) <= 1:
		if len(r.Sub) == 1 {
			return true, r.Sub[0]
		}

		return true, &syntax.Regexp{
			Op:    syntax.OpEmptyMatch,
			Flags: r.Flags,
		}

	case r.Op == syntax.OpRepeat && r.Min == r.Max && r.Min == 1:
		return true, r.Sub[0]
	}

	return false, r
}

// concatRepetition concatenates 2 consecutive repeated sub-patterns into a
// repetition of length 2.
func concatRepetition(r *syntax.Regexp) (bool, *syntax.Regexp) {

	if r.Op != syntax.OpConcat {
		// don't iterate sub-expressions if top-level is no OpConcat
		return false, r
	}

	// check if concatenated op is already a repetition
	if isConcatRepetition(r) {
		return false, r
	}

	// concatenate repetitions in sub-expressions first
	var subs []*syntax.Regexp
	changed := false
	for _, sub := range r.Sub {
		changedSub, tmp := concatRepetition(sub)
		changed = changed || changedSub
		subs = append(subs, tmp)
	}

	var concat []*syntax.Regexp
	lastMerged := -1
	for i, j := 0, 1; j < len(subs); i, j = j, j+1 {
		if subs[i].Op == syntax.OpRepeat && eqRegex(subs[i].Sub[0], subs[j]) {
			r := subs[i]
			concat = append(concat,
				&syntax.Regexp{
					Op:    syntax.OpRepeat,
					Sub:   r.Sub,
					Min:   r.Min + 1,
					Max:   r.Max + 1,
					Flags: r.Flags,
				},
			)

			lastMerged = j
			changed = true
			j++
			continue
		}

		if isConcatRepetition(subs[i]) && eqRegex(subs[i].Sub[0], subs[j]) {
			r := subs[i]
			concat = append(concat,
				&syntax.Regexp{
					Op:    syntax.OpConcat,
					Sub:   append(r.Sub, r.Sub[0]),
					Flags: r.Flags,
				},
			)

			lastMerged = j
			changed = true
			j++
			continue
		}

		if eqRegex(subs[i], subs[j]) {
			r := subs[i]
			concat = append(concat,
				&syntax.Regexp{
					Op:    syntax.OpRepeat,
					Sub:   []*syntax.Regexp{r},
					Min:   2,
					Max:   2,
					Flags: r.Flags,
				},
			)

			lastMerged = j
			changed = true
			j++
			continue
		}

		concat = append(concat, subs[i])
	}

	if lastMerged+1 != len(subs) {
		concat = append(concat, subs[len(subs)-1])
	}

	r = &syntax.Regexp{
		Op:    syntax.OpConcat,
		Sub:   concat,
		Flags: r.Flags,
	}
	return changed, r
}

// flattenRepetition flattens nested repetitions
func flattenRepetition(r *syntax.Regexp) (bool, *syntax.Regexp) {
	if r.Op != syntax.OpConcat {
		// don't iterate sub-expressions if top-level is no OpConcat
		return false, r
	}

	sub := r.Sub
	inRepetition := false
	if isConcatRepetition(r) {
		sub = sub[:1]
		inRepetition = true

		// create flattened regex repetition mulitplying count
		// if nexted expression is also a repetition
		if s := sub[0]; isConcatRepetition(s) {
			count := len(s.Sub) * len(r.Sub)
			return true, &syntax.Regexp{
				Op:    syntax.OpRepeat,
				Sub:   s.Sub[:1],
				Min:   count,
				Max:   count,
				Flags: r.Flags | s.Flags,
			}
		}
	}

	// recursively check if we can flatten sub-expressions
	changed := false
	for i, s := range sub {
		upd, tmp := flattenRepetition(s)
		changed = changed || upd
		sub[i] = tmp
	}

	if !changed {
		return false, r
	}

	// fix up top-level repetition with modified one
	tmp := *r
	if inRepetition {
		for i := range r.Sub {
			tmp.Sub[i] = sub[0]
		}
	} else {
		tmp.Sub = sub
	}
	return changed, &tmp
}
