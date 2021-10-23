package validators

type Interface interface {
	LinkWith(validatorInterface Interface) Interface
	Check(email, password string) error
	checkNext(email, password string) error
}

type chain struct {
	nextValidator Interface
}

// Validator - For simulate "abstract" class when other structs compose with him
type Validator struct {
	Interface
	chain
}

func (v *Validator) Check(email, password string) error {
	return nil
}

func (v *Validator) checkNext(email, password string) error {
	if v.nextValidator != nil {
		return v.nextValidator.Check(email, password)
	} else {
		return nil
	}
}

func (v *Validator) LinkWith(validatorInterface Interface) Interface {
	v.nextValidator = validatorInterface
	return validatorInterface
}
