package mjsonl_test

import (
	"bytes"
	"strings"
	"testing"

	_ "embed"

	mjsonl "github.com/nikolaydubina/multiline-jsonl"
)

//go:embed testdata/example.jsonl
var example string

//go:embed testdata/example_short.jsonl
var exampleShort string

//go:embed testdata/example_expanded.jsonl
var exampleExpanded string

func TestFormatJSONL(t *testing.T) {
	tests := []struct {
		name      string
		expand    bool
		input     string
		expOutput string
		expErr    string
	}{
		{
			name:      "when to short, then shortened",
			expand:    false,
			input:     example,
			expOutput: exampleShort,
		},
		{
			name:      "when to expand, then expanded",
			expand:    true,
			input:     example,
			expOutput: exampleExpanded,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			inBuffer := strings.NewReader(tc.input)
			var outBuffer bytes.Buffer

			err := mjsonl.FormatJSONL(inBuffer, &outBuffer, tc.expand)
			if tc.expErr == "" && err != nil {
				t.Errorf("expected no error, but got error(%v)", err)
			}

			out := outBuffer.String()
			if tc.expOutput != out {
				t.Errorf("expected:\n%v\n but got:\n%v\n", tc.expOutput, out)
			}
		})
	}
}
