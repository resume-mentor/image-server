package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.StaticFS("/api/v1/haojianl/img", http.Dir("./images"))
	r.Run(":8081")
}
