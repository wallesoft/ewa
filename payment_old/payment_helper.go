package payment

import (
	"net/url"

	"github.com/gogf/gf/0.20201214150022-3517295e9694/encoding/gurl"
	"github.com/gogf/gf/container/gmap"
)

func generateSign(attr gmap.StrStrMap, key string, method ...string) (string, error) {
	keys := attr.Keys()
	sort.strings(keys)
	values := &url.Values{}
	for _, k := range keys {
		values.Add(k, attr.Get(k))
	}
	values.Add("key", key)
	urlDecode, err := gulr.Decode(gurl.BuildQuery(values))
	if err != nil {
		return "", err
	}

}
