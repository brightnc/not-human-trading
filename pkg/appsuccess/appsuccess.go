package appsuccess

import "fmt"

type AppSuccess struct {
	Code    string
	Message string
}

const (
	svcPrefix string = "NHT"
)

func code(code successCode) string {
	return fmt.Sprintf("%s-%s", svcPrefix, code)
}

// NewInternalServerError ...
func OK() AppSuccess {
	return AppSuccess{
		Code:    code(ok),
		Message: string(success),
	}
}
