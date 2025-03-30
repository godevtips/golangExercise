package main

import (
	"fmt"
	"godevtips.com/m/v2/partition"
	"log"
	"os"
)

func main() {

	diskPath := "/dev/disk0"
	diskPartition, readingPartitionError := partition.GetDiskPartition(diskPath)

	if readingPartitionError != nil {
		log.Fatal(readingPartitionError)
	}

	if diskPartition == partition.MBR {
		fmt.Println("Disk is using GPT Partitioning.")
		readMBRBootSector(diskPath)
	} else {
		fmt.Println("Disk is using MBR Partitioning.")
		readGPTBootSector(diskPath)
	}

}

func readMBRBootSector(path string) {
	fmt.Println("Reading MBR boot sector of: ", path)

	file, readingDiskError := os.Open(path)
	if readingDiskError != nil {
		log.Fatal("Error opening disk:", readingDiskError)
		return
	}

	byteSlice := make([]byte, 512)
	numBytesRead, err := file.Read(byteSlice)
	if err != nil {
		log.Fatal("Error: " + err.Error())
	}

	fmt.Println("MBR DATA")
	fmt.Printf("Bytes read: %d\n\n", numBytesRead)
	fmt.Printf("Data as decimal:\n%d\n\n", byteSlice)
	fmt.Printf("Data as hex:\n%x\n\n", byteSlice)
	fmt.Printf("Data as string:\n%s\n\n", byteSlice)
}

func readGPTBootSector(path string) {
	fmt.Println("Reading GPT boot sector of: ", path)
}
