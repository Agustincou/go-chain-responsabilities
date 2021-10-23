package checks

import (
	"go-chain-responsabilities/src/server/checks/validators"
)

type Login struct {
	validators.Validator
}

func (e *Login) Check(email, password string) error {
	loginChecks := new(validators.Email)

	return loginChecks.Check(email, password)
}
