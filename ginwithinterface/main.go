package main

import (
	"ginwithinterface/Config"
	"ginwithinterface/Routing"
)

func main() {

	Config.DataMigration()
	Routing.Handler()
}
