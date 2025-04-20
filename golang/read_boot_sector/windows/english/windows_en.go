package main

import (
	"boot_sector/windows_partition"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func main() {

	physicalDrive := 0 // your own physical disk number
	partitionStyle, partitionError := windows_partition.GetDiskPartition(physicalDrive)
	if partitionError != nil {
		log.Fatal(partitionError)
	}

	if partitionStyle == windows_partition.MBR {
		fmt.Printf("\n\nDisk is using MBR Partitioning.\n")
		readMBRBootSector(physicalDrive)
	} else {
		fmt.Printf("\nDisk is using GPT Partitioning.\n\n")
		readGPTData(physicalDrive)
	}
}

func readMBRBootSector(physicalDisk int) {
	fmt.Printf("Reading MBR boot sector of physical disk '%d'........", physicalDisk)

	path := windows_partition.GetPhysicalDrivePath(physicalDisk)

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

	fmt.Printf("\n\n|----------------- MBR DATA ----------------|\n")
	fmt.Printf("Bytes read: %d\n", numBytesRead)
	fmt.Printf("Data as decimal:\n%d\n", byteSlice)
	fmt.Printf("Data as hex:\n%x\n", byteSlice)
	fmt.Printf("Data as string:\n%s\n", byteSlice)
	fmt.Printf("|----------------- MBR DATA ----------------|\n")
}

type GPTHeader struct {
	Signature         [8]byte  // "EFI PART" signature
	Revision          uint32   // GPT version
	HeaderSize        uint32   // Size of GPT header
	_                 [4]byte  // CRC32 checksum (ignored here)
	_                 [4]byte  // Reserved
	CurrentLBA        uint64   // LBA of GPT header
	BackupLBA         uint64   // LBA of backup GPT
	FirstUsableLBA    uint64   // First usable LBA for partitions
	LastUsableLBA     uint64   // Last usable LBA for partitions
	DiskGUID          [16]byte // Unique disk identifier
	PartitionTableLBA uint64   // LBA of windows_partition table
	NumPartitions     uint32   // Number of windows_partition entries
	PartitionSize     uint32   // Size of each windows_partition entry
}

func readGPTData(physicalDisk int) {

	physicalDiskPath := windows_partition.GetPhysicalDrivePath(physicalDisk)

	const (
		sectorSize      = 512
		gptHeaderOffset = sectorSize // LBA 1
	)

	file, err := os.Open(physicalDiskPath)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	// Seek to LBA 1 (GPT Header) -> 512 bytes offset
	_, err = file.Seek(gptHeaderOffset, 0)
	if err != nil {
		panic(err)
	}

	// Read GPT Header (92 bytes)
	buf := make([]byte, sectorSize)
	_, err = file.Read(buf)
	if err != nil {
		panic(err)
	}

	// Parse GPT header
	var gptHeader GPTHeader
	reader := bytes.NewReader(buf)
	err = binary.Read(reader, binary.LittleEndian, &gptHeader)
	if err != nil {
		fmt.Println("Error parsing GPT header:", err)
		return
	}

	// Verify GPT signature
	signature := gptHeader.Signature[:]
	if string(signature) != "EFI PART" {
		fmt.Println("Invalid GPT Signature!")
		return
	}

	// Print GPT details
	fmt.Printf("|---------------- GPT Header Data -----------------|\n\n")
	fmt.Printf("  Disk GUID: %x\n", gptHeader.DiskGUID)
	fmt.Printf("  First Usable LBA: %d\n", gptHeader.FirstUsableLBA)
	fmt.Printf("  Last Usable LBA: %d\n", gptHeader.LastUsableLBA)
	fmt.Printf("  Partition Table LBA: %d\n", gptHeader.PartitionTableLBA)
	fmt.Printf("  Number of Partitions: %d\n", gptHeader.NumPartitions)
	fmt.Printf("  Partition Entry Size: %d bytes\n\n", gptHeader.PartitionSize)
	fmt.Printf("|---------------- GPT Header Data -----------------|\n\n")
}
