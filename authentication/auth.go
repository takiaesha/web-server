package authentication

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

var samplekey = []byte("dynamic")

func GenerateJWTKey(SigningKey []byte) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	claims["authorized"] = "true"
	claims["user"] = "Esha"

	tokenString, err := token.SignedString(SigningKey)

	if err != nil {
		fmt.Println("Are you really a Human?!!")
		return "", nil
	}
	return tokenString, nil
}

func VerifyAuth(endpoint func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var sampleJWT = []byte("dynamic")

		bearer := r.Header.Get("Authorization")
		if len(bearer) > 0 {
			bearer = bearer[7:]

			token, err := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
				_, alg := token.Method.(*jwt.SigningMethodHMAC)

				if !alg {
					return nil, fmt.Errorf("invalid Error")
				}
				return sampleJWT, nil
			})

			if err != nil {
				_, err2 := fmt.Fprintf(w, err.Error())
				if err2 != nil {
					return
				}
			}

			if token.Valid {
				endpoint(w, r)
			} else {
				_, err2 := fmt.Fprintf(w, "! You are not authorized")
				if err2 != nil {
					return
				}
			}
		}
	})

}
