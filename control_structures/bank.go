package main

import (
	"bank/fileops"
	"fmt"
	"os"
)

const accountFile = "balance.txt"

func main() {

	accountBalance, err := fileops.ReadyBalanceFile(accountFile)
	if err != nil {
		fmt.Println("No money account!")
	}
	for {
		presentOptions()
		var choice int
		fmt.Print("\nYour choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Money account: %v\n", accountBalance)
		case 2:
			var newDeposit float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&newDeposit)

			if newDeposit <= 0 {
				fmt.Println("Invalid amount!")
				continue
				// return - tambem funciona para sair do loop
			} else {
				accountBalance := accountBalance + newDeposit
				fmt.Printf("Accoumt Balance update! %v\n", accountBalance)
				fileops.WriteBalanceToFile(accountFile, accountBalance)
			}
		case 3:
			var withDraw float64
			fmt.Print("withdraw amount: ")
			fmt.Scan(&withDraw)

			if withDraw > accountBalance || withDraw <= 0 {
				fmt.Println("Invalid amount!")
				continue

			} else {
				accountBalance = accountBalance - withDraw
				fmt.Printf("Your withdraw = %v\n", withDraw)
				fmt.Printf("Account Balance = %v\n", accountBalance)
			}
		default:
			fmt.Println("See you later!")
			os.Exit(1)
		}
	}
}
