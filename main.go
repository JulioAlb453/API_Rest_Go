package main

import "API_ejemplo/src/routes"


func main() {
	router := routes.SetupRouter()

	router.Run(":8080")

}
