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
