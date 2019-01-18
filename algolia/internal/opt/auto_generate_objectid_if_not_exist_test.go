package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestAutoGenerateObjectIDIfNotExist(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.AutoGenerateObjectIDIfNotExistOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.AutoGenerateObjectIDIfNotExist(false),
		},
		{
			opts:     []interface{}{opt.AutoGenerateObjectIDIfNotExist(true)},
			expected: opt.AutoGenerateObjectIDIfNotExist(true),
		},
		{
			opts:     []interface{}{opt.AutoGenerateObjectIDIfNotExist(false)},
			expected: opt.AutoGenerateObjectIDIfNotExist(false),
		},
		{
			opts: []interface{}{
				opt.AutoGenerateObjectIDIfNotExist(false),
				opt.AutoGenerateObjectIDIfNotExist(true),
			},
			expected: opt.AutoGenerateObjectIDIfNotExist(false),
		},
		{
			opts: []interface{}{
				opt.AutoGenerateObjectIDIfNotExist(true),
				opt.AutoGenerateObjectIDIfNotExist(false),
			},
			expected: opt.AutoGenerateObjectIDIfNotExist(true),
		},
	} {
		var (
			in  = ExtractAutoGenerateObjectIDIfNotExist(c.opts...)
			out opt.AutoGenerateObjectIDIfNotExistOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, c.expected, out)
	}
}
