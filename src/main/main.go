package main

import (
	"bufio"
	"fmt"
	server2 "go-chain-responsabilities/src/server"
	"go-chain-responsabilities/src/server/validators"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	server := new(server2.Server).Initialize()

	registerChecks := new(validators.Email)
	registerChecks.
		LinkWith(new(validators.ShortPassword).Initialize(4)).
		LinkWith(new(validators.WeakPassword))

	loginChecks := new(validators.Email)

	server.RegisterInputChecks = registerChecks
	server.LoginInputChecks = loginChecks

	for {
		inputString := requestUserAction()
		var userEmail string
		var userPassword string

		switch inputString {
		case "LOGIN":
			userEmail = requestUserEmail()
			userPassword = requestUserPassword()
			if loginError := server.Login(userEmail, userPassword); loginError != nil {
				fmt.Println(loginError.Error())
				break
			}
			return
		case "REGISTER":
			userEmail = requestUserEmail()
			userPassword = requestUserPassword()
			if registerError := server.Register(userEmail, userPassword); registerError != nil {
				fmt.Println(registerError.Error())
				break
			}
		}
	}
}

func requestUserAction() string {
	fmt.Println("Ingrese LOGIN si desea loguearse o ingrese REGISTER si desea registrarse")
	text, _ := reader.ReadString('\n')
	return strings.Replace(text, "\n", "", -1)
}

func requestUserEmail() string {
	fmt.Println("Ingrese su email")
	text, _ := reader.ReadString('\n')
	return strings.Replace(text, "\n", "", -1)
}

func requestUserPassword() string {
	fmt.Println("Ingrese su contraseña")
	text, _ := reader.ReadString('\n')
	return strings.Replace(text, "\n", "", -1)
}
