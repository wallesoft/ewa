package http

import (
	"bytes"
	"net/http"

	"github.com/gogf/gf/encoding/gparser"
	"github.com/gogf/gf/util/gconv"
)

// Response
type Response struct {
	*ResponseWriter
	Writer *ResponseWriter
}

const (
	EXCEPTION_EXIT = "exit"
)

// GetResponse
func GetResponse(w http.ResponseWriter) *Response {
	r := &Response{
		ResponseWriter: &ResponseWriter{
			writer: w,
			buffer: bytes.NewBuffer(nil),
		},
	}
	r.Writer = r.ResponseWriter
	return r
}

// Write writes <content> to the response buffer.
func (r *Response) Write(content ...interface{}) {
	if r.hijacked || len(content) == 0 {
		return
	}
	if r.Status == 0 {
		r.Status = http.StatusOK
	}
	for _, v := range content {
		switch value := v.(type) {
		case []byte:
			r.buffer.Write(value)
		case string:
			r.buffer.WriteString(value)
		default:
			r.buffer.WriteString(gconv.String(v))
		}
	}
}

// WriteXml writes <content> to the response with XML format.
func (r *Response) WriteXml(content interface{}, rootTag ...string) error {
	// If given string/[]byte, response it directly to client.
	switch content.(type) {
	case string, []byte:
		r.Header().Set("Content-Type", "application/xml")
		r.Write(gconv.String(content))
		return nil
	}
	// Else use gparser.VarToXml function to encode the parameter.
	if b, err := gparser.VarToXml(content, rootTag...); err != nil {
		return err
	} else {
		r.Header().Set("Content-Type", "application/xml")
		r.Write(b)
	}
	return nil
}

// Note that do not set Content-Type header here.
func (r *Response) WriteStatus(status int, content ...interface{}) {
	r.WriteHeader(status)
	if len(content) > 0 {
		r.Write(content...)
	} else {
		r.Write(http.StatusText(status))
	}
}

func (r *Response) WriteStatusExit(status int, content ...interface{}) {
	r.WriteStatus(status, content)
	panic(EXCEPTION_EXIT)
}

// ClearBuffer clears the response buffer.
func (r *Response) ClearBuffer() {
	r.buffer.Reset()
}

// Output outputs the buffer content to the client and clears the buffer.
func (r *Response) Output() {
	r.Writer.Flush()
}
