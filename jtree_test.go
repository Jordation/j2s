package j2s

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestDiffer(t *testing.T) {
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

		for k, sets := range result.Sets {
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
}

func TestBuildTypes(t *testing.T) {
	testSets := map[string][]*setW{}

	testSets["type_a"] = []*setW{
		{
			set:        GetNewTestSet("abce"),
			Name:       "type_a",
			ParentPath: "type_b.type_d",
			Hash:       "a-b-c-e",
		},
		{
			set:        GetNewTestSet("abcd"),
			Name:       "type_a",
			ParentPath: "type_b.type_d",
			Hash:       "a-b-c-d",
		},
		{
			set:        GetNewTestSet("abcxyz"),
			Name:       "type_a",
			ParentPath: "type_b.type_c",
			Hash:       "a-b-c-x-y-z",
		},
	}

	output := buildTypeRepresentations(testSets)
	writeJson("built types.json", output)

}

func GetNewTestSet(chars string) set {
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
