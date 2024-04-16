package main

import (
	"io"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

// func main() {
// 	// client := g.Client()
// 	// response, _ := client.Get("http://test.com/index.php?s=aaaa", g.Map{
// 	// 	"uid": 1,
// 	// })
// 	// defer response.Close()
// 	// // response.
// 	// g.Dump(response.RawRequest())
// 	testUseFunc()
// }

func testUseFunc() {
	response, _ := g.Client().Use(afterFunc).Post(gctx.New(), "http://test.com/index.php?s=aaaa", g.Map{"test": "tttest body"})
	defer response.Close()
	response.RawDump()
}

// func afterFunc(c *gclient.Client, r *http.Request) (resp *gclient.Response, err error) {
// 	reqBodyContent, _ := io.ReadAll(r.Body)
// 	g.Dump(reqBodyContent)
// 	r.Body = NewReadCloser(reqBodyContent, false)

// 	resp, err = c.Next(r)

// 	//
// 	bodyContent := resp.ReadAll()
// 	// auth
// 	s := gconv.String(bodyContent)
// 	g.Dump(s)
// 	// resp.SetBodyContent(bodyContent)
// 	return resp, err
// }

// ReadCloser implements the io.ReadCloser interface
// which is used for reading request body content multiple times.
//
// Note that it cannot be closed.
type ReadCloser struct {
	index      int    // Current read position.
	content    []byte // Content.
	repeatable bool   // Mark the content can be repeatable read.
}

// NewReadCloser creates and returns a RepeatReadCloser object.
func NewReadCloser(content []byte, repeatable bool) io.ReadCloser {
	return &ReadCloser{
		content:    content,
		repeatable: repeatable,
	}
}

// Read implements the io.ReadCloser interface.
func (b *ReadCloser) Read(p []byte) (n int, err error) {
	// Make it repeatable reading.
	if b.index >= len(b.content) && b.repeatable {
		b.index = 0
	}
	n = copy(p, b.content[b.index:])
	b.index += n
	if b.index >= len(b.content) {
		return n, io.EOF
	}
	return n, nil
}

// Close implements the io.ReadCloser interface.
func (b *ReadCloser) Close() error {
	return nil
}
