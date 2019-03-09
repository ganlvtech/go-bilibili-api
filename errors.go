package api

type ResponseCodeNotZero struct {
	Message string
}

func (r *ResponseCodeNotZero) Error() string {
	return "response code not zero: " + r.Message
}

type ResponseJsonDecodeError struct {
	Message string
	Err     error
}

func (r *ResponseJsonDecodeError) Error() string {
	if r.Message != "" {
		return r.Message
	}
	return r.Err.Error()
}
