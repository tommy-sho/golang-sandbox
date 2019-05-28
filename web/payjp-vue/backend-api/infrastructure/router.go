package infrastructure

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tommy-sho/golang-sandbox/web/payjp-vue/backend-api/handler"
)

// Router - router api server
var Router *gin.Engine

func init() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("CLIENT_CORS_ADDR")},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	router.GET("/api/v1/items", func(c *gin.Context) { handler.GetLists(c) })
	router.GET("/api/v1/items/:id", func(c *gin.Context) { handler.GetItem(c) })
	router.POST("/api/v1/charge/items/:id", func(c *gin.Context) { handler.Charge(c) })

	Router = router
}
