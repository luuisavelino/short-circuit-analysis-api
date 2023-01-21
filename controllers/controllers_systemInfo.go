package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/short-circuit-analysis-elements/models"
	"github.com/luuisavelino/short-circuit-analysis-elements/pkg/functions"
	"github.com/xuri/excelize/v2"
)

func SystemSize(c *gin.Context) {
	file, err := Files(c)
	if err != nil {
		return
	}

	tabelaDados, err := excelize.OpenFile(path + file.Nome)
	if err != nil {
		jsonError(c, err)
		return
	}

	err = functions.SystemInfo(tabelaDados)
	if err != nil {
		jsonError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"size": models.System.Size,
	})
}

func SystemBars(c *gin.Context) {
	file, err := Files(c)
	if err != nil {
		jsonError(c, err)
		return
	}

	tabelaDados, err := excelize.OpenFile(path + file.Nome)
	if err != nil {
		jsonError(c, err)
		return
	}

	err = functions.SystemInfo(tabelaDados)
	if err != nil {
		jsonError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"bars": models.System.Bars,
	})
}

func Files(c *gin.Context) (models.File, error) {
	fileId, err := strconv.Atoi(c.Params.ByName("fileId"))
	files := new(models.File)

	if err != nil {
		return *files, err
	}

	for position, file := range models.Files {
		if position == fileId {
			return file, nil
		}
	}

	err = errors.New("id do arquivo invalido")

	return *files, err
}

func TypeId(c *gin.Context) (string, error) {
	typeId := c.Params.ByName("typeId")

	_, exit := models.ElementTypes[typeId]
	if !exit {
		return "", errors.New("id do arquivo invalido")
	}

	return typeId, nil
}

func jsonError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"Error": err.Error(),
	})
}
