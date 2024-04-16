package main

import (
	"io"
	"net/http"

	"gitee.com/wallesoft/ewa/kernel/log"

	"gitee.com/wallesoft/ewa/internal/utils"
	ehttp "gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	client := &ehttp.Client{
		Client:       gclient.New(),
		BaseUri:      "http://test.com/",
		AfterReponse: afterFunc,
	}
	logger := log.New()
	logger.SetPath("d:/tmp/logs")
	client.Logger = logger
	client.Logger.LogStdout = true
	client.Logger.AccessLogEnabled = true
	client.RequestJson(gctx.New(), "GET", "index.php", g.Map{"test": "tttest body"})
}

func afterFunc(c *gclient.Client, r *http.Request) (resp *gclient.Response, err error) {
	reqBodyContent, _ := io.ReadAll(r.Body)
	g.Dump(reqBodyContent)
	r.Body = utils.NewReadCloser(reqBodyContent, false)

	resp, err = c.Next(r)

	//
	bodyContent := resp.ReadAll()
	// auth
	// resp.RawDump()
	// s := gconv.String(bodyContent)
	g.Dump(resp.Raw())
	resp.SetBodyContent(bodyContent)
	return resp, err
}
