package common

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Event metadata constants. These keys are used within libbeat to identify
// metadata stored in an event.
const (
	FieldsKey = "fields"
	TagsKey   = "tags"
)

var (
	// ErrKeyNotFound indicates that the specified key was not found.
	ErrKeyNotFound = errors.New("key not found")
)

// EventMetadata contains fields and tags that can be added to an event via
// configuration.
type EventMetadata struct {
	Fields          MapStr
	FieldsUnderRoot bool `config:"fields_under_root"`
	Tags            []string
}

// MapStr is a map[string]interface{} wrapper with utility methods for common
// map operations like converting to JSON.
type MapStr map[string]interface{}

// Update copies all the key-value pairs from d to this map. If the key
// already exists then it is overwritten. This method does not merge nested
// maps.
func (m MapStr) Update(d MapStr) {
	for k, v := range d {
		m[k] = v
	}
}

// DeepUpdate recursively copies the key-value pairs from d to this map.
// If the key is present and a map as well, the sub-map will be updated recursively
// via DeepUpdate.
func (m MapStr) DeepUpdate(d MapStr) {
	for k, v := range d {
		switch val := v.(type) {
		case map[string]interface{}:
			m[k] = deepUpdateValue(m[k], MapStr(val))
		case MapStr:
			m[k] = deepUpdateValue(m[k], val)
		default:
			m[k] = v
		}
	}
}

func deepUpdateValue(old interface{}, val MapStr) interface{} {
	if old == nil {
		return val
	}

	switch sub := old.(type) {
	case MapStr:
		sub.DeepUpdate(val)
		return sub
	case map[string]interface{}:
		tmp := MapStr(sub)
		tmp.DeepUpdate(val)
		return tmp
	default:
		return val
	}
}

// Delete deletes the given key from the map.
func (m MapStr) Delete(key string) error {
	_, err := walkMap(key, m, opDelete)
	return err
}

// CopyFieldsTo copies the field specified by key to the given map. It will
// overwrite the key if it exists. An error is returned if the key does not
// exist in the source map.
func (m MapStr) CopyFieldsTo(to MapStr, key string) error {
	v, err := walkMap(key, m, opGet)
	if err != nil {
		return err
	}

	_, err = walkMap(key, to, mapStrOperation{putOperation{v}, true})
	return err
}

// Clone returns a copy of the MapStr. It recursively makes copies of inner
// maps.
func (m MapStr) Clone() MapStr {
	result := MapStr{}

	for k, v := range m {
		if innerMap, ok := tryToMapStr(v); ok {
			v = innerMap.Clone()
		}
		result[k] = v
	}

	return result
}

// HasKey returns true if the key exist. If an error occurs then false is
// returned with a non-nil error.
func (m MapStr) HasKey(key string) (bool, error) {
	hasKey, err := walkMap(key, m, opHasKey)
	if err != nil {
		return false, err
	}

	return hasKey.(bool), nil
}

// GetValue gets a value from the map. If the key does not exist then an error
// is returned.
func (m MapStr) GetValue(key string) (interface{}, error) {
	return walkMap(key, m, opGet)
}

func (m MapStr) GetValueNoErr(key string) interface{} {
	o, err := m.GetValue(key)
	if err != nil {
		return nil
	}
	return o
}

func (m MapStr) GetStringValueWithDefault(key string, df string) string {
	vi, err := m.GetValue(key)
	if err != nil || vi == nil {
		return df
	}

	v, ok := vi.(string)
	if !ok {
		return df
	} else {
		return v
	}
}

func (m MapStr) GetStringValue(key string) string {
	return m.GetStringValueWithDefault(key, "")
}

// Get value from path, and convert the value to int64
// string and intX uintX will be converted
// any failed will cause df return
func (m MapStr) GetInt64ValueWithDefault(key string, df int64) int64 {
	vi, err := m.GetValue(key)
	if err != nil || vi == nil {
		return df
	}

	switch i := vi.(type) {
	case int, int8, int16, int32, int64:
		return int64(reflect.ValueOf(i).Int())
	case uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(i).Uint())
	case float32, float64:
		return int64(reflect.ValueOf(i).Float())
	case string:
		v, err := strconv.Atoi(i)
		if err != nil {
			return df
		}
		return int64(v)
	default:
		return df
	}

}

// Get value from path, and convert the value to int64
// string and intX uintX will be converted
// any failed will cause 0 return
func (m MapStr) GetInt64Value(key string) int64 {
	return m.GetInt64ValueWithDefault(key, 0)
}

