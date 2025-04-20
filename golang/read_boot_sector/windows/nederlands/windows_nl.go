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

	schijfPad := "/dev/disk0" // Vul uw eigen schijfpad in
	schijfPartitie, partitieLezenError := windows_partition.SchijfPartitieOphalen(schijfPad)

	if partitieLezenError != nil {
		log.Fatal(partitieLezenError)
	}

	if schijfPartitie == windows_partition.MBR {
		fmt.Printf("\n\nSchijf maakt gebruik van MBR-partitionering.\n")
		lessMBRBootSector(schijfPad)
	} else {
		fmt.Printf("\nSchijf maakt gebruik van GPT-partitionering.\n\n")
		leesGPTBootSector(schijfPad)
	}

}

func lessMBRBootSector(path string) {
	fmt.Printf("Lezen van MBR-bootsector van '%s'........", path)

	bestand, schijfOpenenError := os.Open(path)
	if schijfOpenenError != nil {
		log.Fatal("Error bij het openen van de schijf: ", schijfOpenenError)
		return
	}

	byteSlice := make([]byte, 512)
	aantalGelezenBytes, err := bestand.Read(byteSlice)
	if err != nil {
		log.Fatal("Error: " + err.Error())
	}

	fmt.Printf("\n\n|----------------- MBR GEGEVENS ----------------|\n")
	fmt.Printf("Gelezen bytes: %d\n", aantalGelezenBytes)
	fmt.Printf("Gegevens als decimaal:\n%d\n", byteSlice)
	fmt.Printf("Gegevens in hexadecimaal:\n%x\n", byteSlice)
	fmt.Printf("Gegevens als string:\n%s\n", byteSlice)
	fmt.Printf("|--------------------- MBR GEGEVENS ----------------|\n")
}

type GPTHeader struct {
	Handtekening        [8]byte  // "EFI PART" handtekening
	Herziening          uint32   // GPT herziening
	HeaderGrootte       uint32   // Grootte van GPT header
	_                   [4]byte  // CRC32 checksum (genegeerd)
	_                   [4]byte  // Gereserveerd
	HuidigeLBA          uint64   // LBA van GPT header
	BackupLBA           uint64   // LBA van back-up GPT
	EersteBruikbareLBA  uint64   // Eerste bruikbare LBA partities
	LaatsteBruikbareLBA uint64   // Laatste bruikbare LBA partities
	SchijfGUID          [16]byte // Unieke schijf-ID
	PartitieTabelLBA    uint64   // LBA van windows_partition tabel
	AantalPartities     uint32   // Aantal partities
	Partitiegrootte     uint32   // Grootte van elke partitie
}

func leesGPTBootSector(pad string) {
	fmt.Printf("Lezen van GPT opstartsector van '%s'..............\n\n", pad)

	bestand, err := os.Open(pad)
	if err != nil {
		fmt.Println("Error bij het openen van de schijf: ", err)
		return
	}

	defer func(bestand *os.File) {
		err := bestand.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(bestand)

	// Zoek naar LBA 1 (GPT-header) -> 512 bytes offset
	_, err = bestand.Seek(512, 0)
	if err != nil {
		fmt.Println("Error bij het zoeken naar GPT-Header: ", err)
		return
	}

	// GPT-header lezen (92 bytes)
	headerGegevens := make([]byte, 92)
	_, err = bestand.Read(headerGegevens)
	if err != nil {
		fmt.Println("Error bij het lezen van GPT-header: ", err)
		return
	}

	// GPT-header bekijken
	var gptHeader GPTHeader
	lezer := bytes.NewReader(headerGegevens)
	err = binary.Read(lezer, binary.LittleEndian, &gptHeader)
	if err != nil {
		fmt.Println("Error bij het lezen van de GPT-header:", err)
		return
	}

	// Handtekening controleer
	handtekening := gptHeader.Handtekening[:]
	if string(handtekening) != "EFI PART" {
		fmt.Println("Ongeldige GPT-handtekening!")
		return
	}

	// GPT-gegevens afdrukken
	fmt.Printf("|---------------- GPT Header Gegevens -----------------|\n\n")
	fmt.Printf("  Schijf GUID: %x\n", gptHeader.SchijfGUID)
	fmt.Printf("  Eertst bruikbare LBA: %d\n", gptHeader.EersteBruikbareLBA)
	fmt.Printf("  Laatste bruikbare LBA: %d\n", gptHeader.LaatsteBruikbareLBA)
	fmt.Printf("  Partitie tabel LBA: %d\n", gptHeader.PartitieTabelLBA)
	fmt.Printf("  Aantal partities: %d\n", gptHeader.AantalPartities)
	fmt.Printf("  Partitie-grootte: %d bytes\n\n", gptHeader.Partitiegrootte)
	fmt.Printf("|---------------- GPT Header Gegevens -----------------|\n\n")
}
