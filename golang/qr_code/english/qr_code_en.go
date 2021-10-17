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
var imageSize = 256

func main() {

	if len(os.Args) == 2 && os.Args[1] == "file" { // Determine which is option must be launch

		outputFileName := "qr_code.png"
		taskError := qrcode.WriteFile(websiteURL, qrcode.Medium, imageSize, outputFileName)

		if taskError != nil {
			log.Fatal("Error generating QR code to file. ", taskError)
		} else {
			fmt.Println("Qrcode file created!")
		}
	} else {
		startServer()
	}
}

func startServer() {

	router := httprouter.New()
	router.GET("/", sendQRCode)

	fmt.Println("server listening on port 3000")
	serverError := http.ListenAndServe(":3000", router)
	if serverError != nil {
		log.Fatal("Unable to start web server, cause: ", serverError)
	}
}

func sendQRCode(response http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	qr_code_image_tag := generateQRCodeHtmlImageTag()
	response.WriteHeader(http.StatusOK)
	responsePageHtml := "<!DOCTYPE html><html><body><h1>QR Code example Go Dev Tips</h1>" + qr_code_image_tag + "</body></html>"
	_, _ = response.Write([]byte(responsePageHtml))
}

func generateQRCodeHtmlImageTag() string {

	qrCodeImageData, taskError := qrcode.Encode(websiteURL, qrcode.High, imageSize)

	if taskError != nil {
		log.Fatalln("Error generating QR code. ", taskError)
	}

	encodedData := base64.StdEncoding.EncodeToString(qrCodeImageData)

	return "<img src=\"data:image/png;base64, " + encodedData + "\">"
}
