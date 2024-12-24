package j2s

import (
	"fmt"
	"io"
	"strings"
)

type TypeWriter struct {
	types           []*typeRepr
	nameTransformer func(string) string
	typeTransformer func(*valueType, func(string) string) string
}

type typeRepr struct {
	Name   string
	Fields map[string]*valueType
}

func (r *TypeWriter) WriteTo(w io.Writer) (int64, error) {
	w.Write([]byte("package j2s\n\n"))
	for _, t := range r.types {
		b, _ := t.toGoStruct(r.nameTransformer)
		if _, err := w.Write(b); err != nil {
			return 0, err
		}
	}

	return 0, nil
}

func (tr *typeRepr) toGoStruct(renamer func(string) string) ([]byte, error) {
	b := strings.Builder{}

	b.WriteString(fmt.Sprintf("type %s struct {\n", renamer(tr.Name)))
	for fieldName, fieldType := range tr.Fields {
		b.WriteString(fmt.Sprintf("%s %s `json:\"%s\"`\n", renamer(fieldName), ToGoType(fieldType, renamer), fieldName))
	}

	b.WriteString("}\n")

	return []byte(b.String()), nil
}

func BuildTypeRepresentations(types map[string][]*setW, renamerFunc func(string) string) *TypeWriter {
	result := &TypeWriter{
		nameTransformer: renamerFunc,
		typeTransformer: ToGoType,
	}

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

	typeReprs := make([]*typeRepr, 0, len(outputSets))
	for name, set := range outputSets {
		typeReprs = append(typeReprs, &typeRepr{
			Name:   name,
			Fields: set.set,
		})
	}

	result.types = typeReprs
	return result
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
			mergeSets(existing, target)
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
	switch classify(s1, s2) {
	case sc_NotEq:
		return nil
	case sc_Sub, sc_Partial:
		mergeSets(s1, s2)
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

func DefaultRenamer(in string) string {
	words := strings.Split(in, "_")
	for i, word := range words {
		words[i] = capitalise(word)
	}

	return strings.Join(words, "")
}

func capitalise(word string) string {
	if len(word) > 0 {
		return strings.ToUpper(string(word[0])) + word[1:]
	}
	return ""
}

func ToGoType(vt *valueType, nameTransformer func(string) string) string {
	switch vt.base {
	default:
		panic("unknown type repr " + vt.String())
	case vt_string:
		return "string"
	case vt_number:
		return "float64"
	case vt_bool:
		return "bool"
	case vt_null:
		return "any"
	case vt_object:
		return nameTransformer(vt.typedName)

	case vt_array:
		switch vt.sub {
		default:
			panic("unknown type repr " + vt.String())
		case vt_string:
			return "[]string"
		case vt_number:
			return "[]float64"
		case vt_bool:
			return "[]bool"
		case vt_null, vt_unknown:
			return "[]any"
		case vt_object:
			return "[]*" + nameTransformer(vt.typedName)
		}

	}
}
