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

const trayectoria_del_projecto = "./golang/noisy_image/espanol"

func main() {

	err := eliminarArchivo("resultado.png")
	if err != nil {
		fmt.Println(err)
	}

	imagen_original, err := obtenerUnaImagenJpeg("imagen")

	if err != nil {
		log.Fatal(err)
	}

	// Imagen superpuesta con ruido
	imagenSuperpuestaConRuido, err := crearSuperpuestoRuidoso(imagen_original, 50)

	if err != nil {
		log.Fatal(err)
	}

	errorDeImagenResultante := fusionarYCrearUnaImagenConRuido(imagen_original, imagenSuperpuestaConRuido)

	if errorDeImagenResultante != nil {
		log.Fatal(errorDeImagenResultante)
	}
}

func eliminarArchivo(archivo string) error {

	// eliminar el archivo si existe utilizando la función Remove()
	resultadoErrorDeEliminacion := os.Remove(fmt.Sprintf("%s/%s", trayectoria_del_projecto, archivo))
	if resultadoErrorDeEliminacion != nil {
		return resultadoErrorDeEliminacion
	}

	return nil
}

func obtenerUnaImagenJpeg(nombreDeArchivo string) (image.Image, error) {

	archivoDeImagen, err := os.Open(fmt.Sprintf("%s/%s.jpg", trayectoria_del_projecto, nombreDeArchivo))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error al abrir el archivo original: %s", err))
	}

	imagenDecodificada, err := jpeg.Decode(archivoDeImagen)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("no se ha podido decodificar la imagen original: %s", err))
	}

	defer func(archivo_de_imagen *os.File) {
		_ = archivo_de_imagen.Close()
	}(archivoDeImagen)

	return imagenDecodificada, nil
}

func crearSuperpuestoRuidoso(imagenDeFondo image.Image, opacidad int) (image.Image, error) {

	var opacidad_solicitada = uint8(0)
	if opacidad > 0 && opacidad < 256 {
		opacidad_solicitada = uint8(math.Round(float64(opacidad) * 2.56))
	}

	encuadernacionDeUnaImagen := imagenDeFondo.Bounds()
	anchuraDeImagen := encuadernacionDeUnaImagen.Dx()
	alturaDeImagen := encuadernacionDeUnaImagen.Dy()
	imagenSuperponer := image.NewRGBA(image.Rect(0, 0, anchuraDeImagen, alturaDeImagen))

	for pixel := 0; pixel < anchuraDeImagen*alturaDeImagen; pixel++ {
		desplazamientoDePixeles := 4 * pixel
		imagenSuperponer.Pix[0+desplazamientoDePixeles] = uint8(rand.Intn(256))
		imagenSuperponer.Pix[1+desplazamientoDePixeles] = uint8(rand.Intn(256))
		imagenSuperponer.Pix[2+desplazamientoDePixeles] = uint8(rand.Intn(256))
		imagenSuperponer.Pix[3+desplazamientoDePixeles] = opacidad_solicitada
	}

	archivoDelImagenConRuidoTemporal, png_err := os.Create(fmt.Sprintf("%s/%s", trayectoria_del_projecto, "imagen_ruidosa_temporal.png"))
	if png_err != nil {
		log.Fatal(png_err)
	}

	error_de_codification := png.Encode(archivoDelImagenConRuidoTemporal, imagenSuperponer)
	if error_de_codification != nil {
		return nil, error_de_codification
	}

	imagenRuidosaTemporal, err := os.Open(fmt.Sprintf("%s/%s", trayectoria_del_projecto, "imagen_ruidosa_temporal.png"))
	if err != nil {
		return nil, err
	}

	// Error de decodificación png
	resultado, errorDecodificacionPng := png.Decode(imagenRuidosaTemporal)
	if errorDecodificacionPng != nil {
		return nil, errorDecodificacionPng
	}

	defer func(imagen_rudiosa_temporal *os.File) {
		_ = imagen_rudiosa_temporal.Close()
	}(imagenRuidosaTemporal)

	// eliminar imagen_ruidosa_temporal.png con la función Remove()
	err = eliminarArchivo("imagen_ruidosa_temporal.png")
	if err != nil {
		return nil, err
	}

	return resultado, nil
}

func fusionarYCrearUnaImagenConRuido(imagen_original image.Image, imagen_de_superposicion_de_ruido image.Image) error {

	compensacion := image.Pt(0, 0)
	limiteDeLaImagenOriginal := imagen_original.Bounds()
	imagenRuidosa := image.NewRGBA(limiteDeLaImagenOriginal)
	draw.Draw(imagenRuidosa, limiteDeLaImagenOriginal, imagen_original, image.Point{}, draw.Src)
	draw.Draw(imagenRuidosa, imagen_de_superposicion_de_ruido.Bounds().Add(compensacion), imagen_de_superposicion_de_ruido, image.Point{}, draw.Over)

	imagenResultante, err := os.Create(fmt.Sprintf("%s/%s", trayectoria_del_projecto, "resultado.png"))
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create result image: %s", err))
	}

	errorDeCodificacionPng := png.Encode(imagenResultante, imagenRuidosa)
	if errorDeCodificacionPng != nil {
		return errorDeCodificacionPng
	}

	defer func(archivo_de_imagen *os.File) {
		_ = archivo_de_imagen.Close()
	}(imagenResultante)

	return nil
}
