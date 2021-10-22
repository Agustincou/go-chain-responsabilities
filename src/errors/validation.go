package errors

import "fmt"

type Validation struct {
	error
	Message string
}

func (l *Validation) Error() string {
	return fmt.Sprintln("Validation error:", l.Message)
}
