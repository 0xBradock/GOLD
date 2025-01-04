package utils

import (
	"github.com/0xBradock/go-srvr/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	key []byte
	t   *jwt.Token
	s   string
}

func Encode(c *config.Config) (*Token, error) {

	jwt.New(jwt.SigningMethodHS256)

	t := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		jwt.MapClaims{
			"iss": "my-server",
		},
	)

	s, err := t.SignedString(c.JWTSecret)
	if err != nil {
		return nil, err
	}

	return &Token{
		key: []byte(c.JWTSecret),
		t:   t,
		s:   s,
	}, nil

}

func Decode() error {
	return nil
}
