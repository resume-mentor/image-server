package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.StaticFS("/api/v1/haojianl/img", http.Dir("./images"))

	r.POST("/api/v1/haojianl/upload", uploadImg)

	r.Run(":8081")
}

func uploadImg(c *gin.Context) {
	checkToken := c.PostForm("token")
	if checkToken != "haojianl098" {
		c.String(http.StatusForbidden, "Wrong Token")
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, "Failed to upload image. Error: %s", err.Error())
		return
	}

	err = c.SaveUploadedFile(file, "./images/"+file.Filename)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to upload image. Error: %s", err.Error())
		return
	}

	c.String(http.StatusOK, "Upload file %s success", file.Filename)
}
