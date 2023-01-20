package elements

import (
	"fmt"
	"testing"

	"github.com/luuisavelino/short-circuit-analysis-elements/models"
	"github.com/xuri/excelize/v2"
)

func preenchePlanilha() *excelize.File {
	var f = excelize.NewFile()
	f.NewSheet("dados_de_transformadores")
	f.SetCellValue("Transformadores", "A3", "895")
	f.SetCellValue("Transformadores", "B3", "814")
	f.SetCellValue("Transformadores", "C3", "Bateias")
	f.SetCellValue("Transformadores", "D3", "0,032")
	f.SetCellValue("Transformadores", "E3", "1,146")
	f.SetCellValue("Transformadores", "F3", "1,26")
	f.SetCellValue("Transformadores", "G3", "4,1")

	f.NewSheet("dados_de_geradores")
	f.SetCellValue("Geradores", "A2", "1")
	f.SetCellValue("Geradores", "B2", "first")
	f.SetCellValue("Geradores", "C2", "4,5")
	f.SetCellValue("Geradores", "D2", "3,4")
	f.SetCellValue("Geradores", "E2", "10")

	f.NewSheet("dados_de_linha")
	f.SetCellValue("Linha", "A3", "824")
	f.SetCellValue("Linha", "B3", "933")
	f.SetCellValue("Linha", "C3", "G.B.Munhoz-Areia")
	f.SetCellValue("Linha", "D3", "0,01")
	f.SetCellValue("Linha", "E3", "0,124")
	f.SetCellValue("Linha", "F3", "0,04")
	f.SetCellValue("Linha", "G3", "0,29")

	return f
}

var f = preenchePlanilha()

func TestTransformadores(t *testing.T) {
	expected := map[string]models.Element{"895-814": {Id: 0, De: "895", Para: "814", Nome: "Bateias", Z_positiva: "(0.032+1.146i)", Z_zero: "(0+13.559999999999999i)"}}
	actual, err := Transformadores(f)
	if err != nil {
		t.Errorf("Expected no error, but got %s", err.Error())
	}

	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestElementosTipo1(t *testing.T) {
	expected := map[string]models.Element{"1": {Id: 0, De: "1", Nome: "first", Z_positiva: "(0+0.045i)", Z_zero: "(0+33.4i)"}}
	actual, err := ElementosTipo1(f)

	if err != nil {
		t.Errorf("Expected no error, but got %s", err.Error())
	}

	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", actual) {
		t.Errorf("\nExpected \t%v\nbut got \t%v", expected, actual)
	}
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

	actual, err := ElementosTipo23(f)

	if err != nil {
		t.Errorf("Expected no error, but got %s", err.Error())
	}

	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", actual) {
		t.Errorf("\nExpected \t%v\nbut got \t%v", expected, actual)
	}
}
