package functions

import (
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func Impedancia(resistenciaLinha string, reatanciaLinha string, impedanciaAtualStr string) complex128 {
	var resistencia, _ = strconv.ParseFloat(resistenciaLinha, 64)
	var reatancia, _ = strconv.ParseFloat(reatanciaLinha, 64)
	var impedanciaAtual, _ = strconv.ParseComplex(impedanciaAtualStr, 128)

	impedancia := complex(resistencia, reatancia)

	if impedanciaAtual != 0 {
		impedancia = (impedancia * impedanciaAtual) / (impedancia + impedanciaAtual)
	}

	return impedancia
}

func StringToFloat(grandezaStr string) float64 {
	grandeza, _ := strconv.ParseFloat(grandezaStr, 64)

	return grandeza
}

func ErrorValidade(err error) {
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func SystemInfo(tabelaExcel *excelize.File) (int, []string) {
	barras, _ := tabelaExcel.GetRows(tabelaExcel.GetSheetList()[0])
	systemSize := len(barras) - 2

	var bars []string

	for x := 2; x < len(barras); x++ {
		bars = append(bars, (barras[x][0]))
	}

	return systemSize, bars
}
