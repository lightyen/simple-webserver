package server

import (
	"path/filepath"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func NewRouter(bundleRoot string) *gin.Engine {
	e := gin.New()
	e.Use(recovery())
	e.Use(gzip.Gzip(gzip.DefaultCompression))

	store := persistence.NewInMemoryStore(time.Second)
	serve := cache.CachePage(store, time.Minute, static.Serve("/", static.LocalFile(bundleRoot, false)))

	webFallback := cache.CachePage(store, time.Minute, fallback(filepath.Join(bundleRoot, "index.html"), true))
	e.NoRoute(serve, webFallback)

	apiFallback := fallback(filepath.Join(bundleRoot, "index.html"), false)
	apis := e.Group("/apis", apiFallback)
	{
		apis.GET("/hello", func(c *gin.Context) {
			c.JSON(200, struct{}{})
		})
	}

	return e
}
