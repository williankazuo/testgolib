package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(string(TestFunction()))
}

func TestFunction() []byte {
	f, err := os.ReadFile("./testfile.txt")
	if err != nil {
		panic(err)
	}

	return f
}
