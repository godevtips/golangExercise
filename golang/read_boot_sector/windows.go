package main

import (
	"fmt"
	"os"
)

func main() {

	diskPath := "/dev/disk0" // replace your own disk path

	file, err := os.Open(diskPath)
	if err != nil {
		fmt.Println("❌ Error opening disk:", err)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

}
