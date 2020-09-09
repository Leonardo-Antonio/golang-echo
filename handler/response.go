package handler

const (
	// ERROR .
	ERROR = "error"
	// MESSAGE .
	MESSAGE = "message"
)

type (
	response struct {
		MessageType string      `json:"message_type"`
		Message     string      `json:"message"`
		Data        interface{} `json:"data"`
		Info        info        `json:"info"`
	}
	info struct {
		Code   int    `json:"code"`
		Path   string `json:"path"`
		Method string `json:"method"`
	}
)

// newResponse .
func newResponse(msjT, msj string, data interface{}, i info) response {
	return response{msjT, msj, data, i}
}