// Get value from path, and convert the value to float64
// string and intX uintX will be converted
// any failed will cause df return
func (m MapStr) GetFloat64ValueWithDefault(key string, df float64) float64 {
	vi, err := m.GetValue(key)
	if err != nil || vi == nil {
		return df
	}

	switch i := vi.(type) {
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(i).Float()
	case uint8, uint16, uint32, uint64:
		return reflect.ValueOf(i).Float()
	case float32, float64:
		return reflect.ValueOf(i).Float()
	case string:
		v, err := strconv.ParseFloat(i, 64)
		if err != nil {
			return df
		}
		return float64(v)
	default:
		return df
	}

}

// Get value from path, and convert the value to float64
// string and intX uintX will be converted
// any failed will cause 0 return
func (m MapStr) GetFloat64Value(key string) float64 {
	return m.GetFloat64ValueWithDefault(key, 0.0)
}
func (m MapStr) GetObject(key string) MapStr {
	return m.GetObjectWithDefault(key, MapStr{})
}

func (m MapStr) GetObjectWithDefault(key string, df MapStr) MapStr {
	vi, err := m.GetValue(key)
	if err != nil || vi == nil {
		return df
	}
	obj, err := toMapStr(vi)
	if err != nil || obj == nil {
		return df
	}
	return obj
}

func (m MapStr) GetObjectArray(key string) []MapStr {
	vi, err := m.GetValue(key)
	if err != nil || vi == nil {
		return nil
	}

	switch array := vi.(type) {
	case []interface{}:
		rt := make([]MapStr, 0, len(array))
		for _, item := range array {
			mItem, err := toMapStr(item)
			if err != nil {
				continue
			}

			rt = append(rt, mItem)
		}

		return rt
	case []MapStr:
		rt := make([]MapStr, 0, len(array))
		for _, item := range array {
			mItem, err := toMapStr(item)
			if err != nil {
				continue
			}

			rt = append(rt, mItem)
		}

		return rt
	}

	return nil
}

func (m MapStr) GetStringArray(key string) []string {
	vi, err := m.GetValue(key)
	if err != nil || vi == nil {
		return nil
	}

	sa := make([]string, 0)
	switch array := vi.(type) {
	case []string:
		for _, f := range array {
			sa = append(sa, f)
		}
	case []interface{}:
		for _, f := range array {
			sa = append(sa, fmt.Sprintf("%v", f))
		}
	default:
		return nil
	}

	return sa

}

// Put associates the specified value with the specified key. If the map
// previously contained a mapping for the key, the old value is replaced and
// returned. The key can be expressed in dot-notation (e.g. x.y) to put a value
// into a nested map.
//
// If you need insert keys containing dots then you must use bracket notation
// to insert values (e.g. m[key] = value).
func (m MapStr) Put(key string, value interface{}) (interface{}, error) {
	return walkMap(key, m, mapStrOperation{putOperation{value}, true})
}

// StringToPrint returns the MapStr as pretty JSON.
func (m MapStr) StringToPrint() string {
	json, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return fmt.Sprintf("Not valid json: %v", err)
	}
	return string(json)
}

// String returns the MapStr as JSON.
func (m MapStr) String() string {
	bytes, err := json.Marshal(m)
	if err != nil {
		return fmt.Sprintf("Not valid json: %v", err)
	}
	return string(bytes)
}

func MapStrUnmarshal(data []byte) (MapStr, error) {
	var ms interface{}
	err := json.Unmarshal(data, &ms)
	if err != nil {
		return nil, err
	}

	return toMapStr(ms)
}

func (m MapStr) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, errors.Wrap(err, "Not valid json")
	}
	return bytes, nil
}

// Flatten flattens the given MapStr and returns a flat MapStr.
//
// Example:
//   "hello": MapStr{"world": "test" }
//
// This is converted to:
//   "hello.world": "test"
//
// This can be useful for testing or logging.
func (m MapStr) Flatten() MapStr {
	return flatten("", m, MapStr{})
}

// flatten is a helper for Flatten. See docs for Flatten. For convenience the
// out parameter is returned.
func flatten(prefix string, in, out MapStr) MapStr {
	for k, v := range in {
		var fullKey string
		if prefix == "" {
			fullKey = k
		} else {
			fullKey = fmt.Sprintf("%s.%s", prefix, k)
		}

		if m, ok := tryToMapStr(v); ok {
			flatten(fullKey, m, out)
		} else {
			out[fullKey] = v
		}
	}
	return out
}

// MapStrUnion creates a new MapStr containing the union of the
// key-value pairs of the two maps. If the same key is present in
// both, the key-value pairs from dict2 overwrite the ones from dict1.
func MapStrUnion(dict1 MapStr, dict2 MapStr) MapStr {
	dict := MapStr{}

	for k, v := range dict1 {
		dict[k] = v
	}

	for k, v := range dict2 {
		dict[k] = v
	}
	return dict
}

