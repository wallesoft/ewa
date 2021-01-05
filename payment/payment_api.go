package payment

import "github.com/gogf/gf/frame/g"

func (p *Payment) Certificates() {
	response := p.getClient().RequestJson("GET", "/v3/certificates")
	g.Dump("certs:----------------------------\n", response.ReadAllString())
}
