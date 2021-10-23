package checks

import (
	"go-chain-responsabilities/src/server/checks/validators"
)

type Registration struct {
	validators.Validator
}

func (e *Registration) Check(email, password string) error {
	registrationChecks := new(validators.Email)
	registrationChecks.
		LinkWith(new(validators.ShortPassword).Initialize(4)).
		LinkWith(new(validators.WeakPassword))

	return registrationChecks.Check(email, password)
}
