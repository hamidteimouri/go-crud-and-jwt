package auth

import (
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func CreateToken(userId uint32) (string, error) {

	/* prepare claim */
	claims := jwt.MapClaims{}
	claims["authorize"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	/* create token */
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	/* sign and convert to string */
	return token.SignedString(os.Getenv("API_SECRET"))
}

func ExtractToken(request *http.Request) string {
	/* get token from query string */
	keys := request.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	/* get token from header */
	bearerToken := request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(request *http.Request) (uint32, error) {
	tokenString := ExtractToken(request)

	/*  parse token */
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}

	/* get user id from claim */
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, err := strconv.ParseInt(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)

		if err != nil {
			return 0, err
		}
		return uint32(userId), nil
	}

	return 0, nil

}

func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(b)
}

func TokenValid(request *http.Request) error {
	tokenString := ExtractToken(request)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method : %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}
