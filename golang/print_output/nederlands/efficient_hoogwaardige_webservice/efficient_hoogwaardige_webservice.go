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

type student struct {
	ID       int    `json:"id"`
	Naam     string `json:"naam"`
	Leeftijd int    `json:"leeftijd"`
}

type serverBericht struct {
	Bericht string `json:"bericht"`
}

var idTeller = 0
var studentenLijst []student

func main() {

	router := httprouter.New()
	router.GET("/student/alle", alleStudentenOphalen)
	router.GET("/student", studentOphalen)
	router.PATCH("/student/update", updateStudent)
	router.POST("/student", studentToevoegen)
	router.DELETE("/student", studentVerwijderen)

	serverFout := http.ListenAndServe(":3000", router)
	if serverFout != nil {
		log.Fatal("Kan webserver niet starten, oorzaak: ", serverFout)
	}
}

func alleStudentenOphalen(reactie http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	stelAlsJson(reactie)
	reactie.WriteHeader(http.StatusOK)

	if len(studentenLijst) == 0 {
		json_resultaat, _ := json.Marshal([]student{})
		_, schrijf_fout := reactie.Write(json_resultaat)
		if schrijf_fout != nil {
			interneServerfoutVerzenden(reactie)
		}
	} else {
		json_resultaat, _ := json.Marshal(studentenLijst)
		_, schrijf_fout := reactie.Write(json_resultaat)
		if schrijf_fout != nil {
			interneServerfoutVerzenden(reactie)
		}
	}
}

func studentOphalen(reactie http.ResponseWriter, verzoek *http.Request, _ httprouter.Params) {

	parameter := verzoek.URL.Query()

	if parameter == nil || len(parameter) == 0 || parameter["id"] == nil || parameter["id"][0] == "" {
		slechteReactieVerzoekenVerzenden(reactie, "ongeldig identiteit")
		return
	}

	studentID := parameter["id"][0]

	studentIdInt, err := strconv.Atoi(studentID)
	if err != nil {
		slechteReactieVerzoekenVerzenden(reactie, "ongeldig identiteit")
		return
	}

	var opgevraagdeStudent = student{
		ID:       -1,
		Naam:     "",
		Leeftijd: -1,
	}

	for _, studentGevonden := range studentenLijst {

		if studentGevonden.ID == studentIdInt {
			opgevraagdeStudent = studentGevonden
		}
	}

	if opgevraagdeStudent.ID == -1 {
		stuurReactieBericht(http.StatusNotFound, reactie, serverBericht{Bericht: "Student niet gevonden"})
	} else {
		reactieVerzenden(http.StatusOK, reactie, opgevraagdeStudent)
	}
}

func studentToevoegen(reactie http.ResponseWriter, verzoek *http.Request, _ httprouter.Params) {

	var nieuweStudent student
	verwerkingsFout := json.NewDecoder(verzoek.Body).Decode(&nieuweStudent)

	if verwerkingsFout != nil {
		if verwerkingsFout == io.EOF {
			slechteReactieVerzoekenVerzenden(reactie, "Onvolledige gegevens")
		} else {
			interneServerfoutVerzenden(reactie)
		}
		return
	} else if nieuweStudent.Naam == "" || nieuweStudent.Leeftijd == 0 {
		slechteReactieVerzoekenVerzenden(reactie, "Ontbrekende gegevens")
		return
	}

	lijstBevatData, studentGevonden := lijstBevatRecord(nieuweStudent.Naam, nieuweStudent.Leeftijd)

	if lijstBevatData {
		reactieVerzenden(http.StatusOK, reactie, studentGevonden)
	} else {
		idTeller++

		studentToevoegen := student{
			ID:       idTeller,
			Naam:     nieuweStudent.Naam,
			Leeftijd: nieuweStudent.Leeftijd,
		}
		studentenLijst = append(studentenLijst, studentToevoegen)
		reactieVerzenden(http.StatusOK, reactie, studentToevoegen)
	}
}

func studentVerwijderen(reactie http.ResponseWriter, verzoek *http.Request, _ httprouter.Params) {

	parameter := verzoek.URL.Query()

	if parameter == nil || len(parameter) == 0 || parameter["id"] == nil || parameter["id"][0] == "" {
		slechteReactieVerzoekenVerzenden(reactie, "ongeldig identiteit")
		return
	}

	studentID := parameter["id"][0]

	studentIdInt, err := strconv.Atoi(studentID)
	if err != nil {
		slechteReactieVerzoekenVerzenden(reactie, "ongeldig identiteit")
		return
	}

	for index, studentFound := range studentenLijst {
		if studentFound.ID == studentIdInt {
			studentenLijst = append(studentenLijst[:index], studentenLijst[index+1:]...)
		}
	}
	stuurReactieBericht(http.StatusOK, reactie, serverBericht{Bericht: "Student verwijderd"})
}

