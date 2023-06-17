package fileread

import "os"

func ReadContent() []byte {
	f, err := os.ReadFile("./testfile.txt")
	if err != nil {
		panic(err)
	}

	return f
}
