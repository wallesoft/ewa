package http

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// Request struct
type Request struct {
	*http.Request
	bodyContent []byte
	parsedQuery bool
	queryMap    map[string]interface{}
}

// Get get query param by key.
func (r *Request) Get(key string, def ...interface{}) interface{} {
	r.parseQuery()
	if len(r.queryMap) > 0 {
		if v, ok := r.queryMap[key]; ok {
			return v
		}
	}
	return nil
}

func (r *Request) GetQuery() map[string]interface{} {
	r.parseQuery()
	return r.queryMap
}
func (r *Request) GetString(key string, def ...interface{}) string {
	return gvar.New(r.Get(key, def...)).String()
}

// parseQuery parses query string into r.queryMap.
func (r *Request) parseQuery() {
	if r.parsedQuery {
		return
	}
	r.parsedQuery = true
	if r.URL.RawQuery != "" {
		var err error
		r.queryMap, err = gstr.Parse(gstr.Trim(r.URL.RawQuery))
		if err != nil {
			panic(err)
		}
	}
}

// GetURL returns current URL of this request.
func (r *Request) GetURL() string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return fmt.Sprintf(`%s://%s%s`, scheme, r.Host, r.URL.String())
}

// GetBody retrieves and returns request body content as bytes.
func (r *Request) GetBody() []byte {
	if r.bodyContent == nil {
		r.bodyContent, _ = ioutil.ReadAll(r.Body)
	}
	//trim

	return gconv.Bytes(gstr.Trim(gconv.String(r.bodyContent)))
}
