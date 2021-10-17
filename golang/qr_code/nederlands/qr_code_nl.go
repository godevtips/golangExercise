package main

import (
	"encoding/base64"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/skip2/go-qrcode"
	"log"
	"net/http"
	"os"
)

var websiteURL = "https://www.godevtips.com"
var afbeeldingsGrootte = 256

func main() {

	if len(os.Args) == 2 && os.Args[1] == "bestand" { // Bepaal welke optie moet worden gestart

		bestands_naam := "qr_code.png"
		_error := qrcode.WriteFile(websiteURL, qrcode.Medium, afbeeldingsGrootte, bestands_naam)

		if _error != nil {
			log.Fatal("Fout bij genereren van QR-code naar bestand ", _error)
		} else {
			fmt.Println("Qr-code bestand aangemaakt!")
		}
	} else {
		serverStarten()
	}
}

func serverStarten() {

	router := httprouter.New()
	router.GET("/", qrCodeVersturen)

	fmt.Println("server aan het luistert op poort 3000")
	serverFout := http.ListenAndServe(":3000", router)
	if serverFout != nil {
		log.Fatal("Kan webserver niet starten, oorzaak: ", serverFout)
	}
}

func qrCodeVersturen(reactie http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	qr_code_afbeelding_tag := genereerQRCodeHtmlImageTag()
	reactie.WriteHeader(http.StatusOK)
	responsePageHtml := "<!DOCTYPE html><html><body><h1>Voorbeeld QR-code Go Dev Tips</h1>" + qr_code_afbeelding_tag + "</body></html>"
	_, _ = reactie.Write([]byte(responsePageHtml))
}

func genereerQRCodeHtmlImageTag() string {

	qrCodeAfbeeldingsGegevens, _error := qrcode.Encode(websiteURL, qrcode.High, afbeeldingsGrootte)

	if _error != nil {
		log.Fatalln("Fout bij het genereren van QR-code. ", _error)
	}

	gecodeerdeGegevens := base64.StdEncoding.EncodeToString(qrCodeAfbeeldingsGegevens)

	return "<img src=\"data:image/png;base64, " + gecodeerdeGegevens + "\">"
}
