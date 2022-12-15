package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/luuisavelino/short-circuit-analysis-api/models"
	"github.com/luuisavelino/short-circuit-analysis-api/pkg/barra"
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

	models.Elements = barra.Elementos_tipo_2_3(tabela_dados)

	json.NewEncoder(w).Encode(models.Elements)
}

func OneElement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	line := vars["line"]

	for _, element := range models.Elements {
		if element.De+"-"+element.Para == line {
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
