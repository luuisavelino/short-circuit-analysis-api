package elements

import "github.com/luuisavelino/short-circuit-analysis-elements/models"

func RemoveElemento(element map[string]models.Element, key string) {



	delete(element, key)
}
