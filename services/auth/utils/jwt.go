package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Generatetoken(email string) (string, error) {
	var jwtkey = []byte(os.Getenv("JWT_KEY"))
	// fmt.Println(jwtkey)

	//payload data
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
		//unix : a number(seconds) that calculated from the time since jan 1 1970 from 00:00:00 , eg:1722028800.
		// It's like a stopwatch counting seconds since 1970.
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //generate the token with HS256 algo
	return token.SignedString(jwtkey)
}
