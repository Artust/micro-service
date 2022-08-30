package stringutil

import (
	"fmt"
	"testing"
	"time"
)

func TestHash(t *testing.T) {
	password := "Hello"

	hashedPassword, err := HashPass(password)
	if err != nil {
		panic(err)
	}

	fmt.Println(hashedPassword)

	time.Sleep(5 * time.Second)

	// Comparing the password with the hash
	err = CheckPasswordHash(hashedPassword, password)
	fmt.Println(err) // nil means it is a match
}
