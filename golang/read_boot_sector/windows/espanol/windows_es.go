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

	diskPath := "/dev/disk0" // Reemplace con su propia ruta de disco
	diskPartition, readingPartitionError := windows_partition.ObtenerParticionDeDisco(diskPath)

	if readingPartitionError != nil {
		log.Fatal(readingPartitionError)
	}

	if diskPartition == windows_partition.MBR {
		fmt.Printf("\n\nEl disco de destino utiliza particionamiento MBR.\n")
		leeSectorDeArranqueMBR(diskPath)
	} else {
		fmt.Printf("\nEl disco de destino utiliza particionamiento GPT.\n\n")
		leeSectorGPT(diskPath)
	}
}

func leeSectorDeArranqueMBR(ruta string) {

	fmt.Printf("Leyendo el sector de arranque del MBR '%s'........", ruta)

	file, readingDiskError := os.Open(ruta)
	if readingDiskError != nil {
		log.Fatal("Error:", readingDiskError)
		return
	}

	byteSlice := make([]byte, 512)
	numBytesRead, err := file.Read(byteSlice)
	if err != nil {
		log.Fatal("Error: " + err.Error())
	}

	fmt.Printf("\n\n|----------------- DATOS MBR ----------------|\n")
	fmt.Printf("Bytes leídos: %d\n", numBytesRead)
	fmt.Printf("Datos en formato decimal:\n%d\n", byteSlice)
	fmt.Printf("Datos en formato hexadecimal:\n%x\n", byteSlice)
	fmt.Printf("Data en formato string:\n%s\n", byteSlice)
	fmt.Printf("|--------------------- DATOS MBR ----------------|\n")
}

type GPTHeader struct {
	Signatura              [8]byte  // Signatura "EFI PART"
	Revision               uint32   // Versión GPT
	Tamano                 uint32   // Tamaño del GPT header
	_                      [4]byte  // Comprobación CRC32 (ignorado)
	_                      [4]byte  // Reservado
	LBA                    uint64   // LBA del GPT header
	BackupLBA              uint64   // LBA backup del GPT
	LBA_Primero            uint64   // Primer LBA utilizable para particiones
	LBA_Ultimo             uint64   // Último LBA utilizable para particiones
	GUID_Del_Disco         [16]byte // Identificador de disco
	TablaDelParticionesLBA uint64   // Tabla de particiones LBA
	NumParticiones         uint32   // Particiones numéricas
	TamanoDelParticion     uint32   // Tamaño de cada partición
}

func leeSectorGPT(ruta string) {

	fmt.Printf("Leyendo el sector de arranque GPT de '%s'..............\n\n", ruta)

	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer func(archivo *os.File) {
		err := archivo.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(archivo)

	// Buscar LBA 1 (encabezado GPT) -> desplazamiento de 512 bytes
	_, err = archivo.Seek(512, 0)
	if err != nil {
		fmt.Println("GPT header error:", err)
		return
	}

	// Leer el encabezado GPT (92 bytes)
	headerData := make([]byte, 92)
	_, err = archivo.Read(headerData)
	if err != nil {
		fmt.Println("Error leer GPT header:", err)
		return
	}

	// Analizar el encabezado GPT
	var gptHeader GPTHeader
	reader := bytes.NewReader(headerData)
	err = binary.Read(reader, binary.LittleEndian, &gptHeader)
	if err != nil {
		fmt.Println("Error analizando GPT header:", err)
		return
	}

	// Comprobar firma
	signature := gptHeader.Signatura[:]
	if string(signature) != "EFI PART" {
		fmt.Println("GPT Signatura invalido!")
		return
	}

	// Imprimir detalles de GPT
	fmt.Printf("|---------------- Datos GPT Header -----------------|\n\n")
	fmt.Printf("  GUID del disco: %x\n", gptHeader.GUID_Del_Disco)
	fmt.Printf("  Primero LBA: %d\n", gptHeader.LBA_Primero)
	fmt.Printf("  Ulimo LBA: %d\n", gptHeader.LBA_Ultimo)
	fmt.Printf("  Tabla del particiones: %d\n", gptHeader.TablaDelParticionesLBA)
	fmt.Printf("  Número de Particiones: %d\n", gptHeader.NumParticiones)
	fmt.Printf("  Tamaño del particion: %d bytes\n\n", gptHeader.TamanoDelParticion)
	fmt.Printf("|---------------- Datos GPT Header -----------------|\n\n")

}
