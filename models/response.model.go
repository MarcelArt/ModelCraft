package models

type JSONResponse struct {
	Items     interface{} `json:"items"`
	IsSuccess bool        `json:"isSuccess"`
	Message   string      `json:"message"`
}

func NewJSONResponse(items interface{}, message string) *JSONResponse {
	err, ok := items.(error)
	if ok {
		if message == "" {
			message = err.Error()
		}
		return &JSONResponse{
			Items:     nil,
			IsSuccess: false,
			Message:   message,
		}
	}

	return &JSONResponse{
		Items:     items,
		IsSuccess: true,
		Message:   message,
	}
}
