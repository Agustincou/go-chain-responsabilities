package validators

type Void struct {
	validator
}

func (e *Void) Check(email, password string) error {
		return nil
}
