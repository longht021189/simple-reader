package parser

import (
	"reflect"
	"strings"
	"testing"
)

func TestJSONParseWithMap(t *testing.T) {
	result := "abc"

	j := &jsonParser{
		key:  "a.b.c",
		data: "",
	}

	r := map[string]interface{}{
		"a": map[string]interface{}{
			"b": map[string]interface{}{
				"c": result,
			},
		},
	}

	j.keyParts = strings.Split(j.key, ".")
	err := j.parse(0, r)

	if err != nil {
		t.Fatalf("Parse Error: %v", err)
	}
	if len(j.output) != 1 {
		t.Fatalf("Output Length <> 1")
	}
	if value := reflect.ValueOf(j.output[0]).String(); value != result {
		t.Fatalf("Wrong Output")
	}
}

func TestJSONParseWithArray(t *testing.T) {
	result := "abc"

	j := &jsonParser{
		key:  "a.b.c",
		data: "",
	}

	r := map[string]interface{}{
		"a": []map[string]interface{}{
			{
				"b": map[string]interface{}{
					"c": result,
				},
			},
			{
				"b": map[string]interface{}{
					"c": result,
				},
			},
		},
	}

	j.keyParts = strings.Split(j.key, ".")
	err := j.parse(0, r)

	if err != nil {
		t.Fatalf("Parse Error: %v", err)
	}
	if len(j.output) != 2 {
		t.Fatalf("Output Length <> 2")
	}
	if value := reflect.ValueOf(j.output[0]).String(); value != result {
		t.Fatalf("Wrong Output")
	}
}

func TestJSONParseWithError(t *testing.T) {
	r := "test"
	j := &jsonParser{
		key:  "a.b.c",
		data: "",
	}

	j.keyParts = strings.Split(j.key, ".")
	err := j.parse(0, r)

	if err == nil {
		t.Fatalf("No Error")
	}
	if err != ErrorTypeNotMatch {
		t.Fatalf("Wrong Error")
	}
}
