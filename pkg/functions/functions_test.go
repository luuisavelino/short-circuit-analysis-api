package functions

import (
	"fmt"
	"testing"

	"github.com/luuisavelino/short-circuit-analysis-elements/models"
	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
)



func TestImpedancia(t *testing.T) {
	// Testa o caso de sucesso com valores válidos
	result, err := Impedancia("3", "4", "0")
	assert.Nil(t, err)
	assert.Equal(t, complex(3, 4), result)

	result, err = Impedancia("7.13", "2.105", "1.13+17.2i")
	assert.Nil(t, err)
	assert.Equal(t, complex(4.946345423746592, 3.5745098782774845), result)

	// Testa o caso de erro com resistenciaLinha inválido
	result, err = Impedancia("a", "4", "5+6i")
	assert.NotNil(t, err)
	assert.Equal(t, 0+0i, result)


	// Testa o caso de erro com reatanciaLinha inválido
	result, err = Impedancia("3", "b", "5+6i")
	assert.NotNil(t, err)
	assert.Equal(t, 0+0i, result)

	// Testa o caso de erro com impedanciaAtualStr inválido
	result, err = Impedancia("3", "4", "c")
	assert.NotNil(t, err)
	assert.Equal(t, 0+0i, result)
}

func TestSystemInfo(t *testing.T) {
	var f = excelize.NewFile()
	f.NewSheet("dados_de_barra")
	f.SetCellValue("dados_de_barra", "A1","")
	f.SetCellValue("dados_de_barra", "A2","")
	f.SetCellValue("dados_de_barra", "A3","1")
	f.SetCellValue("dados_de_barra", "A4","2")
	f.SetCellValue("dados_de_barra", "A5","3")
	f.SetCellValue("dados_de_barra", "A6","4")
	f.SetCellValue("dados_de_barra", "A7","5")

	err := SystemInfo(f)
	fmt.Println(err)

	assert.Nil(t, err)
	assert.Equal(t, 5, models.System.Size, "O tamanho do sistema deveria ser 5")
	assert.Equal(t, []string{"1", "2", "3", "4", "5"}, models.System.Bars, "As barras do sistema deveria ser '1', '2', '3', '4', '5'")
}
