package payment

import "github.com/gogf/gf/frame/g"

func (p *Payment) Certificates() {
	certs := p.getClient().RequestJson("GET", "/v3/certificates")
	g.Dump("certs:----------------------------\n", certs)
}
