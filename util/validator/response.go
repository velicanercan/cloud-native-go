package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ErrResponse struct {
	Errors []string `json:"errors"`
}

func ToErrResponse(errs error) *ErrResponse {
	if fieldErrors, ok := errs.(validator.ValidationErrors); ok {
		resp := ErrResponse{
			Errors: make([]string, len(fieldErrors)),
		}
		for i, err := range fieldErrors {
			switch err.Tag() {
			case "required":
				resp.Errors[i] = fmt.Sprintf("%s is required", err.Field())
			case "max":
				resp.Errors[i] = fmt.Sprintf("%s cannot be longer than %s", err.Field(), err.Param())
			case "url":
				resp.Errors[i] = fmt.Sprintf("%s is not a valid URL", err.Field())
			case "datetime":
				if err.Param() == "2006-01-02" {
					resp.Errors[i] = fmt.Sprintf("%s is not a valid date", err.Field())
				} else {
					resp.Errors[i] = fmt.Sprintf("%s must be in %s format", err.Field(), err.Param())
				}
			case "alphaspace":
				resp.Errors[i] = fmt.Sprintf("%s can only contain letters and spaces", err.Field())
			default:
				resp.Errors[i] = fmt.Sprintf("something went wrong with %s; %s", err.Field(), err.Tag())
			}
		}
		return &resp
	}
	return nil
}