// MergeFields merges the top-level keys and values in each source map (it does
// not perform a deep merge). If the same key exists in both, the value in
// fields takes precedence. If underRoot is true then the contents of the fields
// MapStr is merged with the value of the 'fields' key in ms.
//
// An error is returned if underRoot is true and the value of ms.fields is not a
// MapStr.
func MergeFields(ms, fields MapStr, underRoot bool) error {
	if ms == nil || len(fields) == 0 {
		return nil
	}

	fieldsMS := ms
	if !underRoot {
		f, ok := ms[FieldsKey]
		if !ok {
			fieldsMS = make(MapStr, len(fields))
			ms[FieldsKey] = fieldsMS
		} else {
			// Use existing 'fields' value.
			var err error
			fieldsMS, err = toMapStr(f)
			if err != nil {
				return err
			}
		}
	}

	// Add fields and override.
	for k, v := range fields {
		fieldsMS[k] = v
	}

	return nil
}

// AddTags appends a tag to the tags field of ms. If the tags field does not
// exist then it will be created. If the tags field exists and is not a []string
// then an error will be returned. It does not deduplicate the list of tags.
func AddTags(ms MapStr, tags []string) error {
	if ms == nil || len(tags) == 0 {
		return nil
	}

	tagsIfc, ok := ms[TagsKey]
	if !ok {
		ms[TagsKey] = tags
		return nil
	}

	existingTags, ok := tagsIfc.([]string)
	if !ok {
		return errors.Errorf("expected string array by type is %T", tagsIfc)
	}

	ms[TagsKey] = append(existingTags, tags...)
	return nil
}

// toMapStr performs a type assertion on v and returns a MapStr. v can be either
// a MapStr or a map[string]interface{}. If it's any other type or nil then
// an error is returned.
func toMapStr(v interface{}) (MapStr, error) {
	m, ok := tryToMapStr(v)
	if !ok {
		return nil, errors.Errorf("expected map but type is %T", v)
	}
	return m, nil
}

func tryToMapStr(v interface{}) (MapStr, bool) {
	switch m := v.(type) {
	case MapStr:
		return m, true
	case map[string]interface{}:
		return MapStr(m), true
	default:
		return nil, false
	}
}

// walkMap walks the data MapStr to arrive at the value specified by the key.
// The key is expressed in dot-notation (eg. x.y.z). When the key is found then
// the given mapStrOperation is invoked.
func walkMap(key string, data MapStr, op mapStrOperation) (interface{}, error) {
	var err error
	keyParts := strings.Split(key, ".")

	// Walk maps until reaching a leaf object.
	m := data
	for i, k := range keyParts[0 : len(keyParts)-1] {
		v, exists := m[k]
		if !exists {
			if op.CreateMissingKeys {
				newMap := MapStr{}
				m[k] = newMap
				m = newMap
				continue
			}
			return nil, errors.Wrapf(ErrKeyNotFound, "key=%v", strings.Join(keyParts[0:i+1], "."))
		}

		m, err = toMapStr(v)
		if err != nil {
			return nil, errors.Wrapf(err, "key=%v", strings.Join(keyParts[0:i+1], "."))
		}
	}

	// Execute the mapStrOperator on the leaf object.
	v, err := op.Do(keyParts[len(keyParts)-1], m)
	if err != nil {
		return nil, errors.Wrapf(err, "key=%v", key)
	}

	return v, nil
}

// mapStrOperation types

// These are static mapStrOperation types that store no state and are reusable.
var (
	opDelete = mapStrOperation{deleteOperation{}, false}
	opGet    = mapStrOperation{getOperation{}, false}
	opHasKey = mapStrOperation{hasKeyOperation{}, false}
)

// mapStrOperation represents an operation that can be applied to map.
type mapStrOperation struct {
	mapStrOperator
	CreateMissingKeys bool
}

// mapStrOperator is an interface with a single function that performs an
// operation on a MapStr.
type mapStrOperator interface {
	Do(key string, data MapStr) (value interface{}, err error)
}

type deleteOperation struct{}

func (op deleteOperation) Do(key string, data MapStr) (interface{}, error) {
	value, found := data[key]
	if !found {
		return nil, ErrKeyNotFound
	}
	delete(data, key)
	return value, nil
}

type getOperation struct{}

func (op getOperation) Do(key string, data MapStr) (interface{}, error) {
	value, found := data[key]
	if !found {
		return nil, ErrKeyNotFound
	}
	return value, nil
}

type hasKeyOperation struct{}

func (op hasKeyOperation) Do(key string, data MapStr) (interface{}, error) {
	_, found := data[key]
	return found, nil
}

type putOperation struct {
	Value interface{}
}

func (op putOperation) Do(key string, data MapStr) (interface{}, error) {
	existingValue, _ := data[key]
	data[key] = op.Value
	return existingValue, nil
}
