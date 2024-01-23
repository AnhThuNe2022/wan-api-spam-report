package main

import (
	"net/http"
	"os"
	"wan-api-spam-report/Controllers"
	"wan-api-spam-report/Initializers"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			News API
//	@version		1.21.0
//	@description	This is a sample server for News model.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	ThanhPham
//	@contact.url	http://weallnet.com
//	@contact.email	pct1003@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8081
//	@BasePath	/api/News
//	@schemes	http

func init() {
	Initializers.LoadEnvironmentVariables()
	Initializers.ConnectToDB()
}
func main() {

	//Set Gin mode
	mode := os.Getenv("GIN_MODE")
	gin.SetMode(mode)

	//Set log time
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	r := gin.Default()

	News := r.Group("/spamreport")
	{
		News.POST("/save", Controllers.SaveorUpdateSpamReports)

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Wrong API path"})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(); err != nil {
		log.Error().Msg("Failed to start server")
	}
}

//20-12-2023
