package validators

import (
	"fmt"
	"go-chain-responsabilities/src/errors"
	"strings"
)

type ShortPassword struct {
	Validator
	MinimumLength int
}

func (s *ShortPassword) Initialize(minimumLength int) *ShortPassword {
	s.MinimumLength = minimumLength

	return s
}

func (s *ShortPassword) Check(email, password string) error {
	if len(strings.Split(password, "")) < s.MinimumLength {
		validationError := new(errors.Validation)
		validationError.Message = fmt.Sprintf("Password too short. Minimum required length: %d", s.MinimumLength)

		return validationError
	} else {
		return s.checkNext(email, password)
	}
}
