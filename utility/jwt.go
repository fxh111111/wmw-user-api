package utility

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	key    = "this is my jwt key"
	MaxAge = time.Hour
)

type myClaim struct {
	jwt.RegisteredClaims
	ID string `json:"id"`
}

func SignAToken(id string) (token string, err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &myClaim{
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(MaxAge))},
		ID:               id,
	})
	token, err = t.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string) (id string, err error) {
	var t *jwt.Token
	t, err = jwt.ParseWithClaims(token, &myClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return "", err
	}
	mc, ok := t.Claims.(*myClaim)
	if !ok {
		return "", errors.New("token parse failed")
	}
	return mc.ID, nil
}
