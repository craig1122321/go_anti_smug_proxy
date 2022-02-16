package main

import (
	"anti-smuggling-proxy/api"
)

func main() {
	router := api.InitRouter()
	router.Run(":8000")
}
