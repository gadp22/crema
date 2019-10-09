// Copyright 2019 The Crema Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crema

import (
	"time"
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("crema")

// SetSigningKey set
// TO DO : documentaion will be updated soon
func SetSigningKey(key string) {
	mySigningKey = []byte(key)
}

// GenerateJWT generates new JWT token
// TO DO : documentaion will be updated soon
func GenerateJWT(data interface{}, expiration time.Duration) (string, error) {
	Printf("[TOKEN] generating JWT ...")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["data"] = data
	claims["exp"] = time.Now().Add(time.Minute * expiration).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		Printf(err)
		return "", err
	}

	return tokenString, nil
}


// ValidateJWT validates JWT token
// TO DO : documentaion will be updated soon
func ValidateJWT(tokenString string) error {
	Printf("[TOKEN] validating JWT ...")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})

	if err == nil {
		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			Printf("[TOKEN] The token is valid ...")
		} else {
			Printf("[TOKEN] error, the token is invalid or already expired. ")
			err = errors.New("The token is invalid or already expired")
		}
	} else {
		Printf("[TOKEN]" + err.Error())
	}

	return err
}
