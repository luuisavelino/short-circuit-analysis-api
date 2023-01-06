package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/luuisavelino/short-circuit-analysis-elements/models"
	"github.com/luuisavelino/short-circuit-analysis-elements/pkg/elements"
	"github.com/luuisavelino/short-circuit-analysis-elements/pkg/functions"
	"github.com/xuri/excelize/v2"
)

const (
	mensagemErroArquivo = "Arquivo não encontrado, tente novamente"
	path                  = "./files/"
)

func Readiness(w http.ResponseWriter, r *http.Request) {
	// Readiness probe do kubernetes
}

func Liveness(w http.ResponseWriter, r *http.Request) {
	// Liveness probe do kubernetes
}

func SystemSize(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileId, _ := strconv.Atoi(vars["fileId"])

	tabelaDados, err := excelize.OpenFile(path + models.Files[fileId].Nome)
	if err != nil {
		log.Fatal(mensagemErroArquivo)
	}

	systemSize, _ := functions.SystemInfo(tabelaDados)
	json.NewEncoder(w).Encode(systemSize)
}

func SystemBars(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileId, _ := strconv.Atoi(vars["fileId"])

	tabelaDados, err := excelize.OpenFile(path + models.Files[fileId].Nome)
	if err != nil {
		log.Fatal(mensagemErroArquivo)
	}

	_, bars := functions.SystemInfo(tabelaDados)
	json.NewEncoder(w).Encode(bars)
}

func AllElements(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileId, _ := strconv.Atoi(vars["fileId"])

	tabelaDados, err := excelize.OpenFile(path + models.Files[fileId].Nome)
	if err != nil {
		log.Fatal(mensagemErroArquivo)
	}

	models.Elements["1"] = elements.Elementos_tipo_1(tabelaDados)
	models.Elements["2"] = elements.Elementos_tipo_2_3(tabelaDados)

	json.NewEncoder(w).Encode(models.Elements)
}

func AllElementsType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileId, _ := strconv.Atoi(vars["fileId"])
	typeId := vars["typeId"]

	tabelaDados, err := excelize.OpenFile(path + models.Files[fileId].Nome)
	if err != nil {
		log.Fatal(mensagemErroArquivo)
	}

	models.Elements["1"] = elements.Elementos_tipo_1(tabelaDados)
	models.Elements["2"] = elements.Elementos_tipo_2_3(tabelaDados)

	json.NewEncoder(w).Encode(models.Elements[typeId])
}

func OneElement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileId, _ := strconv.Atoi(vars["fileId"])
	typeId := vars["typeId"]
	elementId, _ := strconv.Atoi(vars["elementId"])

	tabelaDados, err := excelize.OpenFile(path + models.Files[fileId].Nome)
	if err != nil {
		log.Fatal(mensagemErroArquivo)
	}

	models.Elements["1"] = elements.Elementos_tipo_1(tabelaDados)
	models.Elements["2"] = elements.Elementos_tipo_2_3(tabelaDados)

	for _, element := range models.Elements[typeId] {
		if element.Id == elementId {
			json.NewEncoder(w).Encode(element)
		}
	}
}

func AllFiles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Files)
}

func OneFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileId, _ := strconv.Atoi(vars["fileId"])

	for position, file := range models.Files {
		if position == fileId {
			json.NewEncoder(w).Encode(file)
		}
	}
}
