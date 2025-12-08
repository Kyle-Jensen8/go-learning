package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	panic("a problem")

	path := filepath.Join(os.TempDir(), "file")
	fmt.Println("path: ", path)
	_, err := os.Create(path)

	if err != nil{
		panic(err)
	}	
}
