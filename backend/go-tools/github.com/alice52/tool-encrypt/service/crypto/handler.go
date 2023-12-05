package crypto

import (
	"github.com/alice52/jasypt-go"
	"github.com/alice52/jasypt-go/constant"
	"github.com/alice52/tool-encrypt/model"
	"github.com/gin-gonic/gin"
	"os"
)

// @Summary Decrypt a value
// @Description Decrypts a value using Jasypt
// @Tags Crypto
// @Accept json
// @Produce json
// @Param key query string false "Jasypt encryption key"
// @Param value query string true "Value to decrypt"
// @Success 200 {object} model.R
// @Failure 400 {object} model.R
// @Router /decrypt [get]
func decryptHandler(c *gin.Context) {
	var req model.Request
	if err := c.ShouldBindQuery(&req); err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}

	if len(req.Key) != 0 {
		os.Setenv(constant.JasyptKey, req.Key)
	}

	crypto := jasypt.New()
	encrypt, err := crypto.Decrypt(req.Value)
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}

	model.OkWithData(encrypt, c)
}

// @Summary Encrypt a value
// @Description Encrypts a value using Jasypt
// @Tags Crypto
// @Accept json
// @Produce json
// @Param key query string false "Jasypt encryption key"
// @Param value query string true "Value to encrypt"
// @Success 200 {object} model.R
// @Failure 400 {object} model.R
// @Router /encrypt [get]
func encryptHandler(c *gin.Context) {
	var req model.Request
	if err := c.ShouldBindQuery(&req); err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}

	if len(req.Key) != 0 {
		os.Setenv(constant.JasyptKey, req.Key)
	}

	crypto := jasypt.New()
	encrypt, err := crypto.Encrypt(req.Value)
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}

	model.OkWithData(encrypt, c)
}
