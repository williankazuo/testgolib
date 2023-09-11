package fileread

import (
	"fmt"
	"os"
)

func ReadContent() []byte {
	a := os.Getenv("TESTE")
	fmt.Println("env test ", a)
	f, err := os.ReadFile("./testfile.txt")
	if err != nil {
		panic(err)
	}

	return f
}
