package appresponse

type responseSuccess struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	//Debug *interface{} `json:"debug,omitempty"`
}

type responseError struct {
	Code    string `json:"int"`
	Message string `json:"message"`
}

func Success(data interface{}) responseSuccess {
	if data == nil {
		type Empty struct{}
		data = Empty{}
	}

	return responseSuccess{
		Code:    "BOT-2000",
		Message: "Success.",
		Data:    data,
	}
}

func Error(message string) responseError {
	return responseError{
		Code:    "BOT-5000",
		Message: message,
	}
}
