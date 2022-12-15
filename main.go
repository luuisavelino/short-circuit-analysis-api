package main

import (
	"fmt"
	"os"

	"github.com/luuisavelino/short-circuit-analysis-api/models"
	"github.com/luuisavelino/short-circuit-analysis-api/routes"
)

func main() {

	readedFiles, _ := os.ReadDir("./files/")
	for _, readedFile := range readedFiles {
		models.Files = append(models.Files, models.File{Nome: readedFile.Name()})
	}

	fmt.Println("Iniciando o servidor Rest com GO")
	routes.HandleRequest()
}