func updateStudent(reactie http.ResponseWriter, verzoek *http.Request, _ httprouter.Params) {

	parameter := verzoek.URL.Query()

	if parameter == nil || len(parameter) == 0 || parameter["id"] == nil || parameter["id"][0] == "" {
		slechteReactieVerzoekenVerzenden(reactie, "ongeldig identiteit")
		return
	}

	studentID := parameter["id"][0]

	studentIdInt, err := strconv.Atoi(studentID)
	if err != nil {
		slechteReactieVerzoekenVerzenden(reactie, "ongeldig identiteit")
		return
	}

	lijstBevatData := lijstBevatRecordDoorID(studentIdInt)

	var geupdateStudent student
	parseError := json.NewDecoder(verzoek.Body).Decode(&geupdateStudent)

	if !lijstBevatData {
		stuurReactieBericht(http.StatusNotFound, reactie, serverBericht{Bericht: "Student niet gevonden"})
		return
	}

	if parseError != nil {
		if parseError == io.EOF {
			slechteReactieVerzoekenVerzenden(reactie, "Onvolledige gegevens")
		} else {
			interneServerfoutVerzenden(reactie)
		}
		return
	}

	for index, studentGevonden := range studentenLijst {

		if studentGevonden.ID == studentIdInt {

			studentOmTeUpdaten := &studentenLijst[index]

			if studentOmTeUpdaten.Naam != "" {
				studentOmTeUpdaten.Naam = geupdateStudent.Naam
			}

			if studentOmTeUpdaten.Leeftijd != 0 {
				studentOmTeUpdaten.Leeftijd = geupdateStudent.Leeftijd
			}

			geupdateStudent = *studentOmTeUpdaten
		}
	}

	reactieVerzenden(http.StatusOK, reactie, geupdateStudent)
}

func lijstBevatRecord(naam string, leeftijd int) (bool, student) {

	bevatGegevens := false
	var student student

	for _, studentGevonden := range studentenLijst {
		if strings.ToLower(studentGevonden.Naam) == strings.ToLower(naam) && studentGevonden.Leeftijd == leeftijd {
			bevatGegevens = true
			student = studentGevonden
		}
	}

	return bevatGegevens, student
}

func lijstBevatRecordDoorID(id int) bool {

	bevatGegevens := false

	for _, studentGevonden := range studentenLijst {
		if studentGevonden.ID == id {
			bevatGegevens = true
		}
	}

	return bevatGegevens
}

func reactieVerzenden(statusCode int, reactie http.ResponseWriter, student student) {

	if statusCode != 0 && reactie != nil {
		stelAlsJson(reactie)
		reactie.WriteHeader(statusCode)
		jsonReactie, jsonFout := json.Marshal(student)
		if jsonFout != nil {
			interneServerfoutVerzenden(reactie)
		} else {
			_, verzendingFout := reactie.Write(jsonReactie)
			if verzendingFout != nil {
				interneServerfoutVerzenden(reactie)
			}
		}
	}
}

func stuurReactieBericht(statusCode int, reactie http.ResponseWriter, bericht serverBericht) {

	if statusCode != 0 && reactie != nil {
		stelAlsJson(reactie)
		reactie.WriteHeader(statusCode)
		jsonReactie, jsonFout := json.Marshal(bericht)
		if jsonFout != nil {
			interneServerfoutVerzenden(reactie)
		} else {
			_, verzendingFout := reactie.Write(jsonReactie)
			if verzendingFout != nil {
				interneServerfoutVerzenden(reactie)
			}
		}
	}
}

func interneServerfoutVerzenden(reactie http.ResponseWriter) {

	stelAlsJson(reactie)
	reactie.WriteHeader(http.StatusInternalServerError)
	serverFoutBericht := serverBericht{Bericht: "Kan verzoek niet verwerken"}
	jsonReactie, jsonFout := json.Marshal(serverFoutBericht)

	if jsonFout != nil {
		_, _ = reactie.Write([]byte(""))
	} else {
		_, _ = reactie.Write(jsonReactie)
	}
}

func slechteReactieVerzoekenVerzenden(reactie http.ResponseWriter, bericht string) {

	stelAlsJson(reactie)
	reactie.WriteHeader(http.StatusBadRequest)
	serverFoutBericht := serverBericht{Bericht: bericht}
	jsonReactie, jsonFout := json.Marshal(serverFoutBericht)

	if jsonFout != nil {
		_, _ = reactie.Write([]byte(""))
	} else {
		_, _ = reactie.Write(jsonReactie)
	}
}

func stelAlsJson(responseWriter http.ResponseWriter) {
	responseWriter.Header().Set("Content-Type", "application/json")
}
