package functions

import (
	"strconv"
	"strings"

	"github.com/luuisavelino/short-circuit-analysis-elements/models"
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

func SystemInfo(tabelaExcel *excelize.File) error {
	data, err := tabelaExcel.GetRows("dados_de_barra")
	var barras models.Info
	if err != nil {
		return err
	}

	models.System.Size = len(data) - 2

	for x := 2; x < len(data); x++ {
		barras.Bars = append(barras.Bars, (data[x][0]))
	}

	models.System.Bars = barras.Bars
	return nil
}

func replaceCommaWithDot(s string) string {
	if !strings.Contains(s, ",") {
		return s
	}
	return strings.Replace(s, ",", ".", -1)
}
