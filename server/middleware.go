package server

import (
	"log"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// reponse index.html when 404 not found
func fallback(filename string, allowAny bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request
		if req.Method != "GET" {
			return
		}
		if strings.Contains(c.Request.Header.Get("Accept"), "text/html") {
			c.Header("Cache-Control", "no-store")
			c.File(filename)
			c.Abort()
			return
		}
		if allowAny && strings.Contains(c.Request.Header.Get("Accept"), "*/*") {
			c.Header("Cache-Control", "no-store")
			c.File(filename)
			c.Abort()
			return
		}
	}
}

func noCacheIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "GET" {
			return
		}
		if c.Request.URL.Path == "/" || c.Request.URL.Path == "/index.html" {
			c.Header("Cache-Control", "no-store")
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

func simpleHandles(bundleRoot string) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		noCacheIndex(),
		cache.CachePage(
			persistence.NewInMemoryStore(5*time.Second),
			time.Minute,
			static.Serve("/", static.LocalFile(bundleRoot, false)),
		),
		fallback(filepath.Join(bundleRoot, "index.html"), true),
	}
}
