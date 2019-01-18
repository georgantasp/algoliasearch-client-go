package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractInsideBoundingBox(opts ...interface{}) *opt.InsideBoundingBoxOption {
	for _, o := range opts {
		if v, ok := o.(opt.InsideBoundingBoxOption); ok {
			return &v
		}
	}
	return nil
}
