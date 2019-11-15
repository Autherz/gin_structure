package main

import (
	"gin_workshop/routes"
)
func main() {

	r := routes.SetupRouter()
	r.Run()
}