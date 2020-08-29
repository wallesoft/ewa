package server

type Request interface {
	Get(key string, def ...interface{}) interface{}
	GetRaw() []byte
	GetUrl() string
}

//Request abstract request.
type DefaultRequest struct {
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
