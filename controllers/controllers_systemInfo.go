package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/short-circuit-analysis-elements/models"
	"github.com/luuisavelino/short-circuit-analysis-elements/pkg/functions"
	"github.com/xuri/excelize/v2"
)


func SystemSize(c *gin.Context) {
	fileId, err := strconv.Atoi(c.Params.ByName("fileId"))
	if err != nil {

		return
	}

	tabelaDados, err := excelize.OpenFile(path + models.Files[fileId].Nome)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Not Found": mensagemErroArquivo,
		})
		return
	}

	systemSize, _, err := functions.SystemInfo(tabelaDados)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": "Erro na tabela dos geradores",
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"size": systemSize,
	})
}

func SystemBars(c *gin.Context) {
	fileId, err := strconv.Atoi(c.Params.ByName("fileId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Not Found": "Id do arquivo inválido",
		})
		return
	}

	tabelaDados, err := excelize.OpenFile(path + models.Files[fileId].Nome)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Not Found": mensagemErroArquivo,
		})
		return
	}

	_, bars, err := functions.SystemInfo(tabelaDados)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": "Erro na tabela dos geradores",
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"bars": bars,
	})
}