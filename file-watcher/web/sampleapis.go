package web

import (
	"file-watcher/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var apiRouter *gin.Engine
var multipleFactor int

func init() {
	Initialize()
	apiRouter = gin.Default()
	// watchAndReload()
}

func Initialize() {
	appConfig := config.LoadConfig()
	multipleFactor = appConfig.Web.Factor
}

func SampleApi() {
	apiRouter.GET("/multiple/:number", func(c *gin.Context) {
		number, err := strconv.Atoi(c.Param("number"))
		if err != nil {
			c.String(http.StatusInternalServerError, "Error while parsing input :"+err.Error())
		}
		c.String(http.StatusOK, strconv.Itoa(number*multipleFactor))
	})
	apiRouter.Run("localhost:8080")
}
