package security

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"review-server/core"
	"time"
)

type RequestTokenGenerator struct {
	secret string
}

func (r *RequestTokenGenerator) Create(name string, exp time.Duration) (tokenStr string) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(exp).Unix(),
		"iat": time.Now().Unix(),
		"u":   fmt.Sprintf("%x", md5.Sum([]byte(name))),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenStr, err := token.SignedString([]byte(r.secret))

	if err != nil {
		tokenStr = ""
	}

	return
}

var tg *RequestTokenGenerator

func NewRequestTokenGenerator() *RequestTokenGenerator {
	if tg == nil {
		tg = &RequestTokenGenerator{secret: core.GetConfig().App.ReLoginSecret}
	}
	return tg
}
