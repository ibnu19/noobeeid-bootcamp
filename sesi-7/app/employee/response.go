package employee

type WebResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Payload any    `json:"payload,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func ApiResponse(msg string, payload any, err error) (response WebResponse) {
	var successTag bool
	switch err == nil {
	case true:
		successTag = true
	default:
		successTag = false
	}

	return WebResponse{
		Success: successTag,
		Message: msg,
		Payload: payload,
		Error:   err,
	}
}
