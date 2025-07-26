package routes

import (
	"github.com/Sudhir0302/secure_vault.git/services/auth/models"
	"github.com/Sudhir0302/secure_vault.git/services/auth/repo"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.JSON(200, map[string]string{"msg": "hello from test"})
}

func Signup(c *gin.Context) {

	// type body struct {
	// 	Username string
	// 	Password string
	// }

	// user := &body{}
	user := &models.User{}
	c.Bind(user)
	_, err := repo.Create(user)
	if err != nil {
		c.JSON(404, `{"msg":"user not created"}`)
	}
	// fmt.Println(user)
	c.JSON(201, user)
}
