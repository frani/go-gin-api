package auth

import (
	"fmt"
	"time"

	"github.com/frani/go-gin-api/src/configs"
	"github.com/golang-jwt/jwt"
)

// For HMAC signing method, the key can be any []byte. It is recommended to generate
// a key using crypto/rand or something equivalent. You need the same key for signing
// and validating.
var hmacSecret = []byte(configs.JWT_SECRET)
var expiryTime = configs.JWT_EXPIRY_TIME
var issuer = configs.JWT_ISSUER

func Authenticate(token string) (result *jwt.Token, err error) {

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	result, err = jwt.Parse(token, func(result *jwt.Token) (interface{}, error) {
		if _, ok := result.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", result.Header["alg"])
		}
		return hmacSecret, nil
	})

	if err != nil {
		return nil, err
	}

	return result, err

}

// TODO: add claims keys as parameters
func SignToken() (tokenString string, err error) {

	duration, err := time.ParseDuration(expiryTime)
	if err != nil {
		fmt.Println("Error parsing duration:", err)
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"StandardClaims": jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
			Issuer:    issuer,
		},
	})
	tokenString, err = token.SignedString(hmacSecret)
	return tokenString, err
}
