package main

import (
	"usersRequest/routes"
	"usersRequest/logger"
)


func main(){
	logger.Init()
	routes.Init()
}