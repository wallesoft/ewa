package encryptor

//Error ecnryptor error
type Error struct {
	//err     error
	code    int
	message string
}

//NewError new error
func NewError(code int, message string) *Error {
	return &Error{
		code:    code,
		message: message,
	}
}

//GetCode return error code
func (e *Error) GetCode() int {
	return e.code
}

//GetMessage alias of Error()
func (e *Error) GetMessage() string {
	return e.message
}

//Error
func (e *Error) Error() string {
	return e.message
}
