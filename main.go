package main

import (
	"api_crud_blog/app/router"
	"api_crud_blog/app/utils"
)

func main() {
	route := router.SetupRouter(utils.GetConfig("GIN_MODE"))

	route.Run(":" + utils.GetConfig("PORT"))
}