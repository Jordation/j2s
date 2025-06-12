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
	typeSet    set
}

func newSetWrapper(set set, name, parentPath string) *setW {
	setw := &setW{
		typeSet:    set,
		ParentPath: parentPath,
		Name:       name,
		Hash:       set.hash(),
	}

	return setw
}

func (sw *setW) MarshalJSON() ([]byte, error) {
	start := []byte(fmt.Sprintf("{\"Path\":\"%s\",\"Set\": ", sw.ParentPath))
	other, err := json.Marshal(sw.typeSet)
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
	basePaths := strings.Split(base.ParentPath, ".")
	classifieePaths := strings.Split(classifiee.ParentPath, ".")
	// this looks fuckey
	// if the parent path (type name) on both sets are the same,
	// we can easily say it is the same type and at least partically equal
	if len(basePaths) > 1 && len(classifieePaths) > 1 {
		if basePaths[len(basePaths)-2] == classifieePaths[len(classifieePaths)-2] {
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

	// this is the main gamble and where maybe edge case breaks smth
	// we assume that if the mismatch count each way is < half the total set size
	// some keys are missing on each side but there is a majority overlap
	if baseDiff < len(base.typeSet)/2 &&
		classDiff < len(classifiee.typeSet)/2 {
		return sc_Partial
	}

	// we didn't match anywhere
	return sc_NotEq
}

func (s1 *setW) Difference(s2 *setW) int {
	diff := 0

	for k := range s1.typeSet {
		if _, ok := s2.typeSet[k]; !ok {
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

	maps.Copy(s1.typeSet, s2.typeSet)
	maps.Copy(s2.typeSet, s1.typeSet)
}

type valueType struct {
	base      valueTypePart
	sub       valueTypePart
	typedName string
}

type valueTypePart int8

const (
	vt_unknown valueTypePart = iota

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

func guessXmlVt(val string) valueTypePart {
	vt := vt_number
	isNumber := func(r rune) bool {
		return r >= '0' && r <= '9' || r == '.'
	}

	for _, r := range val {
		if !isNumber(r) {
			vt = vt_string
			break
		}
	}

	return vt
}

func toValueType(val any) (valueTypePart, valueTypePart, error) {
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

func nodeToValueType(node *xmlNode) *valueType {
	vt := &valueType{}

	if len(node.Nodes) == 0 && len(node.Attrs) > 0 {
		// we only have attrs, simple object
		vt.base = vt_object
		vt.typedName = node.XMLName.Local
	}
	if len(node.Nodes) > 0 {
		allSame := true
		commonName := node.Nodes[0].XMLName.Local

		for _, child := range node.Nodes {
			if child.XMLName.Local != commonName {
				allSame = false
				break
			}
		}

		if allSame {
			vt.base = vt_array
			vt.sub = nodeToValueType(node.Nodes[0]).base
		}
	}

	return vt
}

func (vt *valueType) MarshalJSON() ([]byte, error) {
	return []byte("{\"type\": \"" + vt.String() + "\"}"), nil
}
