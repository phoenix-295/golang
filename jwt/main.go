package main

import (
	"fmt"
	"time"

	"encoding/base64"

	"github.com/golang-jwt/jwt"
)

var hmacSampleSecret = []byte("secret_key")

func main() {
	ts := encode()

	decode(ts)
}

// Encode user data in json token
func encode() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   "1",
		"user_name": "eric_vice@gmail.com",
		"time":      time.Now(),
		"isSuper":   true,
	})
	tokenString, _ := token.SignedString(hmacSampleSecret)

	fmt.Println("Token", tokenString)

	algo, _ := base64.StdEncoding.DecodeString("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9")
	payload, _ := base64.RawStdEncoding.DecodeString("eyJpZCI6MzIsImV4cCI6MTY1MTA1MDIyMiwiaWF0IjoxNjUxMDUwMTYyfQ")

	fmt.Printf("Decoded text: %s\n", algo)
	fmt.Printf("Payload %s\n", payload)
	return tokenString
}

// Decode JSON data in the token
func decode(tokenString string) {
	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("User ID:", claims["user_id"], "\nUser Name: ", claims["user_name"], "\nIs Super: ", claims["isSuper"], "\nTime:", claims["time"])
	} else {
		fmt.Println(err)
	}
}
