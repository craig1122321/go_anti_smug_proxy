package service

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Regex(reqHeader, reqStream string) map[string]string {

	teRegex, _ := regexp.Compile(`(?ims)[^\W]*[\S]*Transfer-Encoding[\W]*[\w]*[$=]?`)
	clRegex, _ := regexp.Compile(`(?ims)[^\W]*[\S]*Content-Length[\W]*[\w]*[$=]?`)
	methodRegex, _ := regexp.Compile(`(?ims)(GET|POST|HEAD|OPTION|PUT|DELETE|TRACE|PATCH|CONNECT)\s([a-z0-9\-._~%!$&'()*+,;=:@/?#]*)\sHTTP/.+`)

	teMatch := teRegex.FindAllString(reqHeader, -1)
	teMatchStr := strings.Join(teMatch, ",") //
	teCount := fmt.Sprint(len(teRegex.FindAllStringIndex(reqHeader, -1)))

	clMatch := clRegex.FindAllString(reqHeader, -1)
	clMatchStr := strings.Join(clMatch, ",")
	clCount := fmt.Sprint(len(clRegex.FindAllStringIndex(reqHeader, -1)))

	methodMatch := methodRegex.FindAllString(reqStream, -1) //
	methodMatchStr := strings.Join(methodMatch, ",")
	methodCount := fmt.Sprint(len(methodRegex.FindAllStringIndex(reqStream, -1)))

	reMap := map[string]string{"teMatch": teMatchStr, "teCount": teCount, "clMatch": clMatchStr, "clCount": clCount, "methodMatch": methodMatchStr, "methodCount": methodCount}

	return reMap
}

func TeCheck(reqHeader string) bool {
	fmt.Println("TE Checking...")
	teMatch := Regex(reqHeader, "")["teMatch"]
	reqTeValue := strings.Split(teMatch, ": ")
	valList := []string{"chunked", "compress", "deflate", "gzip", "identity"}
	teSum, _ := strconv.Atoi(Regex(reqHeader, "")["teCount"])
	var output bool
	tmpList := append(reqTeValue, valList...)
	tmpList = RemoveDuplicateElement(tmpList)

	switch {
	case teSum == 1: // TE header check
		te := strings.Replace(teMatch, ":", "", -1)
		if te != ("Transfer-Encoding") {
			output = false
			fmt.Println("TE Check Fail #1")
		} else {
			output = true
		}
	case teSum > 1:
		output = false
		fmt.Println("TE Check Fail #2")

	case len(tmpList) > 6: // TE value check
		output = false
		fmt.Println("TE Check Fail #3", len(tmpList))
	default:
		output = true
	}

	return output
}
func ClCheck(reqHeader string) bool {
	fmt.Println("CL Checking...")
	clMatch := Regex(reqHeader, "")["clMatch"]
	clSum, _ := strconv.Atoi(Regex(reqHeader, "")["clCount"])
	var output bool
	switch {
	case clSum == 1:
		cltmp := strings.Split(clMatch, "=")
		cl := cltmp[0]

		if cl != ("Content-Length") {
			output = false
			fmt.Println("CL Check Fail #1: ", cl)
		} else {
			output = true
		}
	case clSum > 1:
		output = false
		fmt.Println("CL Check Fail #2", clSum, "\n", clMatch)
	default:
		output = true
	}

	return output
}
func GetLengthCheck(reqMethod, reqHeader string) bool {
	fmt.Println("GET Method Len Checking...")
	//clMatch := Regex(reqHeader, "")["clMatch"]
	//teMatch := Regex(reqHeader, "")["teMatch"]
	clSum, _ := strconv.Atoi(Regex(reqHeader, "")["clCount"])
	teSum, _ := strconv.Atoi(Regex(reqHeader, "")["teCount"])
	var output bool
	switch {
	case reqMethod == "GET":
		if clSum > 0 || teSum > 0 {
			output = false
			fmt.Println("GET Method CL/TE Check Fail #1")
		} else {
			output = true
		}
	default:
		output = true
	}
	return output
}

func ReqStreamCheck(reqHeader, reqMethod, reqStream string) bool {
	fmt.Println("Request Stream Checking...")
	var output bool
	//methodMatch := Regex(reqHeader, "")["methodMatch"]
	reqStreamMethodSum, _ := strconv.Atoi(Regex(reqHeader, "")["methodCount"])
	switch {
	case len(reqMethod) != 0:
		if reqStreamMethodSum > 0 {
			output = false
			fmt.Println("Req Stream Check Fail #1")
		} else {
			output = true
		}
	default:
		output = true
	}

	return output
}
