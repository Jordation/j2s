package main

import (
	"encoding/json"
	"io"
)

type jsonTranslator struct {
	RootName string
	Data     io.Reader
	original map[string]any
}

func NewTranslator(rootName string, jsonData io.Reader) *jsonTranslator {
	return &jsonTranslator{
		RootName: rootName,
		Data:     jsonData,
	}
}

func (jt *jsonTranslator) Translate() (map[string][]*setW, error) {
	model := make(map[string]any)
	if err := json.NewDecoder(jt.Data).Decode(&model); err != nil {
		return nil, err
	}

	jt.original = model

	return jt.translate(model)
}

func (jt *jsonTranslator) translate(model map[string]any) (map[string][]*setW, error) {
	var (
		sets = make(map[string][]*setW)
	)

	jt.buildKeySets("", jt.RootName, model, sets)

	return sets, nil
}

func (jt *jsonTranslator) buildKeySets(
	pathTo string,
	currKey string,
	currObj map[string]any,
	setsByField map[string][]*setW,
) {
	thisSet := set{}

	if pathTo != "" {
		pathTo += "." + currKey
	} else {
		pathTo = currKey
	}

	for fieldName, val := range currObj {
		base, sub, err := toValueType(val)
		if err != nil {
			panic(err)
		}

		vt := &valueType{
			base: base,
			sub:  sub,
		}

		if base == vt_object || sub == vt_object {
			vt.typedName = fieldName
		}
		thisSet[fieldName] = vt

		switch childVal := val.(type) {
		case int:
			panic("aren't they all floats?")
		case map[string]any:
			jt.buildKeySets(pathTo, fieldName, childVal, setsByField)
		case []any:
			for _, sample := range childVal {
				arrObjMember, isObj := sample.(map[string]any)
				if !isObj {
					continue
				}
				jt.buildKeySets(pathTo, fieldName, arrObjMember, setsByField)
			}
		}
	}

	setOfChildren := newSetWrapper(thisSet, currKey, pathTo)
	setsByField[currKey] = append(setsByField[currKey], setOfChildren)
}
