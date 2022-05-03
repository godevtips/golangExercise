package main

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
)

const project_pad = "./golang/noisy_image/nederlands"

func main() {

	VerwijderingFout := bestandVerwijderen("resultaat.png")
	if VerwijderingFout != nil {
		fmt.Println(VerwijderingFout)
	}

	originele_afbeelding, fout := jpegAfbeeldingUithalen("afbeelding")

	if fout != nil {
		log.Fatal(fout)
	}

	riusLaagBeeld, secondError := riusBeeldLaagAanmaken(originele_afbeelding, 50)

	if secondError != nil {
		log.Fatal(secondError)
	}

	resulterendeBeeldFout := ruisendBeeldAanmaken(originele_afbeelding, riusLaagBeeld)

	if resulterendeBeeldFout != nil {
		log.Fatal(resulterendeBeeldFout)
	}
}

func bestandVerwijderen(bestand string) error {

	// verwijder bestand als het bestaat met de functie Remove()
	verwijderingFout := os.Remove(fmt.Sprintf("%s/%s", project_pad, bestand))
	if verwijderingFout != nil {
		return verwijderingFout
	}

	return nil
}

func jpegAfbeeldingUithalen(bestandsnaam string) (image.Image, error) {

	beeldbestand, fout := os.Open(fmt.Sprintf("%s/%s.jpg", project_pad, bestandsnaam))
	if fout != nil {
		return nil, errors.New(fmt.Sprintf("origineel bestand kan niet geopend: %s", fout))
	}

	gedecodeerdBeeld, fout := jpeg.Decode(beeldbestand)
	if fout != nil {
		return nil, errors.New(fmt.Sprintf("origineel beeld kan niet worden gedecodeerd: %s", fout))
	}

	defer func(beeld_bestand *os.File) {
		_ = beeld_bestand.Close()
	}(beeldbestand)

	return gedecodeerdBeeld, nil
}

func riusBeeldLaagAanmaken(achtergrondafbeelding image.Image, doorzichtigheid int) (image.Image, error) {

	var gewensteDoorzichtigHeid = uint8(0)
	if doorzichtigheid > 0 && doorzichtigheid < 256 {
		gewensteDoorzichtigHeid = uint8(math.Round(float64(doorzichtigheid) * 2.56))
	}

	afbeeldingGrenzen := achtergrondafbeelding.Bounds()
	afbeeldingBreedte := afbeeldingGrenzen.Dx()
	afbeeldingHoogte := afbeeldingGrenzen.Dy()
	ruis_afbeelding_laag := image.NewRGBA(image.Rect(0, 0, afbeeldingBreedte, afbeeldingHoogte))

	for p := 0; p < afbeeldingBreedte*afbeeldingHoogte; p++ {
		pixelverschuving := 4 * p
		ruis_afbeelding_laag.Pix[0+pixelverschuving] = uint8(rand.Intn(256))
		ruis_afbeelding_laag.Pix[1+pixelverschuving] = uint8(rand.Intn(256))
		ruis_afbeelding_laag.Pix[2+pixelverschuving] = uint8(rand.Intn(256))
		ruis_afbeelding_laag.Pix[3+pixelverschuving] = gewensteDoorzichtigHeid
	}

	tijdelijkRuisendBeeldbestand, png_fout := os.Create(fmt.Sprintf("%s/%s", project_pad, "tijdelijk_ruisend_beeld.png"))
	if png_fout != nil {
		log.Fatal(png_fout)
	}

	coderingsfout := png.Encode(tijdelijkRuisendBeeldbestand, ruis_afbeelding_laag)
	if coderingsfout != nil {
		return nil, coderingsfout
	}

	tijdelijkRuisendBeeld, beeld_opening_fout := os.Open(fmt.Sprintf("%s/%s", project_pad, "tijdelijk_ruisend_beeld.png"))
	if beeld_opening_fout != nil {
		return nil, beeld_opening_fout
	}

	resultaat, png_decoderingsfout := png.Decode(tijdelijkRuisendBeeld)
	if png_decoderingsfout != nil {
		return nil, png_decoderingsfout
	}

	defer func(tijdelijk_ruisend_beeld *os.File) {
		_ = tijdelijk_ruisend_beeld.Close()
	}(tijdelijkRuisendBeeld)

	// verwijder tijdelijk_ruisend_beeld.png met remove() functie
	verwijderingsfout := bestandVerwijderen("tijdelijk_ruisend_beeld.png")
	if verwijderingsfout != nil {
		return nil, verwijderingsfout
	}

	return resultaat, nil
}

func ruisendBeeldAanmaken(origineleAfbeelding image.Image, beeldRiusLaag image.Image) error {

	verschuving := image.Pt(0, 0)
	origineleAfbeeldingsGrenzen := origineleAfbeelding.Bounds()
	ruisBeeld := image.NewRGBA(origineleAfbeeldingsGrenzen)
	draw.Draw(ruisBeeld, origineleAfbeeldingsGrenzen, origineleAfbeelding, image.Point{}, draw.Src)
	draw.Draw(ruisBeeld, beeldRiusLaag.Bounds().Add(verschuving), beeldRiusLaag, image.Point{}, draw.Over)

	resulterendeAfbeeldingBestand, err := os.Create(fmt.Sprintf("%s/%s", project_pad, "resultaat.png"))
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create result image: %s", err))
	}

	pngCoderingsfout := png.Encode(resulterendeAfbeeldingBestand, ruisBeeld)
	if pngCoderingsfout != nil {
		return pngCoderingsfout
	}

	defer func(afbeelding_bestand *os.File) {
		_ = afbeelding_bestand.Close()
	}(resulterendeAfbeeldingBestand)

	return nil
}
