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

var websiteURL = "https://www.godevtips.com" // GoDevTips URL
var imageSize = 256                          // 256 x 256 pixels

func main() {

	if len(os.Args) == 2 && os.Args[1] == "file" { // Determine which is option to launch

		// Generate only the raw PNG bytes
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

	router := httprouter.New()  // router instance
	router.GET("/", sendQRCode) // landing page

	fmt.Println("server listening on port 3000")
	serverError := http.ListenAndServe(":3000", router)
	if serverError != nil {
		log.Fatal("Unable to start web server, cause: ", serverError)
	}
}

func sendQRCode(response http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	// Generate QR code HTML image tag
	qr_code_image_tag := generateQRCodeHtmlImageTag()
	// Set response header to Status Ok (200)
	response.WriteHeader(http.StatusOK)

	// HTML image with embedded QR code
	responsePageHtml := "<!DOCTYPE html><html><body><h1>QR Code example Go Dev Tips</h1>" + qr_code_image_tag + "</body></html>"

	// Send HTML response back to client
	_, _ = response.Write([]byte(responsePageHtml))
}

func generateQRCodeHtmlImageTag() string {

	qrCodeImageData, taskError := qrcode.Encode(websiteURL, qrcode.High, imageSize)

	if taskError != nil {
		log.Fatalln("Error generating QR code. ", taskError)
	}

	// // Encode raw QR code data to base 64
	encodedData := base64.StdEncoding.EncodeToString(qrCodeImageData)

	return "<img src=\"data:image/png;base64, " + encodedData + "\">"
}
