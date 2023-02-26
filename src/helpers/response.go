package helpers

type Response struct {
	StatusCode int         `json:"statusCode"`
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func BuildResponse(message string) Response {
	res := Response{
		StatusCode: 200,
		Status:     "sucess",
		Message:    message,
		// Data:       data,
	}
	return res
}

func ErrorResponse(statusCode int, status string, message string) Response {
	res := Response{
		StatusCode: statusCode,
		Status:     status,
		Message:    message,
	}
	return res
}
