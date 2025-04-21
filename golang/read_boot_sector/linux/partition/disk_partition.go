package partition

import (
	"errors"
	"fmt"
	"log"
	"os"
)

const sectorSize = 512      // Standard disk sector size
const sectorgrootte = 512   // Standard disk sector size
const tamanoDelSector = 512 // Standard disk sector size

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

func SchijfPartitieOphalen(schijfpad string) (DiskPartition, error) {

	bestand, err := os.Open(schijfpad)
	if err != nil {
		return UNKNOWN, errors.New(fmt.Sprintf("Error opening disk: %s", err))
	}

	defer func(bestand *os.File) {
		err := bestand.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(bestand)

	//Less LBA 0 (MBR)
	mbr := make([]byte, sectorgrootte)
	_, err = bestand.Read(mbr)
	if err != nil {
		return UNKNOWN, errors.New(fmt.Sprintf("Error bij het lezen van MBR: %s", err))
	}

	// Controleer op MBR-handtekening (0x55AA op offset 510-511)
	if mbr[510] == 0x55 && mbr[511] == 0xAA {

		// Zoek naar LBA 1 (GPT header)
		_, err = bestand.Seek(sectorgrootte, 0)
		if err != nil {
			return UNKNOWN, errors.New(fmt.Sprintf("Error bij het zoeken naar GPT header: %s", err))
		}

		// GPT-header lezen
		gptHeader := make([]byte, 8)
		_, err = bestand.Read(gptHeader)
		if err != nil {
			return UNKNOWN, errors.New(fmt.Sprintf("Error bij het lezen van GPT header: %s", err))
		}

		// Controleer GPT-handtekening
		gptHeaderTekst := string(gptHeader)
		if gptHeaderTekst == "EFI PART" {
			return GPT, nil
		} else {
			return MBR, nil
		}

	} else {
		return UNKNOWN, errors.New("geen geldige MBR gevonden. Schijf is mogelijk niet geïnitialiseerd")
	}
}

func ObtenerParticionDeDisco(ruta string) (DiskPartition, error) {

	archivo, err := os.Open(ruta)
	if err != nil {
		return UNKNOWN, errors.New(fmt.Sprintf("Error al abrir el disco: %s", err))
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(archivo)

	//Leer LBA 0 (MBR)
	mbr := make([]byte, tamanoDelSector)
	_, err = archivo.Read(mbr)
	if err != nil {
		return UNKNOWN, errors.New(fmt.Sprintf("Error reading MBR: %s", err))
	}

	// Check for MBR signature (0x55AA at offset 510-511)
	if mbr[510] == 0x55 && mbr[511] == 0xAA {

		// Buscar LBA 1 (del GPT header)
		_, err = archivo.Seek(tamanoDelSector, 0)
		if err != nil {
			return UNKNOWN, errors.New(fmt.Sprintf("Error al intentar acceder GPT header: %s", err))
		}

		// Leer GPT header
		gptHeader := make([]byte, 8)
		_, err = archivo.Read(gptHeader)
		if err != nil {
			return UNKNOWN, errors.New(fmt.Sprintf("Error al leer el GPT header: %s", err))
		}

		// Check signatura del GPT
		gptHeaderTexto := string(gptHeader)
		if gptHeaderTexto == "EFI PART" {
			return GPT, nil
		} else {
			return MBR, nil
		}

	} else {
		return UNKNOWN, errors.New("no se encontró un MBR válido. Es posible que el disco no esté inicializado")
	}
}
