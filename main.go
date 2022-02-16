package main

import "github.com/craig1122321/go_anti_smug_proxy/api"

func main() {
	router := api.InitRouter()
	router.Run(":8000")
}
