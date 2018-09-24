package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	initializeRoutes()
	router.Run()
}

func initializeRoutes() {
	// router.POST("/api", handlePost)
	// router.OPTIONS("/api", )
	router.GET("/", handleBaseGet)
	router.GET("/api", handleGet)
}

func handleBaseGet(c *gin.Context) {
	//
}

func handleGet(c *gin.Context) {
	message, _ := c.GetQuery("m")
	c.String(http.StatusOK, "Get works! you sent: "+message+"\n")
}
