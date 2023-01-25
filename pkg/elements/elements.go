package elements

import (
	"fmt"

	"github.com/luuisavelino/short-circuit-analysis-elements/models"
	"github.com/luuisavelino/short-circuit-analysis-elements/pkg/functions"
	"github.com/xuri/excelize/v2"
)

func Transformadores(tabela_excel *excelize.File) (map[string]models.Element, error) {

	dadosTransformadores, err := tabela_excel.GetRows("dados_de_transformadores")
	dadosTransformadores = dadosTransformadores[2:]
	if err != nil {
		return nil, err
	}

	elementos_transformadores := make(map[string]models.Element)
	for x := 0; x < len(dadosTransformadores); x++ {

		transformador := dadosTransformadores[x][0] + "-" + dadosTransformadores[x][1]
		impedanciaAtualP := elementos_transformadores[transformador].Z_positiva
		impedanciaAtualZ := elementos_transformadores[transformador].Z_zero

		z_positiva, err := functions.Impedancia(dadosTransformadores[x][3], dadosTransformadores[x][4], impedanciaAtualP)
		if err != nil {

			return nil, err
		}

		z_zero, err := functions.Impedancia("0", dadosTransformadores[x][5], impedanciaAtualZ)
		if err != nil {

			return nil, err
		}

		zn, err := functions.Impedancia("0", dadosTransformadores[x][6], impedanciaAtualZ)
		if err != nil {

			return nil, err
		}

		elementos_transformadores[transformador] = models.Element{
			De:         dadosTransformadores[x][0],
			Para:       dadosTransformadores[x][1],
			Nome:       dadosTransformadores[x][2],
			Z_positiva: fmt.Sprint(z_positiva),
			Z_zero:     fmt.Sprint(z_zero + 3*zn),
		}
	}

	return elementos_transformadores, nil
}

func ElementosTipo1(tabela_excel *excelize.File) (map[string]models.Element, error) {
	var elementosTipo1 = make(map[string]models.Element)

	dadosLinhas, err := tabela_excel.GetRows("dados_de_geradores")
	dadosLinhas = dadosLinhas[1:]

	if err != nil {
		return nil, err
	}

	for x := 0; x < len(dadosLinhas); x++ {

		z_positiva, err := functions.StringToFloat(dadosLinhas[x][2])
		if err != nil {
			return nil, err
		}

		z_zero, err := functions.StringToFloat(dadosLinhas[x][3])
		if err != nil {
			return nil, err
		}

		zn, err := functions.StringToFloat(dadosLinhas[x][4])
		if err != nil {
			return nil, err
		}

		elementosTipo1[dadosLinhas[x][0]] = models.Element{
			De:         dadosLinhas[x][0],
			Nome:       dadosLinhas[x][1],
			Z_positiva: fmt.Sprint("(0+", z_positiva/100, "i)"),
			Z_zero:     fmt.Sprint("(0+", z_zero+3*zn, "i)"),
		}
	}

	return elementosTipo1, nil
}

func ElementosTipo23(tabela_excel *excelize.File) (map[string]models.Element, error) {
	var elementosTipo23 = make(map[string]models.Element)

	dadosLinhas, err := tabela_excel.GetRows("dados_de_linha")
	dadosLinhas = dadosLinhas[2:]
	if err != nil {
		return nil, err
	}

	transformadores, err := Transformadores(tabela_excel)
	if err != nil {
		return nil, err
	}

	for _, dado_do_transformador := range transformadores {
		elementosTipo23[dado_do_transformador.De+"-"+dado_do_transformador.Para] = dado_do_transformador
	}
	elementId := len(elementosTipo23)

	for x := 0; x < len(dadosLinhas); x++ {
		z_positiva, err := functions.Impedancia(dadosLinhas[x][3], dadosLinhas[x][4], "0")
		if err != nil {
			return nil, err
		}

		z_zero, err := functions.Impedancia(dadosLinhas[x][5], dadosLinhas[x][6], "0")
		if err != nil {
			return nil, err
		}

		elementosTipo23[dadosLinhas[x][0]+"-"+dadosLinhas[x][1]] = models.Element{
			De:         dadosLinhas[x][0],
			Para:       dadosLinhas[x][1],
			Nome:       dadosLinhas[x][2],
			Z_positiva: fmt.Sprint(z_positiva),
			Z_zero:     fmt.Sprint(z_zero),
		}
		elementId++
	}

	return elementosTipo23, nil
}
