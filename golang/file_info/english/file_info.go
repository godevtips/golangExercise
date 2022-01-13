package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {

	// File example
	fileInfo, err := os.Stat("./golang/file_info/english/file_directory/file.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("|----------- File info -----------|")
	fmt.Println("File name:     ", fileInfo.Name())
	fmt.Println("Size:          ", getFileSize(fileInfo))
	fmt.Println("Permissions:   ", fileInfo.Mode())
	fmt.Println("Last modified: ", fileInfo.ModTime())
	fmt.Println("Is Directory:  ", isDirectory(fileInfo))
	fmt.Println("|----------- File info -----------|\n")

	// In addition, if you want to obtain the underlying data source you can use:
	// fileInfo.Sys(). It can return nil in some cases.

	// Directory example
	directoryInfo, err := os.Stat("./golang/file_info/english/file_directory")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("|----------- Directory info -----------|")
	fmt.Println("File name:     ", directoryInfo.Name())
	fmt.Println("Size:          ", getFileSize(directoryInfo))
	fmt.Println("Permissions:   ", directoryInfo.Mode())
	fmt.Println("Last modified: ", directoryInfo.ModTime())
	fmt.Println("Is Directory:  ", isDirectory(directoryInfo))
	fmt.Println("|----------- Directory info -----------|")

	// Same as the previous example, if you want to obtain the underlying data source you can use:
	// directoryInfo.Sys(). It can return nil in some cases.
}

func isDirectory(fileInfo fs.FileInfo) string {
	var result = "No"
	if fileInfo.IsDir() {
		result = "yes"
	}
	return result
}

func getFileSize(fileInfo fs.FileInfo) string {

	size := fileInfo.Size()
	if size < 1000 { // File size is in bytes
		return fmt.Sprintf("%d bytes", size)
	} else { // file size is in kilobytes or megabytes
		fileSizeInKb := float64(size) / float64(1000)
		if fileSizeInKb > 1000 { // file size is in megabytes
			fileSizeInMb := fileSizeInKb / float64(1000)
			return fmt.Sprintf("%d Mb", fileSizeInMb)
		} else { // file size is in kilobytes
			return fmt.Sprintf("%d Kb", fileSizeInKb)
		}
	}
}
