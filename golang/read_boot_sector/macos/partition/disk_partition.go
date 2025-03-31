package partition

import (
	"errors"
	"fmt"
	"log"
	"os"
)

const sectorSize = 512 // Standard disk sector size

type DiskPartition int

const (
	UNKNOWN = iota
	MBR     = iota
	GPT     = iota
)

func (partition DiskPartition) String() string {
	switch partition {
	case MBR:
		return "MBR"
	case GPT:
		return "GPT"
	default:
		return "Unknown"
	}
}

func GetDiskPartition(diskPath string) (DiskPartition, error) {

	file, err := os.Open(diskPath)
	if err != nil {
		return UNKNOWN, errors.New(fmt.Sprintf("Error opening disk: %s", err))
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	//Read LBA 0 (MBR)
	mbr := make([]byte, sectorSize)
	_, err = file.Read(mbr)
	if err != nil {
		return UNKNOWN, errors.New(fmt.Sprintf("Error reading MBR: %s", err))
	}

	// Check for MBR signature (0x55AA at offset 510-511)
	if mbr[510] == 0x55 && mbr[511] == 0xAA {

		// Seek to LBA 1 (GPT header)
		_, err = file.Seek(sectorSize, 0)
		if err != nil {
			return UNKNOWN, errors.New(fmt.Sprintf("Error seeking to GPT header: %s", err))
		}

		// Read GPT Header
		gptHeader := make([]byte, 8)
		_, err = file.Read(gptHeader)
		if err != nil {
			return UNKNOWN, errors.New(fmt.Sprintf("Error reading GPT header: %s", err))
		}

		// Check GPT signature
		gptHeaderText := string(gptHeader)
		if gptHeaderText == "EFI PART" {
			return GPT, nil
		} else {
			return MBR, nil
		}

	} else {
		return UNKNOWN, errors.New("no valid MBR found. Disk may be uninitialized")
	}
}
