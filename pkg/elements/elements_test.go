package elements

import (
	"testing"

	"github.com/luuisavelino/short-circuit-analysis-elements/models"
	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
)

func preenchePlanilha() *excelize.File {
	var f = excelize.NewFile()
	f.NewSheet("dados_de_transformadores")
	f.SetCellValue("dados_de_transformadores", "A3", "895")
	f.SetCellValue("dados_de_transformadores", "B3", "814")
	f.SetCellValue("dados_de_transformadores", "C3", "Bateias")
	f.SetCellValue("dados_de_transformadores", "D3", "0,032")
	f.SetCellValue("dados_de_transformadores", "E3", "1,146")
	f.SetCellValue("dados_de_transformadores", "F3", "1,26")
	f.SetCellValue("dados_de_transformadores", "G3", "4,1")

	f.NewSheet("dados_de_geradores")
	f.SetCellValue("dados_de_geradores", "A2", "1")
	f.SetCellValue("dados_de_geradores", "B2", "first")
	f.SetCellValue("dados_de_geradores", "C2", "4,5")
	f.SetCellValue("dados_de_geradores", "D2", "3,4")
	f.SetCellValue("dados_de_geradores", "E2", "10")

	f.NewSheet("dados_de_linha")
	f.SetCellValue("dados_de_linha", "A3", "824")
	f.SetCellValue("dados_de_linha", "B3", "933")
	f.SetCellValue("dados_de_linha", "C3", "G.B.Munhoz-Areia")
	f.SetCellValue("dados_de_linha", "D3", "0,01")
	f.SetCellValue("dados_de_linha", "E3", "0,124")
	f.SetCellValue("dados_de_linha", "F3", "0,04")
	f.SetCellValue("dados_de_linha", "G3", "0,29")

	return f
}

var f = preenchePlanilha()

func TestTransformadores(t *testing.T) {
	expected := map[string]models.Element{"895-814": {Id: 0, De: "895", Para: "814", Nome: "Bateias", Z_positiva: "(0.032+1.146i)", Z_zero: "(0+13.559999999999999i)"}}
	result, err := Transformadores(f)

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestElementosTipo1(t *testing.T) {
	expected := map[string]models.Element{"1": {Id: 0, De: "1", Nome: "first", Z_positiva: "(0+0.045i)", Z_zero: "(0+33.4i)"}}
	result, err := ElementosTipo1(f)

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestElementosTipo23(t *testing.T) {
	expected := map[string]models.Element{
		"895-814": {
			Id: 0, De: "895", Para: "814", Nome: "Bateias", Z_positiva: "(0.032+1.146i)", Z_zero: "(0+13.559999999999999i)",
		},
		"824-933": {
			Id: 1, De: "824", Para: "933", Nome: "G.B.Munhoz-Areia", Z_positiva: "(0.01+0.124i)", Z_zero: "(0.04+0.29i)",
		},
	}

	result, err := ElementosTipo23(f)

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}
