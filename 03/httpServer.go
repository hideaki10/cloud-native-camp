package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// setting log format
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("Client ip is %s - Status code is %d\n",
			param.ClientIP,
			param.StatusCode,
		)
	}))

	r.GET("/", index)

	r.GET("/healthz", healthz)

	r.Run(":8080")

}

func writerHeaderInfoToResponse(c *gin.Context) {
	headInfo := c.Request.Header
	for key, _ := range headInfo {
		c.Writer.Header().Add(key, c.Request.Header.Get(key))
	}
}

func healthz(c *gin.Context) {
	writerHeaderInfoToResponse(c)
	getVersionInfo(c)
	c.Writer.WriteHeader(200)
}

func index(c *gin.Context) {
	writerHeaderInfoToResponse(c)
	getVersionInfo(c)
	io.WriteString(c.Writer, "index")
}

func getVersionInfo(c *gin.Context) {
	version := os.Getenv("VERSION")
	c.Writer.Header().Add("VERSION", version)
}
