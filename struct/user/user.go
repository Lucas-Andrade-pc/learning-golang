package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firtName  string
	lastName  string
	Birthdate string
	createdAt time.Time
}

func (u *User) OutputDetailsUser() {
	fmt.Println(u.firtName, u.lastName, u.Birthdate, u.createdAt)
}

func GetInfoUser(data string) string {
	fmt.Print(data)
	var info string
	fmt.Scanln(&info)

	return info

}
func New(fn, ln, bd string) (*User, error) { // construtor
	if fn == "" || ln == "" || bd == "" {
		return nil, errors.New("firt name, last name , birthdate is required")
	}
	return &User{
		firtName:  fn,
		lastName:  ln,
		Birthdate: bd,
	}, nil
}
