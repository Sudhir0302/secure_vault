package routes

import (
	"fmt"
	"io"

	"github.com/Sudhir0302/secure_vault.git/services/storage/utils"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "hello from storage test"})
}

func UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	// filetype := header.Header["Content-Type"]
	// fmt.Printf("%T", filetype)

	if err != nil || file == nil || header == nil {
		c.JSON(400, gin.H{"msg": "file not found"})
		return
	}
	defer file.Close()
	data, _ := io.ReadAll(file)
	enc := utils.Encrypt(data)
	fmt.Println(enc)
}
