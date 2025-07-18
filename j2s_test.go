package main

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestTypeWritersFromJSON(t *testing.T) {
	tests := []struct {
		input, output, rootName string
	}{
		{
			input:    "samples/inputs/ship.json",
			output:   "samples/outputs/ship",
			rootName: "Summary",
		},
		{
			input:    "samples/inputs/mlbsummary.json",
			output:   "samples/outputs/mlbsummary",
			rootName: "Ship",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			srcFile, err := os.OpenFile(test.input, 0, 0644)
			assert.NoError(t, err)

			defer srcFile.Close()

			setsOfTypes, err := NewTranslator(test.rootName, srcFile).Translate()
			if err != nil {
				logrus.WithError(err).Fatal("failed to build sets")
			}

			tws := map[string]io.WriterTo{
				".go": NewGoTypeWriter(setsOfTypes),
				".ts": NewTsTypeWriter(setsOfTypes),
			}

			for ext, tw := range tws {
				outputFile, err := os.OpenFile(test.output+ext, os.O_CREATE|os.O_RDWR, 0644)
				assert.NoError(t, err)
				defer outputFile.Close()

				outputFile.Truncate(0)

				_, err = tw.WriteTo(outputFile)
				assert.NoError(t, err)
			}
		})
	}
}

// TODO: set up proper tests for xml translator..
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

func writeJson(path string, data any) {
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}

	if err = os.WriteFile(path, jsonBytes, 0644); err != nil {
		panic(err)
	}
}
