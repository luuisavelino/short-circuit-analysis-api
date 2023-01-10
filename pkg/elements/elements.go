package elements

import (
	"fmt"
	"strconv"

	"github.com/luuisavelino/short-circuit-analysis-elements/models"
	"github.com/luuisavelino/short-circuit-analysis-elements/pkg/functions"
	"github.com/xuri/excelize/v2"
)

func Transformadores(tabela_excel *excelize.File) (map[string]models.Element, error) {

	dadosTransformadores, err := tabela_excel.GetRows(tabela_excel.GetSheetList()[3])
	dadosTransformadores = dadosTransformadores[2:]
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	elementos_transformadores := make(map[string]models.Element)
	for x := 0; x < len(dadosTransformadores); x++ {

		transformador := dadosTransformadores[x][0] + "-" + dadosTransformadores[x][1]
		impedanciaAtualP := elementos_transformadores[transformador].Z_positiva
		impedanciaAtualZ := elementos_transformadores[transformador].Z_zero

		z_positiva := functions.Impedancia(dadosTransformadores[x][3], dadosTransformadores[x][4], impedanciaAtualP)
		z_zero := functions.Impedancia("0", dadosTransformadores[x][5], impedanciaAtualZ) + 3*functions.Impedancia("0", dadosTransformadores[x][6], impedanciaAtualZ)

		elementos_transformadores[transformador] = models.Element{
			Id:         x,
			De:         dadosTransformadores[x][0],
			Para:       dadosTransformadores[x][1],
			Nome:       dadosTransformadores[x][2],
			Z_positiva: strconv.FormatComplex(z_positiva, 'g', 'g', 64),
			Z_zero:     strconv.FormatComplex(z_zero, 'g', 'g', 64),
		}
	}

	return elementos_transformadores, nil
}

func Elementos_tipo_1(tabela_excel *excelize.File) (map[string]models.Element, error) {
	var elementosTipo1 = make(map[string]models.Element)

	dadosLinhas, err := tabela_excel.GetRows(tabela_excel.GetSheetList()[2])
	dadosLinhas = dadosLinhas[1:]
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for x := 0; x < len(dadosLinhas); x++ {
		elementosTipo1[dadosLinhas[x][0]] = models.Element{
			Id:         x,
			De:         dadosLinhas[x][0],
			Nome:       dadosLinhas[x][1],
			Z_positiva: strconv.FormatComplex(complex(0, functions.StringToFloat(dadosLinhas[x][2])/100), 'g', 'g', 64),
			Z_zero:     strconv.FormatComplex(complex(0, functions.StringToFloat(dadosLinhas[x][3])+3*functions.StringToFloat(dadosLinhas[x][4])), 'g', 'g', 128),
		}
	}

	return elementosTipo1, nil
}

func Elementos_tipo_2_3(tabela_excel *excelize.File) (map[string]models.Element, error) {
	var elementosTipo23 = make(map[string]models.Element)

	dadosLinhas, err := tabela_excel.GetRows(tabela_excel.GetSheetList()[1])
	dadosLinhas = dadosLinhas[2:]
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	transformadores, _ := Transformadores(tabela_excel)

	// Adicionando todos os elementos dos transformadores como tipos 2 e 3
	for _, dado_do_transformador := range transformadores {
		elementosTipo23[dado_do_transformador.De+"-"+dado_do_transformador.Para] = dado_do_transformador
	}
	elementId := len(elementosTipo23)

	// Adicionando as linhas como elementos do tipo 2 e 3
	for x := 0; x < len(dadosLinhas); x++ {
		z_positiva := functions.Impedancia(dadosLinhas[x][2], dadosLinhas[x][3], "0")
		z_zero := functions.Impedancia(dadosLinhas[x][4], dadosLinhas[x][5], "0")

		elementosTipo23[dadosLinhas[x][0]+"-"+dadosLinhas[x][1]] = models.Element{
			Id:         elementId,
			De:         dadosLinhas[x][0],
			Para:       dadosLinhas[x][1],
			Nome:       dadosLinhas[x][2],
			Z_positiva: strconv.FormatComplex(z_positiva, 'g', 'g', 64),
			Z_zero:     strconv.FormatComplex(z_zero, 'g', 'g', 64),
		}
		elementId++
	}

	return elementosTipo23, nil
}
