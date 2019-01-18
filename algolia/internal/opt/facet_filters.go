package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractFacetFilters(opts ...interface{}) *opt.FacetFiltersOption {
	for _, o := range opts {
		if v, ok := o.(opt.FacetFiltersOption); ok {
			return &v
		}
	}
	return nil
}
