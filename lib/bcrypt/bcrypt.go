package bcryptlib

import "golang.org/x/crypto/bcrypt"

type BcryptStruct struct{}

func (r BcryptStruct) Hash(input string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

var BcryptLib = new(BcryptStruct)
