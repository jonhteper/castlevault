package internal

import (
	"encoding/json"
	"errors"
	"os"
)

const dbPermission = 0755

// TODO check permissions
type Password struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type PasswordsVault struct {
	filePath  string
	key       string
	passwords []Password
}

func (p *PasswordsVault) Open() error {
	dataFile, err := os.ReadFile(p.filePath)
	if err != nil {
		return err
	}

	data, err := DecryptAES(dataFile, []byte(p.key))
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &p.passwords)
}

func (p PasswordsVault) Passwords() []Password {
	return p.passwords
}

func (p PasswordsVault) Get(name string) (Password, error) {
	for _, password := range p.passwords {
		if password.Name == name {
			return password, nil
		}
	}

	return Password{}, errors.New("password not be in vault")
}

func (p *PasswordsVault) Add(password Password) error {
	for _, pass := range p.passwords {
		if pass.Name == password.Name {
			return errors.New("the password already exist")
		}
	}

	p.passwords = append(p.passwords, password)

	return p.Save()
}

func (p PasswordsVault) Save() error {
	file, err := os.OpenFile(p.filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, dbPermission)
	if err != nil {
		return err
	}
	defer file.Close()

	err = file.Truncate(0)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(p.passwords, "", "  ")
	if err != nil {
		return err
	}

	m, err := EncryptAES(data, []byte(p.key))
	if err != nil {
		return err
	}

	_, err = file.Write(m)

	return err
}

func NewPasswordVault(filePath, key string) PasswordsVault {
	return PasswordsVault{
		filePath: filePath,
		key:      key,
	}
}
