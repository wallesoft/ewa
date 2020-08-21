package server

type Response struct {
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
