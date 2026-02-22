package either_test

import (
	"testing"

	"github.com/psychdonim/rfc9635_gonap/internal/structs/either"
)

func TestEither_IsLeft(t *testing.T) {
	left := either.NewLeft[int, string](1)
	if !left.IsLeft() {
		t.Errorf("Expected left, got right")
	}

	right := either.NewRight[int, string]("hello")
	if right.IsLeft() {
		t.Errorf("Expected right, got left")
	}
}

func TestEither_IsRight(t *testing.T) {
	left := either.NewLeft[int, string](1)
	if left.IsRight() {
		t.Errorf("Expected left, got right")
	}

	right := either.NewRight[int, string]("hello")
	if !right.IsRight() {
		t.Errorf("Expected right, got left")
	}
}

func TestEither_Left(t *testing.T) {
	expected := 1
	var defaultExpected int

	left := either.NewLeft[int, string](expected)
	leftValue := left.Left()
	if leftValue != expected {
		t.Errorf("Expected left value %v, got %v", expected, leftValue)
	}

	right := either.NewRight[int, string]("hello")
	if right.Left() != defaultExpected {
		t.Errorf("Expected %v, got %v", defaultExpected, right.Left())
	}
}

func TestEither_Right(t *testing.T) {
	expected := "hello"
	var defaultExpected string

	left := either.NewLeft[int, string](1)
	if left.Right() != defaultExpected {
		t.Errorf("Expected %v, got %v", defaultExpected, left.Right())
	}

	right := either.NewRight[int, string](expected)
	rightValue := right.Right()
	if rightValue != expected {
		t.Errorf("Expected right value %v, got %v", expected, rightValue)
	}
}
