package main

import (
	"fmt"
	"github.com/alice52/tool-encrypt/config"
	"github.com/alice52/tool-encrypt/middleware"
	"github.com/alice52/tool-encrypt/routers"
	"github.com/alice52/tool-encrypt/service/crypto"
	"github.com/gin-gonic/gin"
)

// @title Swagger Crypto API
// @version 1.0
// @description This is a crypto server by golang.
// @termsOfService https://github.com/alice52

// @contact.name alice52
// @contact.url https://github.com/alice52/tools
// @contact.email zzhang_xz@163.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath
func main() {
	// 1. get Engine
	r := gin.Default()

	// 2. register swag, which before middleware, so mw will not enable for swag
	config.Swag(r)

	// 3. register middleware
	r.Use(middleware.Cors())
	r.Use(middleware.LogInterceptor)
	r.Use(middleware.Recovery()) // recovery by custom code

	// 4. register router
	routers.Include(r, crypto.Routers)

	// 5. run container
	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
