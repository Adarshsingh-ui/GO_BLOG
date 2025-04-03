package main

import (
	"fmt"
	"myapp/Routes"
)

func main() {
	r := Routes.SetupRouter()
    fmt.Println("Server is running on port 3060")
    r.Run(":3060")
}
