package middleware

import (
	"fmt"
	"github.com/alice52/tool-encrypt/model"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Recovered from panic: ", err)
				model.FailWithMessage(fmt.Sprintf("Internal Error: %v", err), c)
			}
		}()

		// 继续处理请求
		c.Next()
	}
}
