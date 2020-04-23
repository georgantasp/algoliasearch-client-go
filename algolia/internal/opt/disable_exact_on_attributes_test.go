// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestDisableExactOnAttributes(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.DisableExactOnAttributesOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.DisableExactOnAttributes([]string{}...),
		},
		{
			opts:     []interface{}{opt.DisableExactOnAttributes("value1")},
			expected: opt.DisableExactOnAttributes("value1"),
		},
		{
			opts:     []interface{}{opt.DisableExactOnAttributes("value1", "value2", "value3")},
			expected: opt.DisableExactOnAttributes("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractDisableExactOnAttributes(c.opts...)
			out opt.DisableExactOnAttributesOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}

func TestDisableExactOnAttributes_CommaSeparatedString(t *testing.T) {
	for _, c := range []struct {
		payload  string
		expected *opt.DisableExactOnAttributesOption
	}{
		{
			payload:  `""`,
			expected: opt.DisableExactOnAttributes([]string{}...),
		},
		{
			payload:  `"value1"`,
			expected: opt.DisableExactOnAttributes("value1"),
		},
		{
			payload:  `"value1,value2,value3"`,
			expected: opt.DisableExactOnAttributes("value1", "value2", "value3"),
		},
	} {
		var got opt.DisableExactOnAttributesOption
		err := json.Unmarshal([]byte(c.payload), &got)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), got.Get())
	}
}
