package utils

import (
	"fmt"
	"testing"
)

func TestNewSecret(t *testing.T) {
	fmt.Println(NewSecret())
}

func TestCalculate2FACode(t *testing.T) {
	secret := "FBSKBJZEJZGH2XCBVCBUH7APBQ"
	fmt.Println(Calculate2FACode(secret))
}

func TestVerify2FACode(t *testing.T) {
	secret := "FBSKBJZEJZGH2XCBVCBUH7APBQ"
	b, err := Verify2FACode(secret, "056692")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(b)
}
