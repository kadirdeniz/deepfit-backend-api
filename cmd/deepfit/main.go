package main

import (
	"deepfit/tools/fiber"
	"deepfit/tools/mongodb"
	"fmt"
)

func main() {
	fmt.Println("Deepfit Application Started!")

	// Connect Database
	mongodb.CreateConnection()

	// Start Server
	fiber.Router()
}
