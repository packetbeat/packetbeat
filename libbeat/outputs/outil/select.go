package outil

import (
	"fmt"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/fmtstr"
	"github.com/elastic/beats/libbeat/processors"
)

type Selector struct {
	sel selector
}

type Settings struct {
	// single selector key and default option keyword
	Key string

	// multi-selector key in config
	MultiKey string

	// if enabled a selector `key` in config will be generated, if `key` is present
	EnableSingleOnly bool

	// Fail building selector if `key` and `multiKey` are missing
	FailEmpty bool
}

type selector interface {
	sel(evt common.MapStr) (string, error)
}

type emptySelector struct{}

type listSelector struct {
	selectors []selector
}

type condSelector struct {
	s    selector
	cond *processors.Condition
}

type constSelector struct {
	s string
}

type fmtSelector struct {
	f         fmtstr.EventFormatString
	otherwise string
}

type mapSelector struct {
	from      selector
	otherwise string
	to        map[string]string
}

var nilSelector selector = &emptySelector{}

// Select runs configured selector against the current event.
// If no matching selector is found, an empty string is returned.
// It's up to the caller to decide if an empty string is an error
// or an expected result.
func (s Selector) Select(evt common.MapStr) (string, error) {
	return s.sel.sel(evt)
}

func (s Selector) IsEmpty() bool {
	return s.sel == nilSelector || s.sel == nil
}

func BuildSelector(cfg *common.Config, settings Settings) (Selector, error) {
	var sel []selector

	key := settings.Key
	multiKey := settings.MultiKey
	found := false

	if cfg.HasField(multiKey) {
		found = true
		sub, err := cfg.Child(multiKey, -1)
		if err != nil {
			return Selector{}, err
		}

		var table []*common.Config
		if err := sub.Unpack(&table); err != nil {
			return Selector{}, err
		}

		for _, config := range table {
			action, err := buildSingle(config, key)
			if err != nil {
				return Selector{}, err
			}

			if action != nilSelector {
				sel = append(sel, action)
			}
		}
	}

	if settings.EnableSingleOnly && cfg.HasField(key) {
		found = true

		// expect event-format-string
		str, err := cfg.String(key, -1)
		if err != nil {
			return Selector{}, err
		}

		fmtstr, err := fmtstr.CompileEvent(str)
		if err != nil {
			return Selector{}, fmt.Errorf("%v in %v", err, cfg.PathOf(key))
		}

		sel = append(sel, &fmtSelector{f: *fmtstr})
	}

	if settings.FailEmpty && !found {
		if settings.EnableSingleOnly {
			return Selector{}, fmt.Errorf("missing required '%v' or '%v' in %v",
				key, multiKey, cfg.Path())
		}

		return Selector{}, fmt.Errorf("missing required '%v' in %v",
			multiKey, cfg.Path())
	}

	switch len(sel) {
	case 0:
		return Selector{nilSelector}, nil
	case 1:
		return Selector{sel[0]}, nil
	default:
		return Selector{&listSelector{sel}}, nil
	}
}

func buildSingle(cfg *common.Config, key string) (selector, error) {
	// TODO: check for unknown fields

	// 3. extract required key-word handler
	if !cfg.HasField(key) {
		return nil, fmt.Errorf("missing %v", cfg.PathOf(key))
	}

	str, err := cfg.String(key, -1)
	if err != nil {
		return nil, err
	}

	evtfmt, err := fmtstr.CompileEvent(str)
	if err != nil {
		return nil, fmt.Errorf("%v in %v", err, cfg.PathOf(key))
	}

	// 2. extract optional `default` value
	var otherwise string
	if cfg.HasField("default") {
		tmp, err := cfg.String("default", -1)
		if err != nil {
			return nil, err
		}
		otherwise = tmp
	}

	// 3. extract optional `mapping`
	mapping := struct {
		Table map[string]string `config:"mappings"`
	}{nil}
	if cfg.HasField("mappings") {
		if err := cfg.Unpack(&mapping); err != nil {
			return nil, err
		}
	}

	// 4. extract conditional
	var cond *processors.Condition
	if cfg.HasField("when") {
		sub, err := cfg.Child("when", -1)
		if err != nil {
			return nil, err
		}

		condConfig := processors.ConditionConfig{}
		if err := sub.Unpack(&condConfig); err != nil {
			return nil, err
		}

		tmp, err := processors.NewCondition(&condConfig)
		if err != nil {
			return nil, err
		}

		cond = tmp
	}

	// 5. build selector from available fields
	var sel selector
	if len(mapping.Table) > 0 {
		if evtfmt.IsConst() {
			str, err := evtfmt.Run(common.MapStr{})
			if err != nil {
				return nil, err
			}

			str = mapping.Table[str]
			if str == "" {
				str = otherwise
			}

			if str == "" {
				sel = nilSelector
			} else {
				sel = &constSelector{str}
			}
		} else {
			sel = &mapSelector{
				from:      &fmtSelector{f: *evtfmt},
				to:        mapping.Table,
				otherwise: otherwise,
			}
		}
	} else {
		if evtfmt.IsConst() {
			str, err := evtfmt.Run(common.MapStr{})
			if err != nil {
				return nil, err
			}
			sel = &constSelector{str}
		} else {
			sel = &fmtSelector{f: *evtfmt, otherwise: otherwise}
		}
	}
	if cond != nil && sel != nilSelector {
		sel = &condSelector{s: sel, cond: cond}
	}

	return sel, nil
}

func (s *emptySelector) sel(evt common.MapStr) (string, error) {
	return "", nil
}

func (s *listSelector) sel(evt common.MapStr) (string, error) {
	for _, sub := range s.selectors {
		n, err := sub.sel(evt)
		if err != nil { // TODO: try
			return n, err
		}

		if n != "" {
			return n, nil
		}
	}

	return "", nil
}

func (s *condSelector) sel(evt common.MapStr) (string, error) {
	if !s.cond.Check(evt) {
		return "", nil
	}
	return s.s.sel(evt)
}

func (s *constSelector) sel(_ common.MapStr) (string, error) {
	return s.s, nil
}

func (s *fmtSelector) sel(evt common.MapStr) (string, error) {
	n, err := s.f.Run(evt)
	if err != nil {
		// err will be set if not all keys present in event ->
		// return empty selector result and ignore error
		return s.otherwise, nil
	}

	if n == "" {
		return s.otherwise, nil
	}
	return n, nil
}

func (s *mapSelector) sel(evt common.MapStr) (string, error) {
	n, err := s.from.sel(evt)
	if err != nil {
		if s.otherwise == "" {
			return "", err
		}
		return s.otherwise, nil
	}

	if n == "" {
		return s.otherwise, nil
	}

	n = s.to[n]
	if n == "" {
		return s.otherwise, nil
	}
	return n, nil
}
