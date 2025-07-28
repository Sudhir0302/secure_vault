package routes

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Sudhir0302/secure_vault.git/services/share/models"
	"github.com/Sudhir0302/secure_vault.git/services/share/repo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func Test(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "hello from share test"})
}

// to store the sharefile data
func AddShare(c *gin.Context) {
	type body struct {
		FileId        uuid.UUID `json:"file_id"`
		UserId        uuid.UUID `json:"user_id"`
		Password      string    `json:"password"`
		ExpiryDays    int       `json:"expiry_days"`
		DownloadLimit int       `json:"download_limit"`
	}
	reqbody := &body{}
	err := c.BindJSON(&reqbody)
	if err != nil {
		c.JSON(400, gin.H{"msg": "send crct data"})
		return
	}
	share := &models.Share{}
	share.ID = uuid.New()
	share.FileId = reqbody.FileId
	share.UserId = reqbody.UserId
	nanoid, err := gonanoid.New()
	if err != nil {
		c.JSON(500, gin.H{"msg": "error generating sharelink"})
		return
	}
	share.ShareLink = nanoid
	share.Password = reqbody.Password
	share.ExpiryDays = reqbody.ExpiryDays
	share.DownloadLimit = reqbody.DownloadLimit

	res, err := repo.Create(share)
	if err != nil {
		c.JSON(500, gin.H{"msg": "failed to store in db"})
		return
	}
	c.JSON(201, gin.H{"msg": "success", "data": res})
}

func GetShare(c *gin.Context) {
	sharelink := c.Query("share_link")
	password := c.Query("password")

	reqbody := &models.Share{}
	_, err := repo.GetFile(reqbody, sharelink, password)

	if err != nil {
		c.JSON(404, gin.H{"msg": "link not found"})
		return
	}

	//constructing the storage-service url
	url := fmt.Sprintf("http://localhost:8081/getfile?userid=%s&file_id=%s", reqbody.UserId.String(), reqbody.FileId.String())
	// fmt.Println(url)

	// http.Get(`http://localhost:8081/getfile?userid=2e0dd72b-af5d-4a5d-9ec2-770e1b371f9a&file_name=test`)
	resdata, _ := http.Get(url)
	// fmt.Println(resdata)

	//setting the header from resdata
	c.Header("Content-Disposition", resdata.Header.Get("Content-Disposition"))
	c.Header("Content-Type", resdata.Header.Get("Content-Type"))

	c.Status(http.StatusOK)
	//io.Copy copies data from resdata.Body to actual req of c i.e /getshare endpoint
	io.Copy(c.Writer, resdata.Body)
}
