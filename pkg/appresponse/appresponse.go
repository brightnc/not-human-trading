package appresponse

type responseSuccess struct {
	Code  int          `json:"code"`
	Data  interface{}  `json:"data"`
	Debug *interface{} `json:"debug,omitempty"`
}

type responseError struct {
	Status  bool   `json:"status" example:"false"`
	Message string `json:"message" example:"example error message"`
}

func Success(data interface{}) responseSuccess {
	if result == nil {
		type Empty struct{}
		result = Empty{}
	}
	return responseSuccess{
		Code: true,
		Data: data,
	}
}

func Error(message string) responseError {
	return responseError{
		Status:  false,
		Message: message,
	}
}
