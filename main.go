package main

import (
	"fmt"
	"myapp/routes"
)

func main() {
	r := routes.SetupRouter()
    fmt.Println("Server is running on port 3060")
    r.Run(":3060")
}
