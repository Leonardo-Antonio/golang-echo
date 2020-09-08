package handler

const (
	// ERROR .
	ERROR = "error"
	// MESSAGE .
	MESSAGE = "message"
)

type response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

// newResponse .
func newResponse(msjT, msj string, data interface{}) response {
	return response{msjT, msj, data}
}
