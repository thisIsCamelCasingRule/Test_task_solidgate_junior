package models

import "fmt"

type Code string

const (
	CodeValidateCardInvalidYear   Code = "001"
	CodeValidateCardInvalidMonth  Code = "002"
	CodeValidateCardExpiredCard   Code = "003"
	CodeValidateCardInvalidNumber Code = "004"
)

type ValidateCardError struct {
	Code    Code
	Message string
}

func (e *ValidateCardError) Error() string {
	return fmt.Sprintf("code: %s \n message: %s", e.Code, e.Message)
}

func NewValidateCardError(err string, code Code) error {
	return &ValidateCardError{
		Code:    code,
		Message: err,
	}
}
