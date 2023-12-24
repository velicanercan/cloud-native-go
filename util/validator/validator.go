package validator

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

const alphaspaceRegexString = "^[a-zA-Z ]+$"

func New() *validator.Validate {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	validate.RegisterValidation("alphaspace", func(fl validator.FieldLevel) bool {
		return regexp.MustCompile(alphaspaceRegexString).MatchString(fl.Field().String())
	})

	return validate
}
