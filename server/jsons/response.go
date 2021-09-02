package jsons

// A structure that resposne to the request in case of error
type Response struct {
	// The corresponding HTTP error code
	code int `json:"code"`

	// Error message that describe the reasons of the error
	message string `json:"message"`
}

func NewResponse(code int, message string) Response {
	return Response{code, message}
}
