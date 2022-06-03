package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/iawia002/lux/app"
	"luxapisrv/service"
)

func Run() {
	engine := gin.Default()
	engine.GET("/extract", handleExtract)

	err := engine.Run("0.0.0.0:666")
	if err != nil {
		log.Fatal(err)
	}
}

func handleExtract(c *gin.Context) {
	url := c.Query("url")
	data, err := service.Extract(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, data)
	}
}
