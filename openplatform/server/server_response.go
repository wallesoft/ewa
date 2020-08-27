package server

type Response struct {
	Raw     []byte
	Content ResponseContent
}

const defualtResponseContent = "success"

type ResponseContent string

func (r *Response) GetContent() (content ResponseContent) {
	content = r.Content
	if content == "" {
		content = defualtResponseContent
	}
	return
}

func (r *Response) GetRaw() {

}
