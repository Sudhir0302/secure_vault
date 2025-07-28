package routes

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Sudhir0302/secure_vault/services/storage/models"
	"github.com/Sudhir0302/secure_vault/services/storage/repo"
	"github.com/Sudhir0302/secure_vault/services/storage/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Test(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "hello from storage test"})
}

func UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil || file == nil || header == nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "file not found"})
		return
	}
	defer file.Close()

	userid := c.DefaultPostForm("userid", "")
	filename := c.DefaultPostForm("filename", "")
	filesize := header.Size
	mimetype := header.Header.Get("Content-Type")

	data, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "failed to read file"})
		return
	}

	enc := utils.Encrypt(data)

	store := &models.Storage{
		ID:            uuid.New(),
		Userid:        uuid.MustParse(userid),
		FileName:      filename,
		FileSize:      filesize,
		Mime_type:     mimetype,
		EncryptedData: enc,
	}

	res, err := repo.Store(store)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "failed to store in db"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "file uploaded", "data": res})
}

func GetFile(c *gin.Context) {
	type body struct {
		Userid string `json:"userid"`
		FileId string `json:"file_id"`
	}
	reqdata := &body{
		Userid: c.Query("userid"),
		FileId: c.Query("file_id"),
	}
	fmt.Println(reqdata)

	res, err := repo.GetFile(reqdata.Userid, reqdata.FileId)
	if err != nil || res == nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "file not found"})
		return
	}

	dec := utils.Decrypt(res.EncryptedData)
	if dec == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "failed to decrypt file"})
		return
	}

	filename := res.FileName
	mime := res.Mime_type

	//add file extension
	if mime == "application/pdf" {
		filename += ".pdf"
	} else if mime == "text/plain" {
		filename += ".txt"
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))

	//c.Data means sending a raw binary or byte slice response to the client
	c.Data(http.StatusOK, res.Mime_type, dec)
}
