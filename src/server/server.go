package server

import (
	"fmt"
	"go-chain-responsabilities/src/errors"
	"go-chain-responsabilities/src/server/validators"
)

type Server struct {
	registeredUsers     map[string]string
	LoginInputChecks    validators.Interface
	RegisterInputChecks validators.Interface
}

func (s *Server) Initialize() *Server{
	s.registeredUsers = make(map[string]string)
	s.LoginInputChecks = new(validators.Void)
	s.RegisterInputChecks = new(validators.Void)

	return s
}

func (s *Server) Login(email, password string) error {
	var errorToReturn error

	if errorToReturn = s.LoginInputChecks.Check(email, password); errorToReturn == nil {
		if s.registeredUsers[email] == password {
			fmt.Println("Login success!")
		} else {
			fmt.Println("Invalid credentials!")
			loginError := new(errors.Login)
			loginError.Message = "Invalid credentials!"
			errorToReturn = loginError
		}
	}

	return errorToReturn
}

func (s *Server) Register(email, password string) error {
	var errorToReturn error

	if errorToReturn = s.RegisterInputChecks.Check(email, password); errorToReturn == nil {
		s.registeredUsers[email] = password
	}

	return errorToReturn
}
