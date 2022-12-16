package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/luuisavelino/short-circuit-analysis-elements/models"
	"github.com/luuisavelino/short-circuit-analysis-elements/pkg/elements"
	"github.com/xuri/excelize/v2"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllElements(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileId, _ := strconv.Atoi(vars["fileId"])

	tabela_dados, err := excelize.OpenFile("./files/" + models.Files[fileId].Nome)
	if err != nil {
		log.Fatal("Arquivo não encontrado, tente novamente")
	}

	models.Elements["1"] = elements.Elementos_tipo_1(tabela_dados)
	models.Elements["2-3"] = elements.Elementos_tipo_2_3(tabela_dados)

	json.NewEncoder(w).Encode(models.Elements)
}

func AllElementsType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileId, _ := strconv.Atoi(vars["fileId"])
	typeId := vars["typeId"]

	tabela_dados, err := excelize.OpenFile("./files/" + models.Files[fileId].Nome)
	if err != nil {
		log.Fatal("Arquivo não encontrado, tente novamente")
	}

	models.Elements["1"] = elements.Elementos_tipo_1(tabela_dados)
	models.Elements["2"] = elements.Elementos_tipo_2_3(tabela_dados)

	json.NewEncoder(w).Encode(models.Elements[typeId])
}

func OneElement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileId, _ := strconv.Atoi(vars["fileId"])
	typeId := vars["typeId"]
	elementId, _ := strconv.Atoi(vars["elementId"])

	tabela_dados, err := excelize.OpenFile("./files/" + models.Files[fileId].Nome)
	if err != nil {
		log.Fatal("Arquivo não encontrado, tente novamente")
	}

	models.Elements["1"] = elements.Elementos_tipo_1(tabela_dados)
	models.Elements["2"] = elements.Elementos_tipo_2_3(tabela_dados)

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
