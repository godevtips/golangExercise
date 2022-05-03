package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {

	// Voorbeeld bestand
	bestandGegevens, err := os.Stat("./golang/file_info/nederlands/bestands_map/bestand.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("|----------- Bestand gegevens -----------|")
	fmt.Println("Bestand naam:     ", bestandGegevens.Name())
	fmt.Println("Bestand grootte:  ", bestandsGrootteOphalen(bestandGegevens))
	fmt.Println("Rechten:          ", bestandGegevens.Mode())
	fmt.Println("Laatst gewijzigd: ", bestandGegevens.ModTime())
	fmt.Println("Is een map:       ", isEenMap(bestandGegevens))
	fmt.Println("|----------- Bestand gegevens -----------|\n")

	// Als je bovendien de onderliggende gegevensbron wilt uitlezen, kun je gebruik maken van:
	// bestandGegevens.Sys(). Het kan in sommige gevallen nil retouneren.

	// Een map voorbeeld
	mapGegevens, err := os.Stat("./golang/file_info/nederlands/bestands_map")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("|----------- Map gegevens -----------|")
	fmt.Println("Bestand naam:        ", mapGegevens.Name())
	fmt.Println("Bestand grootte:     ", bestandsGrootteOphalen(mapGegevens))
	fmt.Println("Rechten:             ", mapGegevens.Mode())
	fmt.Println("Laatst gewijzigd:    ", mapGegevens.ModTime())
	fmt.Println("Is een map:          ", isEenMap(mapGegevens))
	fmt.Println("|----------- Map gegevens -----------|")

	// Net zoals de vorige voorbeeld, als je de onderliggende gegevensbron wilt ophalen,
	// kunt je het volgende gebruiken:
	// mapGegevens.Sys().
	//Het kan in sommige gevallen nul opleveren.

}

func isEenMap(bestandGegevens fs.FileInfo) string {
	var resultaat = "Nee"
	if bestandGegevens.IsDir() {
		resultaat = "Ja"
	}
	return resultaat
}

//bestandsgrootte ophalen
func bestandsGrootteOphalen(bestandGegevens fs.FileInfo) string {

	bestandsGrootte := bestandGegevens.Size()
	if bestandsGrootte < 1000 { // Bestandsgrootte is in bytes
		return fmt.Sprintf("%d bytes", bestandsGrootte)
	} else { // bestandsgrootte is in kilobytes of megabytes
		bestandsGrootteInKb := float64(bestandsGrootte) / float64(1000)
		if bestandsGrootteInKb > 1000 { // bestandsgrootte is in megabytes
			bestandsGrootteInMb := bestandsGrootteInKb / float64(1000)
			return fmt.Sprintf("%d Mb", bestandsGrootteInMb)
		} else { // bestandsgrootte is in kilobytes
			return fmt.Sprintf("%d Kb", bestandsGrootteInKb)
		}
	}
}