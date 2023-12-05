package main

import (
	"fmt"
	"github.com/alice52/tool-encrypt/config"
	"github.com/alice52/tool-encrypt/routers"
	"github.com/alice52/tool-encrypt/service/crypto"
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

// @host localhost:8080
// @BasePath
func main() {

	r := routers.Include(crypto.Routers)
	r.Use(config.Cors())
	config.Swag(r)

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
