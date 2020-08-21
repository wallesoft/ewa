package server

//Request
type Request struct {
	Timestamp    string
	Nonce        string
	EncryptType  string
	MsgSignature string
	BodyRaw      []byte
}

//CreateRequest  create request
func CreateRequest() (request *Request) {
	request = &Request{}
	return
}
