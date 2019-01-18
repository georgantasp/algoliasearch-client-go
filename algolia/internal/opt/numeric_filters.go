package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractNumericFilters(opts ...interface{}) *opt.NumericFiltersOption {
	for _, o := range opts {
		if v, ok := o.(opt.NumericFiltersOption); ok {
			return &v
		}
	}
	return nil
}
