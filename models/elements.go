package models

type Element struct {
	Id         int    `json:"id"`
	De         string `json:"de"`
	Para       string `json:"para"`
	Nome       string `json:"nome"`
	Z_positiva string `json:"z_positiva"`
	Z_zero     string `json:"z_zero"`
}

var Elements = make(map[string]map[string]Element)

type ElementType struct {
	Info string `json:"info"`
}

var ElementTypes = make(map[string]ElementType)