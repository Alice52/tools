package config

import (
	"github.com/alice52/tool-encrypt/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func Swag(r *gin.Engine) {
	_ = docs.SwaggerInfo.BasePath

	// Swagger documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Redirect /swagger to /swagger/index.html
	swagRedirectHandler := func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	}
	r.GET("/swagger", swagRedirectHandler)
	r.GET("/api", swagRedirectHandler)
	r.GET("/", swagRedirectHandler)
}
