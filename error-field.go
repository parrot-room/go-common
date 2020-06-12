package common

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

//UsernameValidate validator
func UsernameValidate(fl validator.FieldLevel) bool {
	r, _ := regexp.Compile("[^a-z0-9_]")
	return r.MatchString(fl.Field().String()) == false
}

//MapErr type map error key value of object
type MapErr map[string]string

//FieldError customize error message
func FieldError(err error) []MapErr {
	var AllErr []MapErr
	var (
		errMsg      string
		GetMapError MapErr
	)
	castedObj, _ := err.(validator.ValidationErrors)
	for _, err := range castedObj {

		switch err.Tag() {
		case "required":
			errMsg = fmt.Sprintf("%v is required field", err.Field())
			break
		case "email":
			errMsg = "Field need email format"
			break
		case "min":
			errMsg = fmt.Sprintf("Character must greater than or equal %v", err.Param())
			break
		case "max":
			errMsg = fmt.Sprintf("Character lower or equal than %v", err.Param())
			break
		case "excludesall":
			errMsg = fmt.Sprintf("unicode not allowed")
			break
		case "alpha":
			errMsg = "character alpha only"
			break
		case "username":
			errMsg = "character only alphanumeric and underscors symbol"
			break
		case "dive":
			errMsg = "only accept array of JSON"
			break
		default:
			errMsg = "something happen, try again"
		}
		strField := fmt.Sprintf("%v", strcase.ToSnake(err.Field()))

		GetMapError = map[string]string{
			strField: errMsg,
		}
		AllErr = append(AllErr, GetMapError)
	}

	return AllErr

}
