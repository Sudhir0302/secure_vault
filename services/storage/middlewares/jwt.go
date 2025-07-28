package middlewares

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		var jwtkey = []byte(os.Getenv("JWT_KEY"))
		header := c.Request.Header.Get("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.JSON(400, gin.H{"msg": "Authorization header missing"})
			c.Abort() //prevents pending hanlders from being called
			return
		}
		// fmt.Println(header)

		tokenstr := strings.TrimPrefix(header, "Bearer ")
		// fmt.Println(tokenstr)

		token, err := jwt.Parse(tokenstr, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return jwtkey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"msg": "token expired"})
			c.Abort()
			return
		}
		c.Next()
	}
}
