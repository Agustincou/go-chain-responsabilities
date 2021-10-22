package errors

import "fmt"

type Login struct {
	error
	Message string
}

func (l *Login) Error() string {
	return fmt.Sprintln("Login error:", l.Message)
}
