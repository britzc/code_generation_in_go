//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := os.Getenv("GOFILE")
	fmt.Println(os.Getenv("TEMP01"))

	fmt.Printf("Running go generate on %s\n", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var count int
	for scanner.Scan() {
		count++
	}

	fmt.Println("Number of lines in file:", count)
}
