package main

import (
	"gin-cloudinary-api/controllers" //add this

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//add
	router.POST("/file", controllers.FileUpload())
	router.POST("/remote", controllers.RemoteUpload())

	router.Run("localhost:6000")
}
