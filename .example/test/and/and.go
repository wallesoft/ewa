package main

import (
	"github.com/gogf/gf/encoding/gbinary"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func main() {
	var a string = "test"
	var b string = "test"
	abit := gbinary.DecodeToInt(gbinary.EncodeString(a))
	bbit := gbinary.DecodeToInt(gbinary.EncodeString(b))
	g.Dump(gconv.Int(a))
	g.Dump(gconv.Int(a) & gconv.Int(b))
	g.Dump(2 | 4)
	g.Dump(gbinary.DecodeToString(gbinary.EncodeInt(abit & bbit)))
}
