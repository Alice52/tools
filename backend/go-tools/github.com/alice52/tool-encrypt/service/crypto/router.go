package crypto

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {

	e.GET("/decrypt", decryptHandler)

	e.GET("/encrypt", encryptHandler)
}
