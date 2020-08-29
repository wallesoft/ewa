package server

//Request abstract request.
type Request struct {
	Signature    string
	Timestamp    string
	Nonce        string
	EncryptType  string
	MsgSignature string
	RawBody      []byte
	Uri          string
}

// // func (r *Request) Config(m map[string]interface{}) r *Reqeust
// type Request interface{}{

// }
