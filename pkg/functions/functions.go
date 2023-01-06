package functions

import (
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func Impedancia(resistencia_linha string, reatancia_linha string, impedanciaAtual_str string) complex128 {
	var resistencia, _ = strconv.ParseFloat(resistencia_linha, 64)
	var reatancia, _ = strconv.ParseFloat(reatancia_linha, 64)
	var impedanciaAtual, _ = strconv.ParseComplex(impedanciaAtual_str, 128)

	impedancia := complex(resistencia, reatancia)

	if impedanciaAtual != 0 {
		impedancia = (impedancia * impedanciaAtual) / (impedancia + impedanciaAtual)
	}

	return impedancia
}

func StringToFloat(grandeza_str string) float64 {
	grandeza, _ := strconv.ParseFloat(grandeza_str, 64)

	return grandeza
}

func ErrorValidade(err error) {
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func SystemInfo(tabela_excel *excelize.File) (int, []string) {
	barras, _ := tabela_excel.GetRows(tabela_excel.GetSheetList()[0])
	systemSize := len(barras) - 2

	var bars []string

	for x := 2; x < len(barras); x++ {
		bars = append(bars, (barras[x][0]))
	}

	return systemSize, bars
}
