package validators

type Validator interface {
	ValidateSomeField(inf interface{}) error
}
type validator struct{}

func New() Validator {
	return &validator{}
}

func (vld *validator) ValidateSomeField(inf interface{}) error {
	return nil
}
