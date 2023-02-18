package services

import (
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"math/rand"
	"time"
)

const (
	codeLegnth   = 4
	tokenLength  = 32
	codeCharset  = "0123456789"
	tokenCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func GenerateRandomCode() string {
	return stringWithCharset(codeLegnth, codeCharset)
}

func GenerateToken() *entities.Tokens {
	token := stringWithCharset(tokenLength, tokenCharset)
	return &entities.Tokens{
		AccessToken:  string(token[:len(token)/2]),
		RefreshToken: string(token[len(token)/2:]),
	}
}

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
