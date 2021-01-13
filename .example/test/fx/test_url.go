package main

import (
	"fmt"
	"net/url"
)

func main() {
	var url = &url.Values{}
	url.Add("test", "test")
	url.Add("test", "test2")
	fmt.Println(fmt.Sprintf("ecode:%s", url.Encode()))
}
