package helper

type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Errors interface{} `json:"errors"`
	Data interface{} `json:"data"`
}

type Empty struct {
	
}

func BuildResponse(status int, message string, data interface{})  Response {
	res := Response{
		Status: status,
		Message: message,
		Errors: nil,
		Data: data,
	}

	return res
}

func BuildErrorResponse(message string,error error) Response {
	res := Response{
		Message: message,
		Errors: error.Error(),
	}

	return res
}