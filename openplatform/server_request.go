package openplatform

//MessageReq
type Request struct {
	Timestamp    string
	Nonce        string
	EncryptType  string
	MsgSignature string
	BodyRaw      []byte
}

func CreateRequest() (request *Request) {
	request = &Request{}
	return
}
