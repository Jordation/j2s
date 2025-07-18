package main

import (
	"io"
	"strings"
)

type fmtFunc func(string) string
type fieldNameAndTypeFmtFunc func(*valueType, fmtFunc) string

type goTypeWriter struct {
	ir []*typeIR
}

func NewGoTypeWriter(
	setsOfTypes map[string][]*setW,
) *goTypeWriter {
	t := &goTypeWriter{
		ir: buildTypeRepresentations(setsOfTypes),
	}

	return t
}

func (tw *goTypeWriter) WriteTo(w io.Writer) (int64, error) {
	fileHead, fileTail := goFileWrap()
	w.Write([]byte(fileHead))
	for _, t := range tw.ir {
		sb := strings.Builder{}
		typeHead, typeTail := goTypeWrap(goRenamer(t.Name))

		sb.WriteString(typeHead)
		t.format(&sb, goRenamer, goTypeFormatter)
		sb.WriteString(typeTail)
		w.Write([]byte(sb.String()))
	}

	w.Write([]byte(fileTail))
	return 0, nil
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

func goFileWrap() (string, string) {
	return "package main\n\n", "\n"
}

func goTypeFormatter(vt *valueType, nameTransformer fmtFunc) string {
	switch vt.base {
	default:
		panic("unknown type repr " + vt.String())
	case vt_string:
		return "string"
	case vt_number:
		if vt.sub == vt_int {
			return "int"
		}
		return "float64"
	case vt_bool:
		return "bool"
	case vt_null:
		return "any"
	case vt_object:
		return "*" + nameTransformer(vt.typedName)

	case vt_array:
		switch vt.sub {
		default:
			panic("unknown type repr " + vt.String())
		case vt_string:
			return "[]string"
		case vt_int:
			return "[]int"
		case vt_float:
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
