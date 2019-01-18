package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractInsidePolygon(opts ...interface{}) *opt.InsidePolygonOption {
	for _, o := range opts {
		if v, ok := o.(opt.InsidePolygonOption); ok {
			return &v
		}
	}
	return nil
}
