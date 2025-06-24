package main

import (
	"encoding/xml"
	"fmt"
	"io"
)

type xmlTranslator struct {
	Data     io.Reader
	original *xmlNode
}

// type xmlMap map[string]any

// this is just a type wrapper to hold the information for an xml node (and its children)
type xmlNode struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr"`
	Content []byte     `xml:",chardata"`
	Nodes   []*xmlNode `xml:",any"`
}

func NewXmlTranslator(xmlData io.Reader) *xmlTranslator {
	return &xmlTranslator{
		Data: xmlData,
	}
}

func (xt *xmlTranslator) Translate() (map[string][]*setW, error) {
	model := new(xmlNode)
	if err := xml.NewDecoder(xt.Data).Decode(model); err != nil {
		return nil, err
	}

	xt.original = model

	return xt.translate(model)
}

func (xt *xmlTranslator) translate(model *xmlNode) (map[string][]*setW, error) {
	var (
		sets = make(map[string][]*setW)
	)

	xt.buildKeySets("", model.XMLName.Local, model, sets)

	return sets, nil
}

func (xt *xmlTranslator) buildKeySets(
	pathTo string,
	currKey string,
	currNode *xmlNode,
	setsByField map[string][]*setW,
) {
	thisSet := set{}

	if pathTo != "" && currKey != "" {
		pathTo += "." + currKey
	} else {
		pathTo = currKey
	}
	for _, attr := range currNode.Attrs {
		// attrs have simple values
		vt := &valueType{base: guessXmlVt(attr.Value)}
		fmt.Printf("assigned value type %s to %s\n", vt.String(), attr.Name.Local)
		thisSet[attr.Name.Local] = vt
	}

	for _, node := range currNode.Nodes {
		vt := nodeToValueType(node)
		thisSet[node.XMLName.Local] = vt

		xt.buildKeySets(pathTo, node.XMLName.Local, node, setsByField)
	}

	setOfChildren := newSetWrapper(thisSet, currKey, pathTo)
	setsByField[currKey] = append(setsByField[currKey], setOfChildren)
}

func (n *xmlNode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	n.Attrs = start.Attr
	n.XMLName = start.Name

	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch t := token.(type) {
		case xml.CharData:
			n.Content = append([]byte{}, t...)
		case xml.StartElement:
			node := new(xmlNode)
			if err := d.DecodeElement(node, &t); err != nil {
				return err
			}
			n.Nodes = append(n.Nodes, node)
		case xml.EndElement:
			return nil
		}
	}
}
