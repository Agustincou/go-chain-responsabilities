package validators

import (
	"go-chain-responsabilities/src/errors"
	"strings"
)

type WeakPassword struct {
	Validator
}

func (e *WeakPassword) Check(email, password string) error {
	if strings.Contains(password, "1234") {
		validationError := new(errors.Validation)
		validationError.Message = "Too weak password. Be more creative!"

		return validationError
	} else {
		return e.checkNext(email, password)
	}
}
