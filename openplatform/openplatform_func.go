package openplatform

import (
	"gitee.com/wallesoft/ewa/openplatform/auth"
)

//SetVrifyTicket
//需要自定义解决ticket的存储及获取问题时需要设置满足相关接口的对象
func (op *OpenPlatform) SetVerifyTicket(ticket auth.VerifyTicket) {
	op.verifyTicket = ticket
}

//GetVerigyTicket
func (op *OpenPlatform) GetVerifyTicket() string {
	return op.verifyTicket.GetTicket()
}
