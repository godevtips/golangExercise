package windows_partition

import (
	"errors"
	"fmt"
	"golang.org/x/sys/windows"
	"log"
	"os"
	"syscall"
	"unsafe"
)

const (
	IOCTL_DISK_GET_PARTITION_INFO_EX = 0x00070048

	// Partition styles
	PARTITION_STYLE_MBR = 0
	PARTITION_STYLE_GPT = 1
	PARTITION_STYLE_RAW = 2
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

func GetPhysicalDrivePath(physicalDrive int) string {
	return fmt.Sprintf("\\\\.\\PhysicalDrive%d", int64(physicalDrive))
}

func GetDiskPartition(physicalDrive int) (DiskPartition, error) {

	path := GetPhysicalDrivePath(physicalDrive)
	utf16Path, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return UNKNOWN, err
	}

	handle, err := windows.CreateFile(
		utf16Path,
		windows.GENERIC_READ,
		windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE,
		nil,
		windows.OPEN_EXISTING,
		0,
		0,
	)
	if err != nil {
		return UNKNOWN, err
	}
	defer func(handle windows.Handle) {
		err := windows.CloseHandle(handle)
		if err != nil {
			log.Fatalf("Error closing window handle %v", err)
		}
	}(handle)

	// Allocate a large enough buffer
	bufSize := 1024
	buffer := make([]byte, bufSize)
	var bytesReturned uint32

	err = windows.DeviceIoControl(
		handle,
		IOCTL_DISK_GET_PARTITION_INFO_EX,
		nil,
		0,
		&buffer[0],
		uint32(len(buffer)),
		&bytesReturned,
		nil,
	)
	if err != nil {
		return -1, err
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return UNKNOWN, errors.New(fmt.Sprintf("Error opening disk: %s", err))
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
	}(file)

	//Read LBA 0 (MBR)
	mbr := make([]byte, sectorSize)
	_, err = file.Read(mbr)
	if err != nil {
		return UNKNOWN, errors.New(fmt.Sprintf("Error reading MBR: %s", err))
	}

	// First 4 bytes of output buffer = PartitionStyle enum (uint32)
	diskPartitionStyle := *(*uint32)(unsafe.Pointer(&buffer[0]))

	switch diskPartitionStyle {

	case PARTITION_STYLE_MBR:
		return MBR, nil
	case PARTITION_STYLE_GPT:
		return GPT, nil
	default:
		return UNKNOWN, nil
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

func ObtenerRutaDelDisco(numeroDeUnidad int) string {
	return fmt.Sprintf("\\\\.\\PhysicalDrive%d", int64(numeroDeUnidad))
}

func ObtenerParticionDeDisco(numeroDeUnidad int) (DiskPartition, error) {

	ruta := ObtenerRutaDelDisco(numeroDeUnidad)
	rutaUtf16, err := syscall.UTF16PtrFromString(ruta)
	if err != nil {
		return UNKNOWN, err
	}

	handler, err := windows.CreateFile(
		rutaUtf16,
		windows.GENERIC_READ,
		windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE,
		nil,
		windows.OPEN_EXISTING,
		0,
		0,
	)
	if err != nil {
		return UNKNOWN, err
	}
	defer func(handle windows.Handle) {
		err := windows.CloseHandle(handle)
		if err != nil {
			log.Fatalf("Error al cerrar el identificador de la ventana %v", err)
		}
	}(handler)

	// Asignar un buffer suficientemente
	tamanoDelBufer := 1024
	bufer := make([]byte, tamanoDelBufer)
	var bytesDevueltos uint32

	err = windows.DeviceIoControl(
		handler,
		IOCTL_DISK_GET_PARTITION_INFO_EX,
		nil,
		0,
		&bufer[0],
		uint32(len(bufer)),
		&bytesDevueltos,
		nil,
	)
	if err != nil {
		return -1, err
	}

	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Println(err)
		return UNKNOWN, errors.New(fmt.Sprintf("Error opening disk: %s", err))
	}

	defer func(archivo *os.File) {
		err := archivo.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(archivo)

	//Leer LBA 0 (MBR)
	mbr := make([]byte, sectorSize)
	_, err = archivo.Read(mbr)
	if err != nil {
		return UNKNOWN, errors.New(fmt.Sprintf("Error al leer el MBR: %s", err))
	}

	// Los primeros 4 bytes del búfer de salida (output) = enumeración particion del disco  (uint32)
	particionDelDisco := *(*uint32)(unsafe.Pointer(&bufer[0]))

	switch particionDelDisco {

	case PARTITION_STYLE_MBR:
		return MBR, nil
	case PARTITION_STYLE_GPT:
		return GPT, nil
	default:
		return UNKNOWN, nil
	}
}
