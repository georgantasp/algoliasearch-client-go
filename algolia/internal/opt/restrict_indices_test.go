// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestRestrictIndices(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.RestrictIndicesOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.RestrictIndices([]string{}...),
		},
		{
			opts:     []interface{}{opt.RestrictIndices("value1")},
			expected: opt.RestrictIndices("value1"),
		},
		{
			opts:     []interface{}{opt.RestrictIndices("value1", "value2", "value3")},
			expected: opt.RestrictIndices("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractRestrictIndices(c.opts...)
			out opt.RestrictIndicesOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}

func TestRestrictIndices_CommaSeparatedString(t *testing.T) {
	for _, c := range []struct {
		payload  string
		expected *opt.RestrictIndicesOption
	}{
		{
			payload:  `""`,
			expected: opt.RestrictIndices([]string{}...),
		},
		{
			payload:  `"value1"`,
			expected: opt.RestrictIndices("value1"),
		},
		{
			payload:  `"value1,value2,value3"`,
			expected: opt.RestrictIndices("value1", "value2", "value3"),
		},
	} {
		var got opt.RestrictIndicesOption
		err := json.Unmarshal([]byte(c.payload), &got)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), got.Get())
	}
}
