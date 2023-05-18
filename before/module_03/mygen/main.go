package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Running go generate on %s, with args %s\n",
		os.Getenv("GOFILE"), os.Args[1:])
}
