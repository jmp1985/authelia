package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFileFilters(t *testing.T) {
	testCases := []struct {
		name   string
		have   []string
		expect string
	}{
		{
			"ShouldErrorOnInvalidFilterName",
			[]string{"abc"},
			"invalid filter named 'abc'",
		},
		{
			"ShouldErrorOnInvalidFilterNameWithDuplicates",
			[]string{"abc", "abc"},
			"invalid filter named 'abc'",
		},
		{
			"ShouldErrorOnInvalidFilterNameWithDuplicatesCaps",
			[]string{"ABC", "abc"},
			"invalid filter named 'abc'",
		},
		{
			"ShouldErrorOnDuplicateFilterName",
			[]string{"template", "template"},
			"duplicate filter named 'template'",
		},
		{
			"ShouldErrorOnDuplicateFilterNameCaps",
			[]string{"TEMPLATE", "template"},
			"duplicate filter named 'template'",
		},
		{
			"ShouldNotErrorOnValidFilters",
			[]string{"template"},
			"",
		},
		{
			"ShouldErrorOnExpandEnvFilter",
			[]string{"expand-env"},
			"invalid filter named 'expand-env'",
		},
		{
			"ShouldNotErrorOnTemplateFilter",
			[]string{"template"},
			"",
		},
		{
			"ShouldNotErrorOnTemplateFilterCaps",
			[]string{"TEMPLATE"},
			"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, theError := NewFileFilters(tc.have, "", "")

			switch tc.expect {
			case "":
				assert.NoError(t, theError)
				assert.Len(t, actual, len(tc.have))
			default:
				assert.EqualError(t, theError, tc.expect)
			}
		})
	}
}
