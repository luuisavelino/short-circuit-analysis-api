package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luuisavelino/short-circuit-analysis-elements/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/files", controllers.AllFiles).Methods("Get")
	r.HandleFunc("/files/{fileId}", controllers.OneFile).Methods("Get")
	r.HandleFunc("/files/{fileId}/elements", controllers.AllElements).Methods("Get")
	r.HandleFunc("/files/{fileId}/elements/type/{typeId}", controllers.AllElementsType).Methods("Get")
	r.HandleFunc("/files/{fileId}/elements/type/{typeId}/element/{elementId}", controllers.OneElement).Methods("Get")
	log.Fatal(http.ListenAndServe(":8001", r))
}
