package main

import (
	"deepfit/tools/fiber"
	"deepfit/tools/mongodb"
	"fmt"
)

func main() {
	fmt.Println("Deepfit Fit Application Started!")

	// Start Server
	fiber.Router()

	// Connect Database
	mongodb.CreateConnection()
}
