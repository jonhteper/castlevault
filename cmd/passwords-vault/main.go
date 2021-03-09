package main

import (
	"flag"
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/jonhteper/castlevault/internal"
	"log"
)

const (
	create = "create"
	get    = "get"
	add    = "add"
)

func printManual() {
	fmt.Println("docs in progress...") // TODO create manual
}

func handleErr(err error) {
	if err != nil {
		log.Fatalf("\nInternal error: %v\n", err)
	}
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		fmt.Println("Error: not enough args")
		printManual()
		return
	}

	switch args[0] {
	case create:
		fmt.Print("passphrase(32 char):")
		//byteKey, err := terminal.ReadPassword(syscall.Stdin)
		byteKey, err := gopass.GetPasswd()
		handleErr(err)

		fmt.Print("\nAdd your first password\n")
		pass, err := internal.NewPassword()
		handleErr(err)

		vault := internal.NewPasswordVault(args[1], string(byteKey))
		err = vault.Add(pass)
		handleErr(err)

		fmt.Println("Password vault created successfully")
	case get:
		if len(args) < 3 {
			fmt.Println("Error: not enough args")
			printManual()
			return
		}

		fmt.Print("passphrase: ")
		byteKey, err := gopass.GetPasswd()
		handleErr(err)

		vault := internal.NewPasswordVault(args[1], string(byteKey))
		err = vault.Open()
		handleErr(err)

		var pass internal.Password
		pass, err = vault.Get(args[2])
		handleErr(err)

		fmt.Printf("\n%v\n", pass.Password)
	case add:
		fmt.Print("passphrase: ")
		byteKey, err := gopass.GetPasswd()
		handleErr(err)

		vault := internal.NewPasswordVault(args[1], string(byteKey))

		err = vault.Open()
		handleErr(err)

		fmt.Println("Add a new password.")
		pass, err := internal.NewPassword()
		handleErr(err)

		err = vault.Add(pass)
		handleErr(err)
		fmt.Println("Password saved")
	default:
		fmt.Println("Error: not correct args")
		printManual()
		return
	}

}
