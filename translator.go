package j2s

import (
	"encoding/json"
	"io"
)

type jsonTranslator struct {
	RootName string
	Data     io.Reader
	original map[string]any
}

type translateResult struct {
	Sets map[string][]*setW
}

func NewTranslator(rootName string, jsonData io.Reader) *jsonTranslator {
	return &jsonTranslator{
		RootName: rootName,
		Data:     jsonData,
	}
}

func (jt *jsonTranslator) Translate() (*translateResult, error) {
	model := make(map[string]any)
	if err := json.NewDecoder(jt.Data).Decode(&model); err != nil {
		return nil, err
	}

	jt.original = model

	return jt.translate(model)
}

func (jt *jsonTranslator) translate(model map[string]any) (*translateResult, error) {
	var (
		sets   = make(map[string][]*setW)
		result = &translateResult{Sets: sets}
	)

	jt.buildKeySets("", jt.RootName, model, sets)

	return result, nil
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
		base, sub, err := jsonValueType(val)
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

	setOfChildren := NewSetWrapper(thisSet, currKey, pathTo)
	setsByField[currKey] = append(setsByField[currKey], setOfChildren)
}
