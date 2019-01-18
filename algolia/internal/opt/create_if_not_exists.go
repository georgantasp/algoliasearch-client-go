package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractCreateIfNotExists(opts ...interface{}) opt.CreateIfNotExistsOption {
	for _, o := range opts {
		if v, ok := o.(opt.CreateIfNotExistsOption); ok {
			return v
		}
	}
	return opt.CreateIfNotExists(true)
}
