package miniprogram

type AppCode struct {
}

//获取小程序码，有次数限制，数量较少的业务场景
func (code *AppCode) Get() *ClientResponse {

}

//获取小程序码，适用于临时的业务场景
func (code *AppCode) GetUlimited() *ClientResponse {

}

//获取下程序二维码
func (code *AppCode) CreateQRCode() *ClientResponse {

}
