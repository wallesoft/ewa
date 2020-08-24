package server

type Response struct {
	Raw     []byte
	Content string
}

const defualtResponseContent = "success"

type ResponseContent string

func (r *Response) GetContent() (content ResponseContent) {
	content = r.Content
	if content == nil {
		content = defualtResponseContent
	}
	return
}

func (r *Response) GetRaw() {

}
