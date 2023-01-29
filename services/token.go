package services

import (
	"douyin/model"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

var TokenService = newTokenService()

func newTokenService() *tokenService {
	return &tokenService{}
}

type tokenService struct {
}

var mySecret = []byte("jwt")

func (s *tokenService) GenToken(userID int64) (string, error) {
	c := model.MyClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(999 * time.Hour).Unix(),
			Issuer: "jwt",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(mySecret)
}

// err 可以更详细的处理，所以保留
func (s *tokenService) ParseToken(tokenString string) (*model.MyClaims, error) {
	var mc = new(model.MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(*jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

func (s *tokenService) RefreshToken(aToken string) (string, error) {
	claims := &model.MyClaims{}
	tkn, err := jwt.ParseWithClaims(aToken, claims, func(*jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil || !tkn.Valid {
		return "", err
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}