package controllers

import (
	"net/http"
	"strconv"

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
	fileId, err := strconv.Atoi(c.Params.ByName("fileId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Not Found": mensagemErroIdArquivo,
		})
		return
	}

	for position, file := range models.Files {
		if position == fileId {
			c.JSON(http.StatusOK, gin.H{
				"file": file,
			})

		}
	}
}

func AllTypes(c *gin.Context) {
	c.JSON(http.StatusOK, models.ElementTypes)
}

func OneType(c *gin.Context) {
	typeId := c.Params.ByName("typeId")
	c.JSON(http.StatusOK, models.ElementTypes[typeId])
}

func AllElementsType(c *gin.Context) {
	fileId, err := strconv.Atoi(c.Params.ByName("fileId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Not Found": mensagemErroIdArquivo,
		})
	}

	typeId := c.Params.ByName("typeId")

	tabelaDados, err := excelize.OpenFile(path + models.Files[fileId].Nome)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Not Found": mensagemErroArquivo,
		})
	}

	models.Elements["1"], err = elements.Elementos_tipo_1(tabelaDados)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": "Erro nos elementos do tipo 1",
		})
	}

	models.Elements["2"], err = elements.Elementos_tipo_2_3(tabelaDados)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": "Erro nos elementos do tipo 2 ou 3",
		})
		return
	}

	if typeId == "0" {
		c.JSON(http.StatusOK, gin.H{
			"1": models.Elements["1"],
			"2": models.Elements["2"],
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		typeId: models.Elements[typeId],
	})
}

func OneElement(c *gin.Context) {
	fileId, err := strconv.Atoi(c.Params.ByName("fileId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Not Found": mensagemErroIdArquivo,
		})
	}

	typeId := c.Params.ByName("typeId")

	elementId, err := strconv.Atoi(c.Params.ByName("elementId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Not Found": "Elemento não encontrado",
		})
	}

	tabelaDados, err := excelize.OpenFile(path + models.Files[fileId].Nome)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Not Found": mensagemErroArquivo,
		})
		return
	}

	models.Elements["1"], err = elements.Elementos_tipo_1(tabelaDados)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": "Erro nos elementos do tipo 1",
		})
		return
	}

	models.Elements["2"], err = elements.Elementos_tipo_2_3(tabelaDados)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": "Erro nos elementos do tipo 2 ou 3",
		})
		return
	}

	for _, element := range models.Elements[typeId] {
		if element.Id == elementId {
			c.JSON(http.StatusOK, gin.H{
				"element": element,
			})
		}
	}
}
