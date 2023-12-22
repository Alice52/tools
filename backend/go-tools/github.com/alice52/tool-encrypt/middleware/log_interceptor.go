package middleware

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
)

func LogInterceptor(c *gin.Context) {
	id, _ := uid()
	fmt.Printf("Request received: "+
		"\n req-id: %s"+
		"\n method: %s"+
		"\n path: %s"+
		"\n query: %s"+
		// "\n headers: %s"+
		"\n body: %s\n\n",
		id, c.Request.Method, c.Request.URL.Path, c.Request.URL.RawQuery /*logHeader(c),*/, logReqBody(c))

	wb := &bodyWriter{
		body:           &bytes.Buffer{},
		ResponseWriter: c.Writer,
	}
	c.Writer = wb

	c.Next()

	fmt.Printf("Response completed: "+
		"\n req-id: %s"+
		//"\n headers: %s"+
		"\n body: %s\n\n",
		id /*logHeader(c),*/, logRespBody(wb))
}

func uid() (string, error) {
	uuidBytes := make([]byte, 16)
	_, err := rand.Read(uuidBytes)
	if err != nil {
		return "", err
	}

	uuidBytes[6] = (uuidBytes[6] & 0x0F) | 0x40

	uuidBytes[8] = (uuidBytes[8] & 0x3F) | 0x80

	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", uuidBytes[0:4],
		uuidBytes[4:6], uuidBytes[6:8], uuidBytes[8:10], uuidBytes[10:])
	return uuid, nil
}

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r bodyWriter) Write(b []byte) (int, error) {
	return r.body.Write(b)
}

func logRespBody(wb *bodyWriter) string {
	originBytes := wb.body
	wb.ResponseWriter.Write(originBytes.Bytes())

	return originBytes.String()
}

func logReqBody(c *gin.Context) string {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}

	// 将请求体重新写回到 Request.Body 中
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	return string(body)
}

func logRespHeader(c *gin.Context) map[string]string {

	headers := make(map[string]string)
	for key, values := range c.Writer.Header() {
		headers[key] = values[0]
	}

	return headers
}
func logHeader(c *gin.Context) map[string]string {

	headers := make(map[string]string)
	for key, values := range c.Request.Header {
		headers[key] = values[0]
	}

	return headers
}

func mapToString(m map[string]string) string {
	result := "{"
	for key, value := range m {
		result += fmt.Sprintf("%s: %s, ", key, value)
	}
	result += "}"
	return result
}
