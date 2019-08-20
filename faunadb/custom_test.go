package faunadb

import (
	"testing"
)

func TestSerializeLetPtr(t *testing.T) {
	bindings := Obj{"foobar": 777}
	newObj := Obj{"bob": 777}
	ptr := Obj{"alice": &newObj}
	query := LetPtr(bindings, &ptr)
	assertJSON(t, query,
		`{"in":{"object":{"alice":{"object":{"bob":777}}}},"let":[{"foobar":777}]}`,
	)
	// Test updating pointer object
	newObj["test"] = 999
	// Success: Does include test: 999
	assertJSON(t, query,
		`{"in":{"object":{"alice":{"object":{"bob":777,"test":999}}}},"let":[{"foobar":777}]}`,
	)
}

func TestSerializeLetOriginalPointerFAIL(t *testing.T) {
	newObj := Obj{"bob": 777}
	ptr := Obj{"alice": &newObj}
	query := Let().Bind("v1", Ref("collections/spells/42")).Bind("v2", Index("spells")).Bind("a1", Index("all_things")).In(&ptr)
	assertJSON(t, query,
		`{"in":{"object":{"alice":{"object":{"bob":777}}}},"let":[{"v1":{"@ref":"collections/spells/42"}},{"v2":{"index":"spells"}},{"a1":{"index":"all_things"}}]}`,
	)
	// Test updating pointer object
	newObj["test"] = 999
	// Fail: Does not include test: 999
	assertJSON(t, query,
		`{"in":{"object":{"alice":{"object":{"bob":777}}}},"let":[{"v1":{"@ref":"collections/spells/42"}},{"v2":{"index":"spells"}},{"a1":{"index":"all_things"}}]}`,
	)
}

// Simplified version
func TestSerializeLetOriginalPointer2FAIL(t *testing.T) {
	newNode := Obj{}
	query := Let().Bind("xxx", "X").In(&newNode)
	assertJSON(t, query, `{"in":{"object":{}},"let":[{"xxx":"X"}]}`)
	// Test updating pointer object
	newNode["test"] = 999
	// Fail: Does not include test: 999
	assertJSON(t, query, `{"in":{"object":{}},"let":[{"xxx":"X"}]}`)
}

func TestSerializeLetWithInPtr(t *testing.T) {
	newNode := Obj{}
	query := Let().Bind("xxx", "X").InPtr(&newNode)
	assertJSON(t, query, `{"in":{"object":{}},"let":[{"xxx":"X"}]}`)
	// Test updating pointer object
	newNode["test"] = 999
	// Success: Does include test: 999
	assertJSON(t, query, `{"in":{"object":{"test":999}},"let":[{"xxx":"X"}]}`)
}
