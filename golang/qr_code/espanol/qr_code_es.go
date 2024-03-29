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

var urlDelSitioWeb = "https://www.godevtips.com/es" // URL de GoDevTips
var tamanoDeImagen = 256                            // 256 x 256 píxeles

func main() {

	if len(os.Args) == 2 && os.Args[1] == "archivo" { // Determine cuál opción debe lanzarse

		// Genere solo los bytes PNG sin procesar
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

	enrutador := httprouter.New()      // instancia de enrutador
	enrutador.GET("/", enviarCodigoQR) // página de destino

	fmt.Println("servidor escuchando atraves del puerto 3000")
	errorDelServidor := http.ListenAndServe(":3000", enrutador)
	if errorDelServidor != nil {
		log.Fatal("No se puede iniciar el servidor web, causa: ", errorDelServidor)
	}
}

func enviarCodigoQR(respuesta http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	// Generar etiqueta de imagen HTML de código QR
	etiqueta_de_imagen_de_codigo_qr := generarEtiquetaDeImagenHtmlDeCodigoQR()

	// Establecer el encabezado de respuesta en Estado Ok (200)
	respuesta.WriteHeader(http.StatusOK)

	// Imagen HTML con código QR incrustado
	paginaDeRepuestaHtml := "<!DOCTYPE html><html><body><h1>Ejemplo de código QR Go Dev Tips</h1>" + etiqueta_de_imagen_de_codigo_qr + "</body></html>"

	// Enviar respuesta HTML
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
