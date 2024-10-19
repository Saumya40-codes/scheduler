package main

import (
	"fmt"

	"github.com/Saumya40-codes/scheduler/cmd/scheduler"
)

func main() {
	fmt.Println("Hello, World!")

	scheduler := scheduler.Create(10)
	scheduler.Run()
}
