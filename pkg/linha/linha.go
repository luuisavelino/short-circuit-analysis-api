package barra

import (
	"encoding/json"
	"fmt"

	"github.com/luuisavelino/short-circuit-analysis-api/models"
)

type Barra struct {
	Key   string `json:"key"`
	Color string `json:"color"`
}

type Linha struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func Linhas_sistema(barras []string, elementos_2_3 map[string]models.Element) {
	var dados_linhas []Linha
	var dados_barras []Barra

	for _, barra := range barras {
		dados_barras = append(dados_barras, Barra{Key: barra, Color: "grey"})
	}

	for _, linha := range elementos_2_3 {
		dados_linhas = append(dados_linhas, Linha{From: linha.De, To: linha.Para})
	}

	a, _ := json.Marshal(dados_linhas)
	b, _ := json.Marshal(dados_barras)

	fmt.Println(b)
	fmt.Println(a)
}
