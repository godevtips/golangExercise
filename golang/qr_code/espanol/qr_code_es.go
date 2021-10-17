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

var urlDelSitioWeb = "https://www.godevtips.com/es"
var tamanoDeImagen = 256

func main() {

	if len(os.Args) == 2 && os.Args[1] == "archivo" { // Determine cuál opción que debe lanzarse

		nombreDelArchivo := "qr_code.png"
		_error := qrcode.WriteFile(urlDelSitioWeb, qrcode.Medium, tamanoDeImagen, nombreDelArchivo)

		if _error != nil {
			log.Fatal("Error generar el código QR. ", _error)
		} else {
			fmt.Println("¡Archivo de código QR creado!")
		}
	} else {
		iniciarServidor()
	}
}

func iniciarServidor() {

	enrutador := httprouter.New()
	enrutador.GET("/", enviarCodigoQR)

	fmt.Println("servidor escuchando atraves del puerto 3000")
	errorDelServidor := http.ListenAndServe(":3000", enrutador)
	if errorDelServidor != nil {
		log.Fatal("No se puede iniciar el servidor web, causa: ", errorDelServidor)
	}
}

func enviarCodigoQR(respuesta http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	etiqueta_de_imagen_de_codigo_qr := generarEtiquetaDeImagenHtmlDeCodigoQR()
	respuesta.WriteHeader(http.StatusOK)
	paginaDeRepuestaHtml := "<!DOCTYPE html><html><body><h1>Ejemplo de código QR Go Dev Tips</h1>" + etiqueta_de_imagen_de_codigo_qr + "</body></html>"
	_, _ = respuesta.Write([]byte(paginaDeRepuestaHtml))
}

func generarEtiquetaDeImagenHtmlDeCodigoQR() string {

	//Datos de imagen de código qr
	datosDeImagenDeCodigoQR, _error := qrcode.Encode(urlDelSitioWeb, qrcode.High, tamanoDeImagen)

	if _error != nil {
		log.Fatalln("Error generar el código QR. ", _error)
	}

	datosCodificados := base64.StdEncoding.EncodeToString(datosDeImagenDeCodigoQR)

	return "<img src=\"data:image/png;base64, " + datosCodificados + "\">"
}
