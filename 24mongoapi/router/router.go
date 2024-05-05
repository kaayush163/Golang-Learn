package router

import (
	"github.com/gorilla/mux"
	controller "github.com/kaayush163/Golang-Learn/blob/main/24mongoapi/controllers"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")

	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/delteallmovie", controller.DeleteAllMovie).Methods("DELETE")

	return router
}
