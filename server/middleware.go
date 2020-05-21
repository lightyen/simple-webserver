package server

import (
	"log"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
)

/** reponse index.html when 404 not found */
func fallback(filename string, allowAny bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request
		if req.Method != "GET" {
			return
		}
		if strings.Contains(c.Request.Header.Get("Accept"), "text/html") {
			c.File(filename)
			c.Abort()
			return
		}
		if allowAny && strings.Contains(c.Request.Header.Get("Accept"), "*/*") {
			c.File(filename)
			c.Abort()
			return
		}
	}
}

func recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				for i := 3; ; i++ {
					_, file, line, ok := runtime.Caller(i)
					if !ok {
						break
					}
					log.Printf("%s:%d", file, line)
				}
			}
		}()
		c.Next()
	}
}
