package kernel

//Request abstract request.
type Request struct {
	Signature    string
	Timestamp    string
	Nonce        string
	EncryptType  string
	MsgSignature string
	RawBody      []byte
}

// //IsSafeMode check the request message is the safe mode.
// func (r *Request) IsSafeMode() bool {
// 	return r.Signature != "" && r.EncryptType == "aes"
// }
