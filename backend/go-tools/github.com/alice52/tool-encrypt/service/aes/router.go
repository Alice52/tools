package aes

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {

	e.GET("/aes-decrypt", decryptHandler)

	e.GET("/aes-encrypt", encryptHandler)
}
