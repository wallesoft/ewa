package payment

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func (p *Payment) Certificates() {
	response := p.getClient().RequestJson(context.TODO(), "GET", "/v3/certificates")
	g.Dump("certs:----------------------------\n", response.ReadAllString())
	// @todo
}
