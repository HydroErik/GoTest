package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPass(old_pass string) (string, error) {
	//encoded := base64.StdEncoding.EncodeToString([]byte(old_pass))
	encoded, err := bcrypt.GenerateFromPassword([]byte(old_pass), bcrypt.DefaultCost)
	return string(encoded), err
}

func main() {
	raw := "Einstein4"

	encryp, err := EncryptPass(raw)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Raw Password: %v\nEncrypted: %v\n", raw, encryp)

}
