package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type estudiante struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad"`
}

type mensajeDeRespuesta struct {
	Mensaje string `json:"mensaje"`
}

var idContador = 0
var listaDeEstudiantes []estudiante

func main() {

	router := httprouter.New()
	router.GET("/estudiante/todos", manejarObtenerTodosLosEstudiantes)
	router.GET("/estudiante", manejarObtenerEstudiantes)
	router.PATCH("/estudiante/actualizar", manejarActualizacionEstudiante)
	router.POST("/estudiante", manejarAgregarEstudiante)
	router.DELETE("/estudiante", manejarEliminarEstudiante)

	errorDelServidor := http.ListenAndServe(":3000", router)
	if errorDelServidor != nil {
		log.Fatal("No se puede iniciar el servidor web, causa: ", errorDelServidor)
	}
}

func manejarObtenerTodosLosEstudiantes(respuesta http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	asignadoComoJson(respuesta)
	respuesta.WriteHeader(http.StatusOK)

	if len(listaDeEstudiantes) == 0 {
		respuestaJson, _ := json.Marshal([]estudiante{})
		_, error := respuesta.Write(respuestaJson)
		if error != nil {
			enviarErrorInternoDelServidor(respuesta)
		}
	} else {
		respuestaJson, _ := json.Marshal(listaDeEstudiantes)
		_, error := respuesta.Write(respuestaJson)
		if error != nil {
			enviarErrorInternoDelServidor(respuesta)
		}
	}
}

func manejarObtenerEstudiantes(respuesta http.ResponseWriter, peticion *http.Request, _ httprouter.Params) {

	parametros := peticion.URL.Query()

	if parametros == nil || len(parametros) == 0 || parametros["id"] == nil || parametros["id"][0] == "" {
		RepuestaDeSolicitudIncorrecta(respuesta, "identificación invalida")
		return
	}

	identificacionDelEstudiante := parametros["id"][0]

	estudianteIDEnFormaEntero, error := strconv.Atoi(identificacionDelEstudiante)
	if error != nil {
		RepuestaDeSolicitudIncorrecta(respuesta, "identificación invalida")
		return
	}

	var estudianteSolicitado = estudiante{
		ID:     -1,
		Nombre: "",
		Edad:   -1,
	}

	for _, estudianteEncontrado := range listaDeEstudiantes {

		if estudianteEncontrado.ID == estudianteIDEnFormaEntero {
			estudianteSolicitado = estudianteEncontrado
		}
	}

	if estudianteSolicitado.ID == -1 {
		enviarRespuestaComoUnMensaje(http.StatusNotFound, respuesta, mensajeDeRespuesta{Mensaje: "Registro no encontrado"})
	} else {
		enviarRepuesta(http.StatusOK, respuesta, estudianteSolicitado)
	}
}

func manejarAgregarEstudiante(respuesta http.ResponseWriter, peticion *http.Request, _ httprouter.Params) {

	var nuevoEstudiante estudiante
	error := json.NewDecoder(peticion.Body).Decode(&nuevoEstudiante)

	if error != nil {
		if error == io.EOF {
			RepuestaDeSolicitudIncorrecta(respuesta, "Datos incompletos")
		} else {
			enviarErrorInternoDelServidor(respuesta)
		}
		return
	} else if nuevoEstudiante.Nombre == "" || nuevoEstudiante.Edad == 0 {
		RepuestaDeSolicitudIncorrecta(respuesta, "Datos perdidos")
		return
	}

	listaContieneDatos, estudianteEncontrado := listaContieneRegistro(nuevoEstudiante.Nombre, nuevoEstudiante.Edad)

	if listaContieneDatos {
		enviarRepuesta(http.StatusOK, respuesta, estudianteEncontrado)
	} else {
		idContador++

		alumnoParaAgregar := estudiante{
			ID:     idContador,
			Nombre: nuevoEstudiante.Nombre,
			Edad:   nuevoEstudiante.Edad,
		}
		listaDeEstudiantes = append(listaDeEstudiantes, alumnoParaAgregar)
		enviarRepuesta(http.StatusOK, respuesta, alumnoParaAgregar)
	}
}

func manejarEliminarEstudiante(respuesta http.ResponseWriter, peticion *http.Request, _ httprouter.Params) {

	parametro := peticion.URL.Query()

	if parametro == nil || len(parametro) == 0 || parametro["id"] == nil || parametro["id"][0] == "" {
		RepuestaDeSolicitudIncorrecta(respuesta, "Identificación invalida")
		return
	}

	identificacionDelEstudiante := parametro["id"][0]

	estudianteIDEnFormaEntero, err := strconv.Atoi(identificacionDelEstudiante)
	if err != nil {
		RepuestaDeSolicitudIncorrecta(respuesta, "Identificación invalida")
		return
	}

	for indice, studentFound := range listaDeEstudiantes {
		if studentFound.ID == estudianteIDEnFormaEntero {
			listaDeEstudiantes = append(listaDeEstudiantes[:indice], listaDeEstudiantes[indice+1:]...)
		}
	}
	enviarRespuestaComoUnMensaje(http.StatusOK, respuesta, mensajeDeRespuesta{Mensaje: "Estudiante eliminado"})
}

