package proxy

import (
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	urlpckg "net/url"
)

func NewReverseProxy(targetURL string) gin.HandlerFunc {
	url, _ := urlpckg.Parse(targetURL)
	proxy := httputil.NewSingleHostReverseProxy(url)

	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
