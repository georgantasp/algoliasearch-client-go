package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractFilters(opts ...interface{}) *opt.FiltersOption {
	for _, o := range opts {
		if v, ok := o.(opt.FiltersOption); ok {
			return &v
		}
	}
	return nil
}
