package main

import (
	"API_ejemplo/src/album/infraestructure"

)

func main() {

	deps := infraestructure.Init()
	router := infraestructure.Routes(deps)
	router.Run(":8080")

}
