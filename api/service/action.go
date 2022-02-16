package service

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func RemoveDuplicateElement(languages []string) []string {
	result := make([]string, 0, len(languages))
	temp := map[string]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func MapToString(m map[string][]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}
func ReverseProxy() gin.HandlerFunc {
	target := "http://127.0.0.1:3000/test"

	return func(context *gin.Context) {
		director := func(req *http.Request) {
			r := context.Request
			req = r
			req.URL.Scheme = "http"
			req.Host, req.URL.Host = target, target
			req.Header = context.Request.Header
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(context.Writer, context.Request)
	}

}
