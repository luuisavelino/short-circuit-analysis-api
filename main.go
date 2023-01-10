package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/luuisavelino/short-circuit-analysis-elements/models"
	"github.com/luuisavelino/short-circuit-analysis-elements/routes"
)

func main() {
	reg, _ := regexp.Compile(`.*\.xlsx`)
	var i int = 0

	models.ElementTypes["0"] = models.ElementType{Info:"Todos os tipos"}
	models.ElementTypes["1"] = models.ElementType{Info:"Tipo 1"}
	models.ElementTypes["2"] = models.ElementType{Info:"Tipo 2 e 3"}

	readedFiles, _ := os.ReadDir("./files/")
	for _, readedFile := range readedFiles {
		if reg.MatchString(readedFile.Name()) {
			models.Files = append(models.Files, models.File{
				Posicao: i,
				Nome:    readedFile.Name(),
			})
			i++
		}
	}

	fmt.Println("Iniciando o servidor Rest com GO")
	routes.HandleRequests()
}
