package signinapple

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"shushutochako/pkg/config"
	"shushutochako/pkg/key"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type ClientSecret struct {
	value string
}

func NewClientSecret(config *config.Config, key *key.Key) (cs *ClientSecret, err error) {
	fmt.Println(config.TeamID)
	fmt.Println(config.ClientID)
	claims := &jwt.StandardClaims{
		Issuer:    config.TeamID,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Unix() + 86400*180,
		Audience:  "https://appleid.apple.com",
		Subject:   config.ClientID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	block, _ := pem.Decode([]byte(key.AppleAuthKey))
	cert, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	tokenString, err := token.SignedString(cert)
	if err != nil {
		return nil, err
	}

	return &ClientSecret{
		value: tokenString,
	}, nil
}
