package util

import "crypto/rand"
import "fmt"

//go:generate counterfeiter ./ CredsGenerator

type CredsGenerator interface {
	Generate() (string, string, error)
}

type credsGenerator struct{}

func NewCredsGenerator() CredsGenerator {
	return &credsGenerator{}
}

func (_ *credsGenerator) Generate() (string, string, error) {
	size := 20
	buffer := make([]byte, size)
	if _, err := rand.Read(buffer); err != nil {
		return "", "", err
	}
	userName := fmt.Sprintf("%x", buffer)
	if _, err := rand.Read(buffer); err != nil {
		return "", "", err
	}
	password := fmt.Sprintf("%x", buffer)

	return userName, password, nil
}
