package router

import (
	"fmt"
	"net/http"
	"re-parteners-tech-task/api/calculate"
	"re-parteners-tech-task/api/pack"
	"re-parteners-tech-task/config"

	"github.com/gorilla/mux"
)

type handlers struct {
	ph *pack.Handler
	ch *calculate.Handler
}

var h = handlers{}

func routes(cfg *config.Config, r *mux.Router) {
	r.HandleFunc(
		fmt.Sprintf("/%s/{pack_id}", "pack"),
		addDefaultHeaders(h.ph.Get),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s", "packs"),
		addDefaultHeaders(h.ph.GetAll),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s", "pack"),
		addDefaultHeaders(h.ph.Set),
	).Methods("POST")
	r.HandleFunc(
		fmt.Sprintf("/%s/{pack_id}", "pack"),
		addDefaultHeaders(h.ph.DeleteByID),
	).Methods("DELETE")

	r.HandleFunc(
		fmt.Sprintf("/%s/{amount}", "calculate"),
		addDefaultHeaders(h.ch.Calculate),
	).Methods("GET")
}

func init() {
	pr := pack.NewRepository()
	cr := calculate.NewRepository()

	h = handlers{
		ph: pack.NewHandler(pr),
		ch: calculate.NewHandler(cr),
	}
}

func addDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// EnableCors(&w)
		// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		// w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, GET, OPTIONS, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Content-Length, Accept-Encoding, X-CSRF-Token")
		fn(w, r)
	}
}

// func EnableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// 	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
// }
