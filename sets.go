package j2s

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"golang.org/x/exp/errors"
	"golang.org/x/exp/maps"
)

// wraps a set of children with some extra metadata
// type set map[string]struct{}

type set map[string]*valueType

func (s set) hash() string {
	strs := maps.Keys(s)
	sort.Strings(strs)
	return strings.Join(strs, "-")
}

type setClassficiation int8

const (
	sc_NotEq setClassficiation = iota
	sc_Eq
	sc_Sub
	sc_Partial
)

type setW struct {
	ParentPath string
	Name       string
	Hash       string // sorted and joined values within the set, used for quick compare
	set        set
}

func NewSetWrapper(set set, name, parentPath string) *setW {
	setw := &setW{
		set:        set,
		ParentPath: parentPath,
		Name:       name,
		Hash:       set.hash(),
	}

	return setw
}

func (sw *setW) MarshalJSON() ([]byte, error) {
	start := []byte(fmt.Sprintf("{\"Path\":\"%s\",\"Set\": ", sw.ParentPath))
	other, err := json.Marshal(sw.set)
	if err != nil {
		return nil, err
	}

	combined := append(start, other...)
	combined = append(combined, '}')

	return combined, nil
}

func reduce(sets []*setW) []*setW {
	hashToSet := map[string]*setW{}

	for _, set := range sets {
		hashToSet[set.Hash] = set
	}

	return maps.Values(hashToSet)
}

func classify(base, classifiee *setW) setClassficiation {
	// if they live at the same path and have the same name, assume they are at least merge-worthy
	bs := strings.Split(base.ParentPath, ".")
	cs := strings.Split(classifiee.ParentPath, ".")
	// the second last path, AKA the parent
	if len(bs) > 1 && len(cs) > 1 {
		if bs[len(bs)-2] == cs[len(cs)-2] {
			return sc_Partial
		}
	}

	baseDiff := base.Difference(classifiee)
	classDiff := classifiee.Difference(base)

	// no difference each way, we are equal
	if baseDiff+classDiff == 0 {
		return sc_Eq
	}

	// base keys not found in classifiee but all keys found in base
	if baseDiff > 0 &&
		classDiff == 0 {
		return sc_Sub
	}

	// some keys are missing on each side but there is a majority overlap
	if baseDiff < len(base.set)/2 &&
		classDiff < len(classifiee.set)/2 {
		return sc_Partial
	}

	// we didn't match anywhere
	return sc_NotEq
}

func (self *setW) Difference(other *setW) int {
	diff := 0

	for k := range self.set {
		if _, ok := other.set[k]; !ok {
			diff++
		}
	}

	return diff
}

func (s1 *setW) AllEq(sets ...*setW) bool {
	for _, s := range sets {
		if !s1.Eq(s) {
			return false
		}
	}

	return true
}

func (s1 *setW) Eq(s2 *setW) bool {
	return s1.Hash == s2.Hash
}

func mergeSets(s1, s2 *setW) {

	maps.Copy(s1.set, s2.set)
	maps.Copy(s2.set, s1.set)
}

type valueType struct {
	base      rawValueType
	sub       rawValueType
	typedName string
}

type rawValueType int8

const (
	vt_unknown rawValueType = iota

	vt_string
	vt_number
	vt_bool
	vt_null

	vt_object
	vt_array
)

func (vt *valueType) String() string {
	switch vt.base {
	case vt_string:
		return "string"
	case vt_number:
		return "number"
	case vt_bool:
		return "bool"
	case vt_null:
		return "null"
	case vt_object:
		return "object"
	case vt_array:
		switch vt.sub {
		case vt_string:
			return "array of strings"
		case vt_number:
			return "array of numbers"
		case vt_bool:
			return "array of bools"
		case vt_null:
			return "array of nulls"
		case vt_object:
			return "array of objects"
		default:
			return "array of any"
			// panic(vt.sub)
		}
	}
	return "unknown"
}

func jsonValueType(val any) (rawValueType, rawValueType, error) {
	switch data := val.(type) {
	case string:
		return vt_string, 0, nil
	case float64:
		return vt_number, 0, nil
	case bool:
		return vt_bool, 0, nil
	case nil:
		return vt_null, 0, nil
	case map[string]any:
		return vt_object, 0, nil
	case []any:
		if len(data) == 0 {
			return vt_array, vt_unknown, nil
		}
		switch st := data[0].(type) {
		case string:
			return vt_array, vt_string, nil
		case float64:
			return vt_array, vt_number, nil
		case map[string]any:
			return vt_array, vt_object, nil
		case bool:
			return vt_array, vt_bool, nil
		case nil:
			return vt_array, vt_null, nil
		default:
			panic(st)
		}
	}

	return -1, -1, errors.New("invalid type")
}

func (vt *valueType) MarshalJSON() ([]byte, error) {
	return []byte("{\"type\": \"" + vt.String() + "\"}"), nil
}
