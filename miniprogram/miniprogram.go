package miniprogram

type MiniProgram struct {
}

func GetApp(cfg *g.Map()) *MiniProgram {
	return &MiniProgram
}

//auth 小程序登录
func (mp *MiniProgram) Auth() *Auth {

}
