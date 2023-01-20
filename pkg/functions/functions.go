package functions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func Impedancia(resistenciaLinha string, reatanciaLinha string, impedanciaAtualStr string) (complex128, error) {
	resistencia, err := strconv.ParseFloat(replaceCommaWithDot(resistenciaLinha), 64)
	if err != nil {
		return 0, err
	}

	reatancia, err := strconv.ParseFloat(replaceCommaWithDot(reatanciaLinha), 64)
	if err != nil {
		return 0, err
	}

	impedancia := complex(resistencia, reatancia)

	if impedanciaAtualStr != "" && impedanciaAtualStr != "0" {
		impedanciaAtual, err := strconv.ParseComplex(replaceCommaWithDot(impedanciaAtualStr), 128)
		if err != nil {
			return 0, err
		}

		impedancia = (impedancia * impedanciaAtual) / (impedancia + impedanciaAtual)
	}

	return impedancia, nil
}

func StringToFloat(grandezaStr string) (float64, error) {
	grandeza, err := strconv.ParseFloat(replaceCommaWithDot(grandezaStr), 64)
	if err != nil {
		return 0, err
	}

	return grandeza, nil
}

func SystemInfo(tabelaExcel *excelize.File) (int, []string, error) {
	barras, err := tabelaExcel.GetRows("Barra")
	if err != nil {
		fmt.Println(err.Error())
		return 0, nil, err
	}

	systemSize := len(barras) - 2

	var bars []string

	for x := 2; x < len(barras); x++ {
		bars = append(bars, (barras[x][0]))
	}

	return systemSize, bars, nil
}

func replaceCommaWithDot(s string) string {
	if !strings.Contains(s, ",") {
		return s
	}
	return strings.Replace(s, ",", ".", -1)
}
