package main

import (
	"github.com/gogf/gf/v2/encoding/gbinary"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func main() {
	var a string = "test"
	var b string = "test"
	abit := gbinary.DecodeToInt(gbinary.EncodeString(a))
	bbit := gbinary.DecodeToInt(gbinary.EncodeString(b))
	g.Dump(gconv.Int(a))
	g.Dump(gconv.Int(a) & gconv.Int(b))
	g.Dump(2 | 4)
	g.Dump((2 | 4) & 2)
	g.Dump(gbinary.DecodeToString(gbinary.EncodeInt(abit & bbit)))
}
