package api

import (
	"anti-smuggling-proxy/api/service"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ParameterFiltering(context *gin.Context) {
	fmt.Println("Checker Loading.....")
	reqHeaederTmp := context.Request.Header
	reqHeader := service.MapToString(reqHeaederTmp)
	reqHeader = strings.Replace(reqHeader, "]\"", "", -1)
	reqHeader = strings.Replace(reqHeader, "\"[", "", -1)
	reqMethod := context.Request.Method

	//streamBuf := make([]byte, 2048, 4096)
	//reqStreamNum, _ := context.Request.Body.Read(streamBuf)
	//reqStream := string(streamBuf[0:reqStreamNum])

	//上面三行可以使用 但在這裡用的話 router那邊的req stream會變空值

	reqStream := "" ///////////////Need EDIT /////////////////////

	//uri := context.Params
	clMatch := service.Regex(reqHeader, "")["clMatch"]
	teMatch := service.Regex(reqHeader, "")["teMatch"]
	clSum, _ := strconv.Atoi(service.Regex(reqHeader, "")["clCount"])
	teSum, _ := strconv.Atoi(service.Regex(reqHeader, "")["teCount"])
	matchAllLen := clSum + teSum
	fmt.Println("Handler Match All len: ", matchAllLen)
	//fmt.Println(reqMethod, uri, "\n", reqHeader, "\n", reqStream, "\n", "clMatch: ", clMatch, "\n", teMatch, "\n", "match ALL Len", matchAllLen)
	if matchAllLen == 0 {
		fmt.Println("match len == 0")
		if !service.ReqStreamCheck(reqHeader, reqMethod, reqStream) {
			context.AbortWithStatus(400)
		}

	} else if matchAllLen == 1 {
		fmt.Println("match len == 1")
		switch {
		case !service.GetLengthCheck(reqMethod, reqHeader):
			context.AbortWithStatus(400)
		case !service.TeCheck(reqHeader):
			context.AbortWithStatus(400)
		case !service.ClCheck(reqHeader):
			context.AbortWithStatus(400)

		}
	} else if matchAllLen > 1 {
		fmt.Println("match len > 1")
		switch {
		case clMatch == "Content-Length" && teMatch == "Transfer-Encoding":
			context.Request.Header.Del("Content-Length")
		case !service.TeCheck(reqHeader):
			context.AbortWithStatus(400)
		case !service.ClCheck(reqHeader):
			context.AbortWithStatus(400)

		}

	}
}
