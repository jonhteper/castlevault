package main

import (
	"flag"
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/jonhteper/castlevault/internal"
	"log"
	"os"
	"path"
)

const (
	createCommand   = "create"
	getCommand      = "get"
	addCommand      = "add"
	manCommand      = "man"
	commandFile     = "manual_commands.txt"
	descriptionFile = "manual_description.txt"
)

func printManual() {
	binDir, err := os.Executable()
	handleErr(err)
	manPath := path.Join(path.Dir(binDir), "manual")

	dataDescription, err := os.ReadFile(path.Join(manPath, descriptionFile))
	handleErr(err)

	dataCommands, err := os.ReadFile(path.Join(manPath, commandFile))
	handleErr(err)

	fmt.Printf("%v\n%v", string(dataDescription), string(dataCommands))
}

func printCommandManual() {
	binDir, err := os.Executable()
	handleErr(err)
	manPath := path.Join(path.Dir(binDir), "manual")

	data, err := os.ReadFile(path.Join(manPath, commandFile))
	handleErr(err)

	fmt.Println(string(data))
}

func handleErr(err error) {
	if err != nil {
		log.Fatalf("\nInternal error: %v\n", err)
	}
}

func main() {
	flag.Parse()
	args := flag.Args()

	switch args[0] {
	case createCommand:
		if len(args) < 2 {
			fmt.Println("Error: not enough args")
			printCommandManual()
			return
		}

		fmt.Print("passphrase(32 char):")
		byteKey, err := gopass.GetPasswd()
		handleErr(err)

		fmt.Print("\nAdd your first password\n")
		pass, err := internal.NewPassword()
		handleErr(err)

		vault := internal.NewPasswordVault(args[1], string(byteKey))
		err = vault.Add(pass)
		handleErr(err)

		fmt.Println("Password vault created successfully")
	case getCommand:
		if len(args) < 3 {
			fmt.Println("Error: not enough args")
			printCommandManual()
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
	case addCommand:
		if len(args) < 2 {
			fmt.Println("Error: not enough args")
			printCommandManual()
			return
		}
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
	case manCommand:
		printManual()
	default:
		fmt.Println("Error: not correct args")
		printCommandManual()
		return
	}
}
