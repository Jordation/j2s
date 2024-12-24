package j2s

import (
	"fmt"
	"io"
	"strings"
)

type typeWriter struct {
	types []*typeRepr
}

func NewTypeWriter(
	ir map[string][]*setW,
) *typeWriter {

	t := &typeWriter{
		types: buildTypeRepresentations(ir),
	}

	return t
}

func (tw *typeWriter) WriteTo(w io.Writer) (int64, error) {
	fileHead, fileTail := goFileWrap("")
	w.Write([]byte(fileHead))
	for _, t := range tw.types {
		sb := strings.Builder{}
		typeHead, typeTail := goTypeWrap(goRenamer(t.Name))

		sb.WriteString(typeHead)
		t.format(&sb)
		sb.WriteString(typeTail)
		w.Write([]byte(sb.String()))
	}

	w.Write([]byte(fileTail))
	return 0, nil
}

type typeRepr struct {
	Name   string
	Fields map[string]*valueType
}

func (tr *typeRepr) format(sb *strings.Builder) {
	for fieldName, fieldType := range tr.Fields {
		sb.WriteString(fmt.Sprintf("%s %s `json:\"%s\"`\n", goRenamer(fieldName), goTypeFormatter(fieldType, goRenamer), fieldName))
	}
}

func goRenamer(in string) string {
	words := strings.Split(in, "_")
	for i, word := range words {
		words[i] = capitalise(word)
	}

	return strings.Join(words, "")
}

func goTypeWrap(typeName string) (string, string) {
	return "type " + typeName + " struct {\n", "}\n\n"
}

func goFileWrap(string) (string, string) {
	return "package main\n\n", "\n"
}

func goTypeFormatter(vt *valueType, nameTransformer func(string) string) string {
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

func capitalise(word string) string {
	if len(word) > 0 {
		return strings.ToUpper(string(word[0])) + word[1:]
	}
	return ""
}
