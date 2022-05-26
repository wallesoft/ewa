package main

import (
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	client := g.Client()
	response, _ := client.Get("http://test.com/index.php?s=aaaa", g.Map{
		"uid": 1,
	})
	defer response.Close()
	// response.
	g.Dump(response.RawRequest())
}
