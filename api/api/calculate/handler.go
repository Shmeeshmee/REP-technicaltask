package calculate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	r *Repository
}

func NewHandler(r *Repository) *Handler {
	return &Handler{r: r}
}

func (h *Handler) Calculate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// mux.
	amount := vars["amount"]
	iAmount, err := strconv.Atoi(amount)
	if err != nil {
		return
	}

	vals, err := h.r.Calculate(iAmount)
	if err != nil {
		return
	}

	jsonResponse, err := json.Marshal(vals)
	if err != nil {
		fmt.Println("cannot marhal the response, ", err)
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		fmt.Println("cannot write the json response, ", err)
	}
}
