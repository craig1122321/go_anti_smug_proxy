package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

const Host = "127.0.0.1:3000/"

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Any("/:path", ReverseProxy, func(context *gin.Context) {

	})
	return router
}

func ReverseProxy(c *gin.Context) {

	fmt.Println("Reverse Proxy Loading.....")
	ParameterFiltering(c) //
	remote, err := url.Parse("http://127.0.0.1:3000/")
	remote.Path = c.FullPath()

	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		streamBuf := make([]byte, 2048, 4096)
		reqStreamNum, _ := c.Request.Body.Read(streamBuf)
		reqStream := string(streamBuf[0:reqStreamNum])

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(reqStream)))
		fmt.Println(c.Request.Body)
		//body, _ := c.GetRawData()
		//c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(body)))

		//fmt.Println("Body:", body)
		//req.URL.Path = c.Param("test")
		//defer req.Body.Close()
		//fmt.Println("Rev Proxy req body ", body)

	}
	//fmt.Println("Rev Proxy c.Request: ")

	proxy.ServeHTTP(c.Writer, c.Request)
}
