package validators

type Interface interface {
	LinkWith(validatorInterface Interface) Interface
	Check(email, password string) error
	CheckNext(email, password string) error
}

type chain struct {
	nextValidator Interface
}

type validator struct {
	Interface
	chain
}

func (v *validator) Check(email, password string) error {
	return nil
}

func (v *validator) CheckNext(email, password string) error {
	if v.nextValidator != nil {
		return v.nextValidator.Check(email, password)
	} else {
		return nil
	}
}

func (v *validator) LinkWith(validatorInterface Interface) Interface {
	v.nextValidator = validatorInterface
	return validatorInterface
}
