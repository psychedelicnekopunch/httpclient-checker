
package infrastructure

import (
	"github.com/psychedelicnekopunch/httpclient-checker/src/interfaces/controllers"
	"github.com/gin-gonic/gin"
)


type Routing struct {
	Port string
	Gin *gin.Engine
}


func NewRouting(config *Config) *Routing {
	return &Routing{
		Port: config.Port,
		Gin: gin.Default(),
	}
}


func (r *Routing) Start() {

	indexController := controllers.NewIndexController()
	uploadsController := controllers.NewUploadsController()

	r.Gin.GET("/", indexController.Get)
	r.Gin.POST("/", indexController.Post)
	r.Gin.PUT("/", indexController.Put)
	r.Gin.PATCH("/", indexController.Patch)
	r.Gin.DELETE("/", indexController.Delete)

	r.Gin.POST("/uploads", uploadsController.Post)

	r.Gin.Run(r.Port)
}
