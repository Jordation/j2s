package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGoTypeBuilder(t *testing.T) {
	base := "./testdata/ships/ship-%d.json"
	bigSet := map[string][]*setW{}
	dupes := map[string][]*setW{}
	deduped := map[string][]*setW{}

	for i := range 30 {
		f, err := os.OpenFile(fmt.Sprintf(base, i), os.O_RDONLY, os.ModeAppend)
		if err != nil {
			t.Fatal(err)
		}

		jt := NewTranslator("Ship", f)
		result, err := jt.Translate()
		if err != nil {
			t.Fatal(err)
		}

		for k, sets := range result {
			merged, ok := bigSet[k]
			if !ok {
				merged = make([]*setW, 0, len(sets))
			}
			merged = append(merged, sets...)
			bigSet[k] = merged

		}

	}

	for k, sets := range bigSet {
		if len(sets) < 2 {
			continue
		}

		sample := sets[0]
		if sample.AllEq(sets[1:]...) {
			deduped[k] = []*setW{sample}
		} else {
			dupes[k] = append(dupes[k], reduce(sets)...)
		}

	}

	types := buildTypeRepresentations(bigSet)
	writeJson("types.json", types)
	writeJson("dupes.json", dupes)
}

func TestTypescriptTypeBuilder(t *testing.T) {
	base := "./testdata/ships/ship-%d.json"
	bigSet := map[string][]*setW{}
	dupes := map[string][]*setW{}
	deduped := map[string][]*setW{}

	for i := range 30 {
		f, err := os.OpenFile(fmt.Sprintf(base, i), os.O_RDONLY, os.ModeAppend)
		if err != nil {
			t.Fatal(err)
		}

		jt := NewTranslator("Ship", f)
		result, err := jt.Translate()
		if err != nil {
			t.Fatal(err)
		}

		for k, sets := range result {
			merged, ok := bigSet[k]
			if !ok {
				merged = make([]*setW, 0, len(sets))
			}
			merged = append(merged, sets...)
			bigSet[k] = merged

		}

	}

	for k, sets := range bigSet {
		if len(sets) < 2 {
			continue
		}

		sample := sets[0]
		if sample.AllEq(sets[1:]...) {
			deduped[k] = []*setW{sample}
		} else {
			dupes[k] = append(dupes[k], reduce(sets)...)
		}

	}

	tw := NewTsTypeWriter(bigSet)
	outFile, err := os.OpenFile("types.ts", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer outFile.Close()

	tw.WriteTo(outFile)
}
func TestXmlTranslator(t *testing.T) {
	resultSets := map[string][]*setW{}

	f, err := os.OpenFile("testdata/sample.xml", os.O_RDONLY, os.ModeAppend)
	if err != nil {
		t.Fatal(err)
	}

	xt := NewXmlTranslator(f)
	result, err := xt.Translate()
	if err != nil {
		t.Fatal(err)
	}

	for k, sets := range result {
		merged, ok := resultSets[k]
		if !ok {
			merged = make([]*setW, 0, len(sets))
		}
		merged = append(merged, sets...)
		resultSets[k] = merged

	}

	dupes := map[string][]*setW{}
	deduped := map[string][]*setW{}
	for k, sets := range resultSets {
		if len(sets) < 2 {
			continue
		}

		sample := sets[0]

		if sample.AllEq(sets[1:]...) {
			deduped[k] = []*setW{sample}
			continue
		}

		dupes[k] = append(dupes[k], reduce(sets)...)
	}

	writeJson("xmltypes.json", resultSets)

	// types := buildTypeRepresentations(resultSets)
	// writejson("types.json", types)
	// writejson("dupes.json", dupes)
}

func TestThat(t *testing.T) {
	type testType struct {
		ShipRawMessage []json.RawMessage
	}

	shipFolder := "./testdata/ships"
	ships, err := os.ReadDir(shipFolder)
	if err != nil {
		t.Fatal(err)
	}

	out := &testType{}
	for _, ship := range ships {
		if ship.IsDir() {
			continue
		}

		shipFile, err := os.OpenFile(filepath.Join(shipFolder, ship.Name()), os.O_RDONLY, os.ModeAppend)
		if err != nil {
			t.Fatal(err)
		}

		shipBytes, err := io.ReadAll(shipFile)
		if err != nil {
			t.Fatal(err)
		}

		out.ShipRawMessage = append(out.ShipRawMessage, json.RawMessage(shipBytes))
	}

	jsonData, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	os.WriteFile("testdata/ShipsCombined.json", jsonData, 0644)
}

func getTestSet(chars string) set {
	out := set{}
	entries := strings.Split(chars, "")

	for _, e := range entries {
		out[e] = &valueType{}
	}

	return out
}

func writeJson(path string, data any) {
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}

	if err = os.WriteFile(path, jsonBytes, 0644); err != nil {
		panic(err)
	}
}
