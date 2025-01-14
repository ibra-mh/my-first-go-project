package main

import (
	"fmt"
	"my-first-go-project/cmd"
)

func main() {
	fmt.Println("Hello, World !")

	cmd.Application.Run()
}
