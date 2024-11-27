package utils

func ValidateStruct(data interface{}) error {
	validate := validator.New()
	return validate.Struct(data)
}
