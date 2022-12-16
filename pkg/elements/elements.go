package elements

import (
	"strconv"

	"github.com/luuisavelino/short-circuit-analysis-elements/models"
	"github.com/luuisavelino/short-circuit-analysis-elements/pkg/functions"
	"github.com/xuri/excelize/v2"
)

func transformadores(tabela_excel *excelize.File) map[string]models.Element {

	dados_transformadores, err := tabela_excel.GetRows(tabela_excel.GetSheetList()[3])
	dados_transformadores = dados_transformadores[2:]
	functions.ErrorValidade(err)

	elementos_transformadores := make(map[string]models.Element)
	for x := 0; x < len(dados_transformadores); x++ {

		transformador := dados_transformadores[x][0] + "-" + dados_transformadores[x][1]
		impedancia_atual_p := elementos_transformadores[transformador].Z_positiva
		impedancia_atual_z := elementos_transformadores[transformador].Z_zero

		z_positiva := functions.Impedancia(dados_transformadores[x][3], dados_transformadores[x][4], impedancia_atual_p)
		z_zero := functions.Impedancia("0", dados_transformadores[x][5], impedancia_atual_z) + 3*functions.Impedancia("0", dados_transformadores[x][6], impedancia_atual_z)

		elementos_transformadores[transformador] = models.Element{
			Id:         x,
			De:         dados_transformadores[x][0],
			Para:       dados_transformadores[x][1],
			Nome:       dados_transformadores[x][2],
			Z_positiva: strconv.FormatComplex(z_positiva, 'g', 'g', 64),
			Z_zero:     strconv.FormatComplex(z_zero, 'g', 'g', 64),
		}
	}

	return elementos_transformadores
}

func Elementos_tipo_1(tabela_excel *excelize.File) map[string]models.Element {
	var elementos_tipo_1 = make(map[string]models.Element)

	dados_linhas, err := tabela_excel.GetRows(tabela_excel.GetSheetList()[2])
	dados_linhas = dados_linhas[1:]
	functions.ErrorValidade(err)

	for x := 0; x < len(dados_linhas); x++ {
		elementos_tipo_1[dados_linhas[x][0]] = models.Element{
			Id:         x,
			De:         dados_linhas[x][0],
			Nome:       dados_linhas[x][1],
			Z_positiva: strconv.FormatComplex(complex(0, functions.StringToFloat(dados_linhas[x][2])/100), 'g', 'g', 64),
			Z_zero:     strconv.FormatComplex(complex(0, functions.StringToFloat(dados_linhas[x][3])+3*functions.StringToFloat(dados_linhas[x][4])), 'g', 'g', 128),
		}
	}

	return elementos_tipo_1
}

func Elementos_tipo_2_3(tabela_excel *excelize.File) map[string]models.Element {
	var elementos_tipo_2_3 = make(map[string]models.Element)

	dados_linhas, err := tabela_excel.GetRows(tabela_excel.GetSheetList()[1])
	dados_linhas = dados_linhas[2:]
	functions.ErrorValidade(err)

	// Adicionando todos os elementos dos transformadores como tipos 2 e 3
	for _, dado_do_transformador := range transformadores(tabela_excel) {
		elementos_tipo_2_3[dado_do_transformador.De+"-"+dado_do_transformador.Para] = dado_do_transformador
	}
	elementId := len(elementos_tipo_2_3)

	// Adicionando as linhas como elementos do tipo 2 e 3
	for x := 0; x < len(dados_linhas); x++ {
		z_positiva := functions.Impedancia(dados_linhas[x][2], dados_linhas[x][3], "0")
		z_zero := functions.Impedancia(dados_linhas[x][4], dados_linhas[x][5], "0")

		elementos_tipo_2_3[dados_linhas[x][0]+"-"+dados_linhas[x][1]] = models.Element{
			Id:         elementId,
			De:         dados_linhas[x][0],
			Para:       dados_linhas[x][1],
			Nome:       dados_linhas[x][2],
			Z_positiva: strconv.FormatComplex(z_positiva, 'g', 'g', 64),
			Z_zero:     strconv.FormatComplex(z_zero, 'g', 'g', 64),
		}
		elementId++
	}

	return elementos_tipo_2_3
}
