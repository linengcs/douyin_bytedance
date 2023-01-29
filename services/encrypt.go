package services

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

var EncryptService = newEncryptService()

func newEncryptService() *encryptService {
	return &encryptService{}
}

type encryptService struct {
}

func (s *encryptService) hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func (s *encryptService) comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
