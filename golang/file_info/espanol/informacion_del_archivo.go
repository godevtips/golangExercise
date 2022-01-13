package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {

	// File example
	informacionDelArchivo, err := os.Stat("./golang/file_info/espanol/directorio_del_archivo/archivo.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("|----------- Información del archivo -----------|")
	fmt.Println("Nombre del archivo:     ", informacionDelArchivo.Name())
	fmt.Println("Tamaño:                 ", obtenerTamanoDeArchivo(informacionDelArchivo))
	fmt.Println("Permisos:               ", informacionDelArchivo.Mode())
	fmt.Println("Última modificación:    ", informacionDelArchivo.ModTime())
	fmt.Println("Es Directorio:          ", esDirectorio(informacionDelArchivo))
	fmt.Println("|----------- Información del archivo -----------|\n")

	// Además, si desea obtener la fuente de datos subyacente, puede usar:
	// fileInfo.Sys(). Puede devolver nil en algunos casos.

	// Ejemplo de directorio
	informacionDelDirectorio, err := os.Stat("./golang/file_info/espanol/directorio_del_archivo")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("|----------- Información del directorio -----------|")
	fmt.Println("Nombre del archivo:     ", informacionDelDirectorio.Name())
	fmt.Println("Tamaño:                 ", obtenerTamanoDeArchivo(informacionDelDirectorio))
	fmt.Println("Permisos:               ", informacionDelDirectorio.Mode())
	fmt.Println("Última modificación:    ", informacionDelDirectorio.ModTime())
	fmt.Println("Es Directorio:          ", esDirectorio(informacionDelDirectorio))
	fmt.Println("|----------- Información del directorio -----------|")

	//Igual que en el ejemplo anterior, si desea obtener la fuente de datos subyacente, puede usar:
	//directoryInfo.Sys(). Puede devolver nil en algunos casos.
}

func esDirectorio(fileInfo fs.FileInfo) string {
	var resultado = "No"
	if fileInfo.IsDir() {
		resultado = "Sí"
	}
	return resultado
}

func obtenerTamanoDeArchivo(fileInfo fs.FileInfo) string {

	tamano := fileInfo.Size()
	if tamano < 1000 { // El tamaño del archivo está en bytes
		return fmt.Sprintf("%d bytes", tamano)
	} else { // el tamaño del archivo está en kilobytes o megabytes
		tamanoDelArchivoEnKb := float64(tamano) / float64(1000)
		if tamanoDelArchivoEnKb > 1000 { // el tamaño del archivo está en megabytes
			tamanoDelArchivoEnMb := tamanoDelArchivoEnKb / float64(1000)
			return fmt.Sprintf("%d Mb", tamanoDelArchivoEnMb)
		} else { // el tamaño del archivo está en kilobytes
			return fmt.Sprintf("%d Kb", tamanoDelArchivoEnKb)
		}
	}
}
