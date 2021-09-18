package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type errorMessage struct {
	Message string `json:"message"`
}

var idCounter = 0
var studentList []student

func main() {

	router := httprouter.New()
	router.GET("/student/all", handleGetAllStudent)
	router.GET("/student", handleGetStudent)
	router.PATCH("/student/update", updateStudent)
	router.POST("/student", addStudent)

	serverError := http.ListenAndServe(":3000", router)
	if serverError != nil {
		log.Fatal("Unable to start web server, cause: ", serverError)
	}
}

func handleGetAllStudent(response http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	mappedAsJson(response)
	response.WriteHeader(http.StatusOK)

	if len(studentList) == 0 {
		marshalledResponse, _ := json.Marshal([]student{})
		_, writerError := response.Write(marshalledResponse)
		if writerError != nil {
			fmt.Println(writerError)
			writeInternalServerError(response)
		}
	} else {
		marshalledResponse, _ := json.Marshal(studentList)
		_, responseWriteError := response.Write(marshalledResponse)
		if responseWriteError != nil {
			fmt.Println(responseWriteError)
			writeInternalServerError(response)
		}
	}
}

func handleGetStudent(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	params := request.URL.Query()

	if params == nil || len(params) == 0 || params["id"] == nil || params["id"][0] == "" {
		BadRequestResponse(response, "Invalid ID")
		return
	}

	studentID := params["id"][0]

	studentIdInt, err := strconv.Atoi(studentID)
	if err != nil {
		BadRequestResponse(response, "Invalid ID")
		return
	}

	var requestedStudent = student{
		ID:   -1,
		Name: "",
		Age:  -1,
	}

	for _, studentFound := range studentList {

		if studentFound.ID == studentIdInt {
			requestedStudent = studentFound
		}
	}

	if requestedStudent.ID == -1 {
		sendResponseMessage(http.StatusNotFound, response, errorMessage{Message: "Record not found"})
	} else {
		sendResponse(http.StatusOK, response, requestedStudent)
	}
}

func addStudent(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	var newStudent student
	parseError := json.NewDecoder(request.Body).Decode(&newStudent)

	if parseError != nil {
		if parseError == io.EOF {
			BadRequestResponse(response, "Incomplete data")
		} else {
			writeInternalServerError(response)
		}
		return
	} else if newStudent.Name == "" || newStudent.Age == 0 {
		BadRequestResponse(response, "Missing data")
		return
	}

	listContainsData, studentFound := listContainsRecord(newStudent.Name, newStudent.Age)

	if listContainsData {
		sendResponse(http.StatusOK, response, studentFound)
	} else {
		idCounter++

		studentToAdd := student{
			ID:   idCounter,
			Name: newStudent.Name,
			Age:  newStudent.Age,
		}
		studentList = append(studentList, studentToAdd)
		sendResponse(http.StatusOK, response, studentToAdd)
	}
}

func updateStudent(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	params := request.URL.Query()

	if params == nil || len(params) == 0 || params["id"] == nil || params["id"][0] == "" {
		BadRequestResponse(response, "Invalid ID")
		return
	}

	studentID := params["id"][0]

	studentIdInt, err := strconv.Atoi(studentID)
	if err != nil {
		BadRequestResponse(response, "Invalid ID")
		return
	}

	listContainsData := listContainsRecordByID(studentIdInt)

	var updatedStudent student
	parseError := json.NewDecoder(request.Body).Decode(&updatedStudent)

	if !listContainsData {
		sendResponseMessage(http.StatusNotFound, response, errorMessage{Message: "Record not found!"})
		return
	}

	if parseError != nil {
		if parseError == io.EOF {
			BadRequestResponse(response, "Incomplete data")
		} else {
			writeInternalServerError(response)
		}
		return
	}

	for index, studentFound := range studentList {

		if studentFound.ID == studentIdInt {

			studentToUpdate := &studentList[index] // point to the address of the item

			if updatedStudent.Name != "" {
				studentToUpdate.Name = updatedStudent.Name
			}

			if updatedStudent.Age != 0 {
				studentToUpdate.Age = updatedStudent.Age
			}

			updatedStudent = *studentToUpdate
		}
	}

	sendResponse(http.StatusOK, response, updatedStudent)
}

func listContainsRecord(name string, age int) (bool, student) {

	containsData := false
	var student student

	for _, studentFound := range studentList {
		if strings.ToLower(studentFound.Name) == strings.ToLower(name) && studentFound.Age == age {
			containsData = true
			student = studentFound
		}
	}

	return containsData, student
}

func listContainsRecordByID(id int) bool {

	containsData := false

	for _, studentFound := range studentList {
		if studentFound.ID == id {
			containsData = true
		}
	}

	return containsData
}

func sendResponse(statusCode int, response http.ResponseWriter, student student) {

	if statusCode != 0 && response != nil {
		mappedAsJson(response)
		response.WriteHeader(statusCode)
		jsonResponse, jsonMarshalError := json.Marshal(student)
		if jsonMarshalError != nil {
			writeInternalServerError(response)
		} else {
			_, jsonResponseError := response.Write(jsonResponse)
			if jsonResponseError != nil {
				writeInternalServerError(response)
			}
		}
	}
}

func sendResponseMessage(statusCode int, response http.ResponseWriter, message errorMessage) {

	if statusCode != 0 && response != nil {
		mappedAsJson(response)
		response.WriteHeader(statusCode)
		jsonResponse, jsonMarshalError := json.Marshal(message)
		if jsonMarshalError != nil {
			writeInternalServerError(response)
		} else {
			_, jsonResponseError := response.Write(jsonResponse)
			if jsonResponseError != nil {
				writeInternalServerError(response)
			}
		}
	}
}

func writeInternalServerError(response http.ResponseWriter) {

	mappedAsJson(response)
	response.WriteHeader(http.StatusInternalServerError)
	errorResponse := errorMessage{Message: "Unable to handle request"}
	errorResponseJson, marshalError := json.Marshal(errorResponse)

	if marshalError != nil {
		_, _ = response.Write([]byte(""))
	} else {
		_, _ = response.Write(errorResponseJson)
	}
}

func BadRequestResponse(response http.ResponseWriter, message string) {

	mappedAsJson(response)
	response.WriteHeader(http.StatusBadRequest)
	errorResponse := errorMessage{Message: message}
	errorResponseJson, marshalError := json.Marshal(errorResponse)

	if marshalError != nil {
		_, _ = response.Write([]byte(""))
	} else {
		_, _ = response.Write(errorResponseJson)
	}
}

func mappedAsJson(responseWriter http.ResponseWriter) {
	responseWriter.Header().Set("Content-Type", "application/json")
}
