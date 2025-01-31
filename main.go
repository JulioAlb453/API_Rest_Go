package main

import (
	"API_ejemplo/album/infraestructure"

)

func main() {

	deps := infraestructure.Init()
	router := infraestructure.Routes(deps)
	router.Run(":8080")

}
