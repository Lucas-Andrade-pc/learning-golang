package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountFile = "balance.txt"

func readyBalanceFile() (float64, error) {
	data, err := os.ReadFile(accountFile)
	if err != nil {
		fmt.Println(err)
		return 1000, errors.New("failed open file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed convert string -> float64")
	}
	return balance, nil

}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountFile, []byte(balanceText), 0644)
}

func main() {

	accountBalance, err := readyBalanceFile()
	if err != nil {
		fmt.Println(err)
		//os.Exit(1)
		panic("Can't continue sorry!")
	}
	for {
		fmt.Println("Welcome to go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
				return
			} else {
				accountBalance := accountBalance + newDeposit
				fmt.Printf("Accoumt Balance update! %v\n", accountBalance)
				writeBalanceToFile(accountBalance)
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
