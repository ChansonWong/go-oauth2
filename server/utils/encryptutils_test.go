package utils

import (
	"fmt"
	"testing"
)

func TestPasswordEncode(t *testing.T) {
	encodePassword, _ := PasswordEncode("admin123", "", 2000)
	fmt.Println(encodePassword)
}

func TestPasswordVerify(t *testing.T) {
	result, _ := PasswordVerify("admin123", "pbkdf2_sha256$2000$Salangid332$3zAP5fqO0XIKAX28gA28YR9Q5DoEIH/+3KYnwmpqMkc=")
	fmt.Println(result)
}
