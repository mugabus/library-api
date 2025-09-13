package main

import "library-api/routes"

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")

}
