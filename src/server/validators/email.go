package validators

import (
	"go-chain-responsabilities/src/errors"
	"strings"
)

type Email struct {
	validator
}

func (e *Email) Check(email, password string) error {
	if !strings.Contains(email, "@") {
		validationError := new(errors.Validation)
		validationError.Message = "Invalid email format"

		return validationError
	} else {
		return e.CheckNext(email, password)
	}
}
