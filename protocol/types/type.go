package types

import (
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
)

var ByteOrder = binary.BigEndian

type Type interface {

	Decode(r io.Reader) (interface{}, error)
	Encode(w io.Writer) error
}

func ErrorMismatchedType(expected string, actual interface{}) error {
	return fmt.Errorf(
		"cannot encode mismatched type (expected: %s, got: %s)",
		expected, reflect.TypeOf(actual),
	)
}
