package main

import (
	"net/http"

	kernel "welshacademy/src"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	kernel.Boot()

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welsh Academy api")
	})

	http.ListenAndServe(":80", router)
}
