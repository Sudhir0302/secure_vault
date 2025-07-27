package routes

import (
	"github.com/Sudhir0302/secure_vault.git/services/auth/models"
	"github.com/Sudhir0302/secure_vault.git/services/auth/repo"
	"github.com/Sudhir0302/secure_vault.git/services/auth/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

	user.Id = uuid.New() //create uuid
	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashed)

	_, err := repo.Create(user)
	if err != nil {
		c.JSON(404, map[string]string{"msg": "user not created"})
	}
	// fmt.Println(user)
	c.JSON(201, user)
}

func Signin(c *gin.Context) {

	type body struct {
		Email    string `json:"email" binding:"required"` //binding add input validation and it's checked with c.BindJSON()
		Password string `json:"password" binding:"required"`
	}
	user := body{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"status": "error", "msg": "invalid req body"})
		return
	}
	// fmt.Println(user)

	res, _ := repo.FindUser(user.Email)
	// fmt.Println(res)
	//compare password
	ch := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(user.Password))

	if ch != nil {
		c.JSON(404, map[string]string{"msg": "password mismatch"})
		return
	}

	token, _ := utils.Generatetoken(res.Email) //generate jwt token

	c.JSON(200, gin.H{"status": "success", "token": token})
}
