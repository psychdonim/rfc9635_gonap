package either

import (
	"bytes"
	"encoding/json"
	"errors"
)

type Either[L any, R any] struct {
	left   L
	right  R
	isLeft bool
}

func NewLeft[L any, R any](left L) Either[L, R] {
	return Either[L, R]{
		left:   left,
		isLeft: true,
	}
}

func NewRight[L any, R any](right R) Either[L, R] {
	return Either[L, R]{
		right:  right,
		isLeft: false,
	}
}

func (e Either[L, R]) IsLeft() bool {
	return e.isLeft
}

func (e Either[L, R]) IsRight() bool {
	return !e.isLeft
}

func (e Either[L, R]) Left() L {
	if e.isLeft {
		return e.left
	}
	var zero L
	return zero
}

func (e Either[L, R]) Right() R {
	if !e.isLeft {
		return e.right
	}
	var zero R
	return zero
}

func (e *Either[L, R]) UnmarshalJSON(data []byte) error {
	var left L
	if err := tryDecode(data, &left); err == nil {
		e.left = left
		e.isLeft = true
		return nil
	}

	var right R
	if err := tryDecode(data, &right); err == nil {
		e.right = right
		e.isLeft = false
		return nil
	}

	return errors.New("JSON doesn't match either types")
}

func tryDecode[T any](data []byte, v *T) error {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()

	if err := dec.Decode(v); err != nil {
		return err
	}

	if dec.More() {
		return errors.New("extra JSON data")
	}

	return nil
}
