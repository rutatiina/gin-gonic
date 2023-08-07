package main

import (
	"wese/_module_template_/services"
)

func main() {
	services.DatabaseConnectAndMigrate()
	routes._route_name_()
}
