package server

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func NewRouter(bundleRoot string) *gin.Engine {
	e := gin.New()
	e.Use(recovery())
	e.Use(gzip.Gzip(gzip.DefaultCompression))
	e.NoRoute(simpleHandles(bundleRoot)...)
	return e
}
