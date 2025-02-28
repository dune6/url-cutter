package response

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMssgs []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMssgs = append(errMssgs, fmt.Sprintf("%s is required", err.Field))
		case "url":
			errMssgs = append(errMssgs, fmt.Sprintf("%s is not a valid URL", err.Field))
		default:
			errMssgs = append(errMssgs, fmt.Sprintf("%s is invalid", err.Field))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMssgs, ", "),
	}
}
