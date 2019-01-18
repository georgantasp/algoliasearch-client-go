package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractAutoGenerateObjectIDIfNotExist(opts ...interface{}) opt.AutoGenerateObjectIDIfNotExistOption {
	for _, o := range opts {
		if v, ok := o.(opt.AutoGenerateObjectIDIfNotExistOption); ok {
			return v
		}
	}
	return opt.AutoGenerateObjectIDIfNotExist(false)
}
