package payment

import (
	"github.com/gogf/gf/crypto/gmd5"
)

type Sandbox struct {
	payment *Payment
}

//Getkey
func (s *Sanbox) Getkey() string {
	if key , err := s.payment.Cache.Get(s.getCacheKey); err == nil {
		return gvar.New(key).String()
	} 
	---- request ---- get --- sign --- key ---
	
}

func (s *Sanbox) getCacheKey() string {
	return "ewa.payment.sanbox." + gmd5.MustEncrypt(s.payment.config.AppID+s.payment.config.MchID)
}
