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

func AllElements(c *gin.Context) {
	fileId, err := strconv.Atoi(c.Params.ByName("fileId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Not Found": mensagemErroIdArquivo,
		})
	}

	tabelaDados, err := excelize.OpenFile(path + models.Files[fileId].Nome)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Not Found": mensagemErroArquivo,
		})
	}

	models.Elements["1"] = elements.Elementos_tipo_1(tabelaDados)
	models.Elements["2"] = elements.Elementos_tipo_2_3(tabelaDados)

	c.JSON(http.StatusOK, models.Elements)
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

	models.Elements["1"] = elements.Elementos_tipo_1(tabelaDados)
	models.Elements["2"] = elements.Elementos_tipo_2_3(tabelaDados)

	c.JSON(http.StatusOK, gin.H{
		"element": models.Elements[typeId],
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
	models.Elements["1"] = elements.Elementos_tipo_1(tabelaDados)
	models.Elements["2"] = elements.Elementos_tipo_2_3(tabelaDados)

	for _, element := range models.Elements[typeId] {
		if element.Id == elementId {
			c.JSON(http.StatusOK, gin.H{
				"element": element,
			})
		}
	}
}
