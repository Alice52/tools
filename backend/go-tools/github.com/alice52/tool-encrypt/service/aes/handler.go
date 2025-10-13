package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/alice52/tool-encrypt/model"
	"github.com/gin-gonic/gin"
	"os"
)

const (
	AesKey = "AES_KEY"
	AesIv  = "AES_IV"
)

// @Summary Decrypt a value by aes
// @Description Decrypts a value by aes
// @Tags Aes
// @Accept json
// @Produce json
// @Param key query string false "aes key"
// @Param iv query string false "aes iv"
// @Param value query string true "Value to decrypt"
// @Success 200 {object} model.R
// @Failure 400 {object} model.R
// @Router /aes-decrypt [get]
func decryptHandler(c *gin.Context) {
	var req model.AesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}

	if len(req.Key) == 0 {
		req.Key = os.Getenv(AesKey)
	}
	if len(req.Iv) == 0 {
		req.Iv = os.Getenv(AesIv)
	}

	encrypted, err := base64.StdEncoding.DecodeString(req.Value)
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}

	block, err := aes.NewCipher([]byte(req.Key))
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}

	mode := cipher.NewCBCDecrypter(block, []byte(req.Iv))
	decrypted := make([]byte, len(encrypted))
	mode.CryptBlocks(decrypted, encrypted)

	padding := decrypted[len(decrypted)-1]
	model.OkWithData(string(decrypted[:len(decrypted)-int(padding)]), c)
}

// @Summary Encrypt a value by aes
// @Description Encrypts a value by aes
// @Tags Aes
// @Accept json
// @Produce json
// @Param key query string false "aes key"
// @Param iv query string false "aes iv"
// @Param value query string true "Value to encrypt"
// @Success 200 {object} model.R
// @Failure 400 {object} model.R
// @Router /aes-encrypt [get]
func encryptHandler(c *gin.Context) {
	var req model.AesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}

	if len(req.Key) == 0 {
		req.Key = os.Getenv(AesKey)
	}
	if len(req.Iv) == 0 {
		req.Iv = os.Getenv(AesIv)
	}

	block, err := aes.NewCipher([]byte(req.Key))
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}

	// 处理数据填充（根据算法块大小）
	blockSize := block.BlockSize()
	padding := blockSize - len(req.Value)%blockSize
	padText := make([]byte, padding)
	for i := range padText {
		padText[i] = byte(padding)
	}
	dataBytes := append([]byte(req.Value), padText...)

	mode := cipher.NewCBCEncrypter(block, []byte(req.Iv))
	ciphertext := make([]byte, len(dataBytes))
	mode.CryptBlocks(ciphertext, dataBytes)

	model.OkWithData(base64.StdEncoding.EncodeToString(ciphertext), c)
}
