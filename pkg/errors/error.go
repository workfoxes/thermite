package errors

// BaseError : Customer Base error Object
type BaseError struct {
	code    string
	message string
	args    map[string]string
}

func (e *BaseError) Error() string {
	return e.message
}

// InvalidCommand : returns an error that formats as the given text.
func InvalidCommand(cmd string) error {
	msg := "Enter Comment Is Invalid : " + cmd
	return &BaseError{code: "ER00001", message: msg}
}
