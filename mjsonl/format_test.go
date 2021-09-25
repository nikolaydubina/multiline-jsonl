package mjsonl_test

import (
	"bytes"
	"strings"
	"testing"

	_ "embed"

	"github.com/nikolaydubina/multiline-jsonl/mjsonl"
)

//go:embed testdata/example.jsonl
var example string

//go:embed testdata/example_short.jsonl
var exampleShort string

//go:embed testdata/example_expanded.jsonl
var exampleExpanded string

//go:embed testdata/example_gin.jsonl
var exampleGin string

//go:embed testdata/example_gin_short.jsonl
var exampleGinShort string

//go:embed testdata/example_gin_expanded.jsonl
var exampleGinExpanded string

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
		{
			name:      "when to short gin, then shortened",
			expand:    false,
			input:     exampleGin,
			expOutput: exampleGinShort,
		},
		{
			name:      "when to expand gin, then expanded",
			expand:    true,
			input:     exampleGin,
			expOutput: exampleGinExpanded,
		},
		{
			name:      "when empty input, then no error returned",
			expand:    true,
			input:     "",
			expOutput: "",
		},
		{
			name:      "when empty multiline input, then no error returned and all collapsed",
			expand:    true,
			input:     "\n\n\n\n",
			expOutput: "",
		},
		{
			name:      "when input contains empty objects, then no error returned",
			expand:    true,
			input:     "{}\n{}\n{}",
			expOutput: "{}\n{}\n{}\n",
		},
		{
			name:      "when not closed multiline, then return error",
			expand:    false,
			input:     `{"a": {"b": 123}}}`,
			expOutput: "{\"a\":{\"b\":123}}\n",
			expErr:    "got more } than {",
		},
		{
			name:      "when lines are arrays, then ignored",
			expand:    false,
			input:     `[]`,
			expOutput: "",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			inBuffer := strings.NewReader(tc.input)
			var outBuffer bytes.Buffer

			err := mjsonl.FormatJSONL(inBuffer, &outBuffer, tc.expand)
			if tc.expErr == "" {
				if err != nil {
					t.Errorf("expected no error, but got error(%v)", err)
				}
			} else {
				if err == nil {
					t.Error("expected error, but got nil")
				} else if !strings.Contains(err.Error(), tc.expErr) {
					t.Errorf("expected contain (%s) but got (%s)", tc.expErr, err.Error())
				}
			}

			out := outBuffer.String()
			if tc.expOutput != out {
				t.Errorf("expected:\n%v\n but got:\n%v\n", tc.expOutput, out)
			}
		})
	}
}
