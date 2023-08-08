package main

import (
	"wese/core/module_template/routes"
	"wese/core/module_template/services"
)

func main() {
	services.DatabaseConnectAndMigrate()
	routes.RouteName()
}
