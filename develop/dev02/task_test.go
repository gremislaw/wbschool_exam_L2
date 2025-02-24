package main

import (
	"testing"
)

type testCase struct {
	name     string
	input    string
	expected string
	wantErr  bool
	errMsg   string
}

func TestUnpack(t *testing.T) {
	testCases := []testCase{
		{name: "Basic repetition", input: "a4bc2d5e", expected: "aaaabccddddde", wantErr: false},
		{name: "No repetition", input: "abcd", expected: "abcd", wantErr: false},
		{name: "Invalid starting with digit", input: "45", expected: "", wantErr: true, errMsg: "invalid string"},
		{name: "Empty string", input: "", expected: "", wantErr: false},
		{name: "Escaped digits", input: `qwe\4\5`, expected: "qwe45", wantErr: false},
		{name: "Escaped digit with repetition", input: `qwe\45`, expected: "qwe44444", wantErr: false},
		{name: "Escaped backslash with repetition", input: `qwe\\5`, expected: `qwe\\\\\`, wantErr: false},
		{name: "Single escaped backslash", input: `qwe\\`, expected: `qwe\`, wantErr: false},
		{name: "Invalid escape sequence", input: `qwe\`, expected: "", wantErr: true, errMsg: "invalid escape sequence"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Unpack(tc.input)

			if tc.wantErr {
				if err == nil {
					t.Errorf("Unpack(%q) expected error but got nil", tc.input)
				} else if err.Error() != tc.errMsg {
					t.Errorf("Unpack(%q) error = %q, expected error = %q", tc.input, err.Error(), tc.errMsg)
				}
			} else {
				if err != nil {
					t.Errorf("Unpack(%q) unexpected error: %v", tc.input, err)
				}
				if result != tc.expected {
					t.Errorf("Unpack(%q) = %q, expected %q", tc.input, result, tc.expected)
				}
			}
		})
	}
}
