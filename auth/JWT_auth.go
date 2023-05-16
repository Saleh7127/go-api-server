package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var SecretKey = []byte("CHAGOL")

func GenerateJWT(SigningKey []byte) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claimInfo := token.Claims.(jwt.MapClaims)
	claimInfo["authorized"] = true
	claimInfo["users"] = "someone"
	claimInfo["exp"] = time.Now().Add(time.Hour * 3).Unix()
	tokenString, err := token.SignedString(SigningKey)

	if err != nil {
		fmt.Printf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		if len(bearer) > 0 {
			bearer = bearer[7:]

			token, err := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an Error")
				}
				return SecretKey, nil
			})
			if err != nil {
				fmt.Fprint(w, err.Error())
			}
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}

	})
}
