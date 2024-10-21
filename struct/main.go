package main

import (
	"fmt"
	"struct/user"
)

func main() {
	userFirtName := user.GetInfoUser("Please enter your firt name: ")
	userLastName := user.GetInfoUser("Please enter your last name: ")
	userBirthdate := user.GetInfoUser("Please enter your birthdate MM/DD/YYYY: ")

	var appUser *user.User

	appUser, err := user.New(userFirtName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputDetailsUser()
}
