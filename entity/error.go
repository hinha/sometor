package entity

import "fmt"

// ApplicationError cover all application error
type ApplicationError struct {
	Err        []error
	HTTPStatus int
}

type APIError struct {
	Status string `json:"status"`
	Detail string `json:"detail"`
}

// ErrorString return error string as an array of string
func (a *ApplicationError) ErrorString() []APIError {
	var appendError []APIError

	for _, e := range a.Err {
		appendError = append(appendError, APIError{
			Detail: e.Error(),
			Status: fmt.Sprintf("%d", a.HTTPStatus),
		})
	}

	return appendError
}

func (a *ApplicationError) Error() string {
	errString := ""

	for _, e := range a.Err {
		errString = errString + e.Error() + ". "
	}

	return errString
}
