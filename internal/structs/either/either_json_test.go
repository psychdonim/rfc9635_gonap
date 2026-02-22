package either_test

import (
	"encoding/json"
	"testing"

	"github.com/psychdonim/rfc9635_gonap/internal/structs/either"
)

func TestEither_Unmarshal_Primitives(t *testing.T) {
	eitherIntString := struct {
		X either.Either[int, string] `json:"x"`
	}{}

	left := "{ \"x\": 1 }"
	json.Unmarshal([]byte(left), &eitherIntString)
	assertLeft(t, eitherIntString.X, 1)

	right := "{ \"x\": \"hello\" }"
	json.Unmarshal([]byte(right), &eitherIntString)
	assertRight(t, eitherIntString.X, "hello")
}

func TestEither_Unmarshal_Structs(t *testing.T) {
	eitherStruct := struct {
		X either.Either[struct{ A int }, struct{ B string }] `json:"x"`
	}{}

	left := "{ \"x\": { \"a\": 1 } }"
	json.Unmarshal([]byte(left), &eitherStruct)
	assertLeft(t, eitherStruct.X, struct{ A int }{A: 1})

	right := "{ \"x\": { \"b\": \"hello\" } }"
	json.Unmarshal([]byte(right), &eitherStruct)
	assertRight(t, eitherStruct.X, struct{ B string }{B: "hello"})
}

func assertLeft[L comparable, R any](t *testing.T, either either.Either[L, R], expected L) {
	if !either.IsLeft() {
		t.Errorf("Expected left, got right")
	}

	if either.Left() != expected {
		t.Errorf("Expected left = %v, got %v", expected, either.Left())
	}
}

func assertRight[L any, R comparable](t *testing.T, either either.Either[L, R], expected R) {
	if !either.IsRight() {
		t.Errorf("Expected right, got left")
	}

	if either.Right() != expected {
		t.Errorf("Expected right = %v, got %v", expected, either.Right())
	}
}
