package models

type File struct {
	Posicao int    `json:"posicao"`
	Nome    string `json:"nome"`
}

var Files []File
