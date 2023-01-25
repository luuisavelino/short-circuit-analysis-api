package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/short-circuit-analysis-elements/models"
	"github.com/luuisavelino/short-circuit-analysis-elements/pkg/elements"
	"github.com/xuri/excelize/v2"
)

const (
	mensagemErroIdArquivo = "Id do arquivo inválido"
	mensagemErroArquivo   = "Arquivo não encontrado"
	path                  = "./files/"
)

func AllFiles(c *gin.Context) {
	c.JSON(http.StatusOK, models.Files)
}

func OneFile(c *gin.Context) {
	file, err := Files(c)

	if err != nil {
		jsonError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"file": file,
	})
}

func AllTypes(c *gin.Context) {
	c.JSON(http.StatusOK, models.ElementTypes)
}

func OneType(c *gin.Context) {
	typeId, err := TypeId(c)
	if err != nil {
		jsonError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.ElementTypes[typeId])
}

func Elements(c *gin.Context) (map[string]map[string]models.Element, error) {
	file, err := Files(c)
	if err != nil {
		return nil, err
	}

	tabelaDados, err := excelize.OpenFile(path + file.Nome)
	if err != nil {
		return nil, err
	}

	models.Elements["1"], err = elements.ElementosTipo1(tabelaDados)
	if err != nil {
		return nil, err
	}

	models.Elements["2"], err = elements.ElementosTipo23(tabelaDados)
	if err != nil {
		return nil, err
	}

	return models.Elements, nil
}

func AllElementsType(c *gin.Context) {
	elements, err := Elements(c)
	if err != nil {
		jsonError(c, err)
		return
	}

	typeId, err := TypeId(c)
	if err != nil {
		jsonError(c, err)
		return
	}

	if typeId == "0" {
		c.JSON(http.StatusOK, gin.H{
			"1": elements["1"],
			"2": elements["2"],
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		typeId: elements[typeId],
	})
}

func OneElement(c *gin.Context) {
	element, err := Element(c)
	if err != nil {
		jsonError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		element.De + "-" + element.Para: element,
	})
}
