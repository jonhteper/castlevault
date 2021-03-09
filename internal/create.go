package internal

import (
	"errors"
	"fmt"
	"github.com/howeyc/gopass"
)

func NewPassword() (pass Password, err error) {
	fmt.Print("Label: ")
	_, err = fmt.Scan(&pass.Name)
	if err != nil {
		return
	}

	fmt.Print("Password: ")
	//bytesPass, err := terminal.ReadPassword(syscall.Stdin)
	bytesPass, err := gopass.GetPasswd()
	if err != nil {
		return
	}

	pass.Password = string(bytesPass)
	if len(pass.Password) < 8 {
		return Password{}, errors.New("enter password with at least 8 characters")
	}

	return
}
