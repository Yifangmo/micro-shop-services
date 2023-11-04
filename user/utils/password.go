package utils

import (
	"crypto/sha512"
	"fmt"
	"strings"

	"github.com/anaskhan96/go-password-encoder"
)

var (
	EncryptPasswordOptions = &password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: sha512.New,
	}
)

func GenStorePassword(pwd string) string {
	salt, encodedPwd := password.Encode(pwd, EncryptPasswordOptions)
	return fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
}

func GetPasswordDetails(storePwd string) (salt string, encodedPwd string) {
	list := strings.Split(storePwd, "$")
	salt, encodedPwd = list[2], list[3]
	return
}

func VerifyPassword(rawPwd, storePwd string) bool {
	salt, encodedPwd := GetPasswordDetails(storePwd)
	return password.Verify(rawPwd, salt, encodedPwd, EncryptPasswordOptions)
}
