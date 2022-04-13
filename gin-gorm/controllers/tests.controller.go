package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Multipart(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}

func UpFile(c *gin.Context) {
	//single file
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	log.Println(os.Getwd())

	//upload the file to specific dst
	if err := c.SaveUploadedFile(file, "./upload/"+file.Filename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func ReadFile(c *gin.Context) {
	f, err := os.Open("./upload/" + c.Param("file"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
		return
	}
	defer f.Close()
	fi, _ := f.Stat()

	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="` + fi.Name() + `"`,
	}

	c.DataFromReader(http.StatusOK, fi.Size(), fi.Mode().Type().String(), f, extraHeaders)
}
