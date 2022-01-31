package jwt

import (
	"log"
	"time"

	"github.com/El-Hendawy/gograph/graph/model"
	"github.com/dgrijalva/jwt-go"
)

// secret key being used to sign tokens
var (
	SecretKey = []byte("secret")
)

// GenerateToken generates a jwt token and assign a username to it's claims and return it
func GenerateToken(data *model.User) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)
	/* Set token claims */
	claims["email"] = data.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating key")
		//return "", err
	}
	return tokenString, err
}

// ParseToken parses a jwt token and returns the username in it's claims
func ParseToken(tokenStr string) (string, error) {
	//fmt.Println(tokenStr)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		return email, nil
	} else {
		return "", err
	}
}
