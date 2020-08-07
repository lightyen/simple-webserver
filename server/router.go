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
	indexCache := cache.CachePage(
		persistence.NewInMemoryStore(0*time.Second),
		0*time.Second,
		static.Serve("/", static.LocalFile(bundleRoot, false)),
	)
	e.GET("/", indexCache)
	e.GET("/index.html", indexCache)
	e.NoRoute(static.Serve("/", static.LocalFile(bundleRoot, false)), fallback(filepath.Join(bundleRoot, "index.html"), true))
	return e
}
