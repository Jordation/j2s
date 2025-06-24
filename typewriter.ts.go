package main

import (
	"fmt"
	"io"
	"strings"
)

type tsTypeWriter struct {
	types []*typeIR
}

func NewTsTypeWriter(
	setsOfTypes map[string][]*setW,
) *tsTypeWriter {
	t := &tsTypeWriter{
		types: buildTypeRepresentations(setsOfTypes),
	}
	return t
}

func (tw *tsTypeWriter) WriteTo(w io.Writer) (int64, error) {
	fileHead, fileTail := tsFileWrap()
	w.Write([]byte(fileHead))
	for _, t := range tw.types {
		sb := strings.Builder{}
		typeHead, typeTail := tsTypeWrap(tsRenamer(t.Name))

		sb.WriteString(typeHead)
		t.formatTs(&sb)
		sb.WriteString(typeTail)
		w.Write([]byte(sb.String()))
	}

	w.Write([]byte(fileTail))
	return 0, nil
}

func (tr *typeIR) formatTs(sb *strings.Builder) {
	for fieldName, fieldType := range tr.Fields {
		sb.WriteString(fmt.Sprintf("  %s: %s;\n", tsRenamer(fieldName), tsTypeFormatter(fieldType, tsRenamer)))
	}
}

func tsRenamer(in string) string {
	words := strings.Split(in, "_")
	for i, word := range words {
		if i == 0 {
			words[i] = strings.ToLower(word)
			continue
		}
		words[i] = capitalise(word)
	}
	return strings.Join(words, "")
}

func tsTypeWrap(typeName string) (string, string) {
	return "interface " + typeName + " {\n", "}\n\n"
}

func tsFileWrap() (string, string) {
	return "", ""
}

func tsTypeFormatter(vt *valueType, renamer fmtFunc) string {
	switch vt.base {
	case vt_string:
		return "string"
	case vt_number:
		return "number"
	case vt_bool:
		return "boolean"
	case vt_null:
		return "null"
	case vt_array:
		return tsTypeFormatter(&valueType{base: vt.sub}, renamer) + "[]"
	case vt_object:
		if vt.typedName != "" {
			return renamer(vt.typedName)
		}
		return "object"
	default:
		return "any"
	}
}
