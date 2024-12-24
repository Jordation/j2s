package j2s

import (
	"strings"
)

type typeRepr struct {
	Name   string
	Fields map[string]*valueType
}

func BuildTypeRepresentations(types map[string][]*setW) []*typeRepr {
	outputSets := make(map[string]*setW, len(types))

	for typeName, fieldSets := range types {
		cont := true
		for cont {
			fieldSets, cont = collapse(outputSets, typeName, fieldSets)
			if !cont {
				break
			}
		}
	}

	out := make([]*typeRepr, 0, len(outputSets))
	for name, set := range outputSets {
		out = append(out, &typeRepr{
			Name:   name,
			Fields: set.set,
		})
	}

	return out
}

func collapse(dst map[string]*setW, typeName string, sets []*setW) ([]*setW, bool) {
	if len(sets) == 0 {
		return sets, false
	}

	var (
		target = sets[0]
		rem    = sets[1:]
	)

	offset := 0
	for i, set := range rem {
		newMerged := handleMerge(target, set)
		if newMerged != nil {
			target = newMerged
			rem = append(rem[:i-offset], rem[i+1-offset:]...)
			offset++
			continue
		}
	}

	existing, ok := dst[typeName]
	if ok {
		target.Name, existing.Name = handleNamingConflict(existing, target)
		if target.Name == existing.Name {
			MergeSets(existing, target)
		}
	}

	dst[target.Name] = target

	return rem, len(rem) != 0
}

func handleNamingConflict(s1, s2 *setW) (string, string) {
	s1Parents := strings.Split(s1.ParentPath, ".")
	s2Parents := strings.Split(s2.ParentPath, ".")

	if len(s1Parents) == 0 && len(s2Parents) == 0 {
		panic("couldn't handle type conflict")
	}

	s1Last := s1Parents[len(s1Parents)-1]
	s2Last := s2Parents[len(s2Parents)-1]
	if s1Last == s2Last {
		// same parent so we can just merge
		return s1Last, s2Last
	}

	return s1Last + s1.Name, s2Last + s2.Name
}

func handleMerge(s1, s2 *setW) *setW {
	switch Classify(s1, s2) {
	case sc_NotEq:
		return nil
	case sc_Sub, sc_Partial:
		MergeSets(s1, s2)
		ensureBestTypes(s1.set, s1.set)
	case sc_Eq:
		ensureBestTypes(s1.set, s1.set)
		// nothing to do
	}

	// copy
	c := *s1
	return &c
}

func ensureBestTypes(s1, s2 set) {
	for k, v := range s1 {
		if s2v, ok := s2[k]; ok {
			s1[k] = bestType(v, s2v)
		}
	}

	for k, v := range s2 {
		if s1v, ok := s1[k]; ok {
			s2[k] = bestType(v, s1v)
		}
	}
}

func bestType(vt1, vt2 *valueType) *valueType {
	if vt1.sub == vt_unknown && vt2.sub != vt_unknown {
		return vt2
	} else if vt1.sub != vt_unknown && vt2.sub == vt_unknown {
		return vt1
	} else if vt1.base == vt2.base && vt1.sub == vt2.sub {
		return vt1
	} else if vt1.base == vt_null && vt2.base != vt_null {
		return vt2
	} else if vt1.base != vt_null && vt2.base == vt_null {
		return vt1
	}

	return nil
}
