package kernel

//Request
type Request struct {
	Timestamp    string
	Nonce        string
	EncryptType  string
	MsgSignature string
	BodyRaw      []byte
}

func (r *Request) IsSafeMode() bool {
	return r.
}
