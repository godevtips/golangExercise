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

	FysiekeSchijf := 0 // Vul uw eigen fysieke schijf nummer in
	partitieStijl, partitieError := windows_partition.SchijfPartitieOphalen(FysiekeSchijf)
	if partitieError != nil {
		log.Fatal(partitieError)
	}

	if partitieStijl == windows_partition.MBR {
		fmt.Printf("\n\nSchijf maakt gebruik van MBR-partitionering.\n")
		leesMBRBootSector(FysiekeSchijf)
	} else {
		fmt.Printf("\nSchijf maakt gebruik van GPT-partitionering.\n\n")
		leesGPTBootSector(FysiekeSchijf)
	}
}

func leesMBRBootSector(fysiekeSchijfNummer int) {

	pad := windows_partition.FysiekSchijfpadOphalen(fysiekeSchijfNummer)
	fmt.Printf("Lezen van MBR-bootsector van '%s'........", pad)

	bestand, schijfError := os.Open(pad)
	if schijfError != nil {
		log.Fatal("Error bij het openen van de schijf: ", schijfError)
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

func leesGPTBootSector(fysiekeSchijfNummer int) {

	pad := windows_partition.FysiekSchijfpadOphalen(fysiekeSchijfNummer)
	const (
		sectorGrootte   = 512
		gptHeaderOffset = sectorGrootte // LBA 1
	)

	fmt.Printf("Lezen van GPT opstartsector van '%s'..............\n\n", pad)

	bestand, err := os.Open(pad)
	if err != nil {
		panic(err)
	}
	defer func(bestand *os.File) {
		err := bestand.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(bestand)

	// Seek to LBA 1 (GPT Header) -> 512 bytes offset
	_, err = bestand.Seek(gptHeaderOffset, 0)
	if err != nil {
		panic(err)
	}

	// Read GPT Header (92 bytes)
	buffer := make([]byte, sectorGrootte)
	_, err = bestand.Read(buffer)
	if err != nil {
		panic(err)
	}

	// GPT-header bekijken
	var gptHeader GPTHeader
	lezer := bytes.NewReader(buffer)
	err = binary.Read(lezer, binary.LittleEndian, &gptHeader)
	if err != nil {
		fmt.Println("Fout bij het parseren van GPT-header:", err)
		return
	}

	// Verify GPT signature
	signature := gptHeader.Handtekening[:]
	if string(signature) != "EFI PART" {
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
