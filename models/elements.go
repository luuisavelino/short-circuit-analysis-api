package models

type Element struct {
	De         string `json:"de"`
	Para       string `json:"para"`
	Nome       string `json:"nome"`
	Z_positiva string `json:"z_positiva"`
	Z_zero     string `json:"z_zero"`
}

var Elements = make(map[string]Element)
