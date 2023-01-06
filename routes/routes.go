package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luuisavelino/short-circuit-analysis-elements/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/health/liveness", controllers.Liveness).Methods("Get")
	r.HandleFunc("/health/readiness", controllers.Readiness).Methods("Get")
	r.HandleFunc("/api/files", controllers.AllFiles).Methods("Get")
	r.HandleFunc("/api/files/{fileId}", controllers.OneFile).Methods("Get")
	r.HandleFunc("/api/files/{fileId}/size", controllers.SystemSize).Methods("Get")
	r.HandleFunc("/api/files/{fileId}/bars", controllers.SystemBars).Methods("Get")
	r.HandleFunc("/api/files/{fileId}/elements", controllers.AllElements).Methods("Get")
	r.HandleFunc("/api/files/{fileId}/elements/type/{typeId}", controllers.AllElementsType).Methods("Get")
	r.HandleFunc("/api/files/{fileId}/elements/type/{typeId}/element/{elementId}", controllers.OneElement).Methods("Get")
	log.Fatal(http.ListenAndServe(":8080", r))
}
