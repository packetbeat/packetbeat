package monitoring

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

// Registry to store variables and sub-registries.
// When adding or retrieving variables, all names are split on the `.`-symbol and
// intermediate registries will be generated.
type Registry struct {
	mu sync.RWMutex

	name    string
	entries map[string]entry

	opts *options
}

type entry struct {
	Var
	Mode
}

// Var interface required for every metric to implement.
type Var interface {
	Visit(Visitor) error
}

// NewRegistry create a new empty unregistered registry
func NewRegistry(opts ...Option) *Registry {
	return &Registry{
		opts:    applyOpts(nil, opts),
		entries: map[string]entry{},
	}
}

func (r *Registry) Do(mode Mode, f func(string, interface{}) error) error {
	return r.doVisit(mode, NewKeyValueVisitor(f))
}

func (r *Registry) VisitMode(mode Mode, vs Visitor) error {
	return r.doVisit(mode, vs)
}

// Visit uses the Visitor interface to iterate the complete metrics hieararchie.
// In case of the visitor reporting an error, Visit will return immediately,
// reporting the very same error.
func (r *Registry) Visit(vs Visitor) error {
	return r.doVisit(Full, vs)
}

func (r *Registry) doVisit(mode Mode, vs Visitor) error {
	if err := vs.OnRegistryStart(); err != nil {
		return err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	first := true
	for key, v := range r.entries {
		var err error

		if v.Mode > mode {
			continue
		}

		if first {
			first = false
		} else {
			if err = vs.OnKeyNext(); err != nil {
				return err
			}
		}

		if err = vs.OnKey(key); err != nil {
			return err
		}

		if reg, ok := v.Var.(*Registry); ok {
			err = reg.doVisit(mode, vs)
		} else {
			err = v.Var.Visit(vs)
		}
		if err != nil {
			return err
		}
	}

	return vs.OnRegistryFinished()
}

// NewRegistry creates and register a new registry
func (r *Registry) NewRegistry(name string, opts ...Option) *Registry {
	v := &Registry{
		name:    fullName(r, name),
		opts:    applyOpts(r.opts, opts),
		entries: map[string]entry{},
	}
	r.Add(name, v, v.opts.mode)
	return v
}

// Get tries to find a registered variable by name.
func (r *Registry) Get(name string) Var {
	v, err := r.find(name)
	if err != nil {
		return nil
	}
	return v.Var
}

// GetRegistry tries to find a sub-registry by name.
func (r *Registry) GetRegistry(name string) *Registry {
	e, err := r.find(name)
	if err != nil {
		return nil
	}

	v := e.Var
	if v == nil {
		return nil
	}

	reg, ok := v.(*Registry)
	if !ok {
		return nil
	}

	return reg
}

// Remove removes a variable or a sub-registry by name
func (r *Registry) Remove(name string) {
	r.removeNames(strings.Split(name, "."))
}

// Clear removes all entries from the current registry
func (r *Registry) Clear() error {
	r.mu.Lock()
	r.mu.Unlock()

	if r.opts.publishExpvar {
		return errors.New("Can not clear registry with metrics being exported via expvar")
	}

	r.entries = map[string]entry{}
	return nil
}

// Add adds a new variable to the registry. The method panics if the variables
// name is already in use.
func (r *Registry) Add(name string, v Var, m Mode) {
	panicErr(r.addNames(strings.Split(name, "."), v, m))
}

func (r *Registry) addNames(names []string, v Var, m Mode) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	name := names[0]
	if len(names) == 1 {
		if _, found := r.entries[name]; found {
			return fmt.Errorf("name %v already used", name)
		}

		r.entries[name] = entry{v, m}
		return nil
	}

	if tmp, found := r.entries[name]; found {
		reg, ok := tmp.Var.(*Registry)
		if !ok {
			return fmt.Errorf("name %v already used", name)
		}

		return reg.addNames(names[1:], v, m)
	}

	sub := NewRegistry()
	sub.opts = r.opts
	if err := sub.addNames(names[1:], v, m); err != nil {
		return err
	}

	r.entries[name] = entry{sub, sub.opts.mode}
	return nil
}

func (r *Registry) find(name string) (entry, error) {
	return r.findNames(strings.Split(name, "."))
}

func (r *Registry) findNames(names []string) (entry, error) {
	switch len(names) {
	case 0:
		return entry{r, r.opts.mode}, nil
	case 1:
		r.mu.RLock()
		defer r.mu.RUnlock()
		return r.entries[names[0]], nil
	}

	r.mu.RLock()
	next, exist := r.entries[names[0]]
	r.mu.RUnlock()

	if !exist {
		return entry{}, errNotFound
	}

	if reg, ok := next.Var.(*Registry); ok {
		return reg.findNames(names[1:])
	}
	return entry{}, errInvalidName
}

func (r *Registry) removeNames(names []string) {
	switch len(names) {
	case 0:
		return
	case 1:
		r.mu.Lock()
		defer r.mu.Unlock()
		delete(r.entries, names[0])
		return
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	next, exists := r.entries[names[0]]

	// if name does not exist => don't remove anything
	if !exists {
		return
	}

	sub, ok := next.Var.(*Registry)
	if ok {
		sub.removeNames(names[1:])
		sub.mu.RLock()
		sub.mu.RUnlock()

		if len(sub.entries) == 0 {
			delete(r.entries, names[0])
		}
	}
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
