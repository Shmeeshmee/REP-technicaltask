package pack

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	r *Repository
}

func NewHandler(r *Repository) *Handler {
	return &Handler{r: r}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	packID := vars["pack_id"]
	packKey, err := h.r.GetByKey(packID)
	if err != nil {
		return
	}

	jsonResponse, err := json.Marshal(packKey)
	if err != nil {
		fmt.Println("cannot marhal the response, ", err)
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		fmt.Println("cannot write the json response, ", err)
	}
}

func (h *Handler) Set(w http.ResponseWriter, r *http.Request) {
	var (
		p Pack
		_ = json.NewDecoder(r.Body).Decode(&p)
	)

	if exists(p.Size) {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("pack size already exists"))
		if err != nil {
			fmt.Println("cannot write the json response, ", err)
		}
		return
	}

	h.r.Set(p)
	fmt.Println("entered Set pack")

	_, err := w.Write([]byte(fmt.Sprintf("value of [%+v] has been set", p)))
	if err != nil {
		fmt.Println("cannot write the json response, ", err)
	}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	res := h.r.GetAll()

	jsonResponse, err := json.Marshal(res)
	if err != nil {
		fmt.Println("cannot marhal the response, ", err)
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		fmt.Println("cannot write the json response, ", err)
	}
}

func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	packID := vars["pack_id"]
	err := h.r.DeleteByID(packID)
	if err != nil {
		return
	}

	_, err = w.Write([]byte(fmt.Sprintf("value of [%s] has been deleted", packID)))
	if err != nil {
		fmt.Println("cannot write the json response, ", err)
	}
}
