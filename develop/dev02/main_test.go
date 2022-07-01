package main

import (
	"errors"
	"testing"
)

func TestUnpack(t *testing.T) {
	unpackTests := []struct {
		arg         string
		expected    string
		expectedErr error
	}{
		{`a4bc2d5e`, "aaaabccddddde", nil},
		{`abcd`, "abcd", nil},
		{`45`, "", errors.New("incorrent s")},
		{"", "", nil},
		{`qwe\4\5`, "qwe45", nil},
		{`qwe\45`, "qwe44444", nil},
		{`qwe\\5`, `qwe\\\\\`, nil},
	}

	for _, test := range unpackTests {
		output, err := Unpack(test.arg)
		if output != test.expected || err != test.expectedErr {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