func manejarActualizacionEstudiante(respuesta http.ResponseWriter, peticion *http.Request, _ httprouter.Params) {

	parametro := peticion.URL.Query()

	if parametro == nil || len(parametro) == 0 || parametro["id"] == nil || parametro["id"][0] == "" {
		RepuestaDeSolicitudIncorrecta(respuesta, "Identificación invalida")
		return
	}

	identificacionDelEstudiante := parametro["id"][0]

	identificacionDelEstudianteEnEntero, err := strconv.Atoi(identificacionDelEstudiante)
	if err != nil {
		RepuestaDeSolicitudIncorrecta(respuesta, "Identificación invalida")
		return
	}

	listaContieneDatos := listaContieneRegistroPorID(identificacionDelEstudianteEnEntero)

	var estudianteActualizado estudiante
	errorDeAnalisis := json.NewDecoder(peticion.Body).Decode(&estudianteActualizado)

	if !listaContieneDatos {
		enviarRespuestaComoUnMensaje(http.StatusNotFound, respuesta, mensajeDeRespuesta{Mensaje: "Registro no encontrado!"})
		return
	}

	if errorDeAnalisis != nil {
		if errorDeAnalisis == io.EOF {
			RepuestaDeSolicitudIncorrecta(respuesta, "Datos incompletos")
		} else {
			enviarErrorInternoDelServidor(respuesta)
		}
		return
	}

	for indice, estudianteEncontrado := range listaDeEstudiantes {

		if estudianteEncontrado.ID == identificacionDelEstudianteEnEntero {

			estudianteParaActualizar := &listaDeEstudiantes[indice]

			if estudianteParaActualizar.Nombre != "" {
				estudianteParaActualizar.Nombre = estudianteActualizado.Nombre
			}

			if estudianteParaActualizar.Edad != 0 {
				estudianteParaActualizar.Edad = estudianteActualizado.Edad
			}

			estudianteActualizado = *estudianteParaActualizar
		}
	}

	enviarRepuesta(http.StatusOK, respuesta, estudianteActualizado)
}

func listaContieneRegistro(nombre string, edad int) (bool, estudiante) {

	contieneDatos := false
	var estudiante estudiante

	for _, estudianteEncontrado := range listaDeEstudiantes {
		if strings.ToLower(estudianteEncontrado.Nombre) == strings.ToLower(nombre) && estudianteEncontrado.Edad == edad {
			contieneDatos = true
			estudiante = estudianteEncontrado
		}
	}

	return contieneDatos, estudiante
}

func listaContieneRegistroPorID(identificacion int) bool {

	contieneDatos := false

	for _, estudianteEncontrado := range listaDeEstudiantes {
		if estudianteEncontrado.ID == identificacion {
			contieneDatos = true
		}
	}

	return contieneDatos
}

func enviarRepuesta(codigoDeEstado int, respuesta http.ResponseWriter, estudiante estudiante) {

	if codigoDeEstado != 0 && respuesta != nil {
		asignadoComoJson(respuesta)
		respuesta.WriteHeader(codigoDeEstado)
		repuestaEnJson, errorJson := json.Marshal(estudiante)
		if errorJson != nil {
			enviarErrorInternoDelServidor(respuesta)
		} else {
			_, respuestaJsonError := respuesta.Write(repuestaEnJson)
			if respuestaJsonError != nil {
				enviarErrorInternoDelServidor(respuesta)
			}
		}
	}
}

func enviarRespuestaComoUnMensaje(codigoDeEstado int, respuesta http.ResponseWriter, mensaje mensajeDeRespuesta) {

	if codigoDeEstado != 0 && respuesta != nil {
		asignadoComoJson(respuesta)
		respuesta.WriteHeader(codigoDeEstado)

		repuestaEnJson, errorJson := json.Marshal(mensaje)
		if errorJson != nil {
			enviarErrorInternoDelServidor(respuesta)
		} else {
			_, repuestaJsonError := respuesta.Write(repuestaEnJson)
			if repuestaJsonError != nil {
				enviarErrorInternoDelServidor(respuesta)
			}
		}
	}
}

func enviarErrorInternoDelServidor(respuesta http.ResponseWriter) {

	asignadoComoJson(respuesta)
	respuesta.WriteHeader(http.StatusInternalServerError)
	mensajeRepuesta := mensajeDeRespuesta{Mensaje: "No se puede manejar la solicitud"}
	respuestaEnJson, errorJson := json.Marshal(mensajeRepuesta)

	if errorJson != nil {
		_, _ = respuesta.Write([]byte(""))
	} else {
		_, _ = respuesta.Write(respuestaEnJson)
	}
}

func RepuestaDeSolicitudIncorrecta(respuesta http.ResponseWriter, mensaje string) {

	asignadoComoJson(respuesta)
	respuesta.WriteHeader(http.StatusBadRequest)
	mensajeRepuesta := mensajeDeRespuesta{Mensaje: mensaje}
	respuestaEnJson, errorJson := json.Marshal(mensajeRepuesta)

	if errorJson != nil {
		_, _ = respuesta.Write([]byte(""))
	} else {
		_, _ = respuesta.Write(respuestaEnJson)
	}
}

func asignadoComoJson(escritorDeRespuesta http.ResponseWriter) {
	escritorDeRespuesta.Header().Set("Content-Type", "application/json")
}
