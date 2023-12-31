package pack

import (
	"encoding/json"
	"fmt"
	"net/http"

	"re-parteners-tech-task/redis"

	"github.com/gorilla/mux"
)

type Handler struct {
	r *Repository
}

func NewHandler(r *Repository) *Handler {
	return &Handler{r: r}
}

// DeleteByID - returns the pack by the provided ID
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	packID := vars["pack_id"]
	packKey, err := h.r.GetByKey(packID)
	if redis.ErrorResponse(w, r, err, 400) {
		return
	}

	redis.JsonResponse(
		w,
		redis.CreateResponse(
			packKey,
			fmt.Sprintf("retrieved the Pack object with ID [%s]", packID),
			nil,
			200,
		))
}

// DeleteByID - sets a new pack with the input value
func (h *Handler) Set(w http.ResponseWriter, r *http.Request) {
	var (
		p   Pack
		err = json.NewDecoder(r.Body).Decode(&p)
	)

	if err != nil {
		redis.JsonResponse(
			w,
			redis.CreateResponse(
				p,
				"please enter number with numeric values / digits [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]",
				err,
				405,
			))
		return
	}

	if p.Size <= 0 {
		redis.JsonResponse(
			w,
			redis.CreateResponse(
				nil,
				"please enter valid number [>0]",
				fmt.Errorf("please enter valid number"),
				405,
			))
		return
	}

	if exists(p.Size) {
		redis.JsonResponse(
			w,
			redis.CreateResponse(
				nil,
				"pack size already exists",
				fmt.Errorf("pack size already exists"),
				400,
			))
		return
	}

	h.r.Set(p)
	redis.JsonResponse(
		w,
		redis.CreateResponse(
			p,
			fmt.Sprintf("set the Pack object [%v]", p),
			nil,
			200,
		))
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	res := h.r.GetAll()

	redis.JsonResponse(
		w,
		redis.CreateResponse(
			res,
			"successfully retrieved all pack objects",
			nil,
			200,
		))
}

// DeleteByID - deletes the pack by the provided ID
func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	packID := vars["pack_id"]
	err := h.r.DeleteByID(packID)
	if redis.ErrorResponse(w, r, err, 400) {
		return
	}

	redis.JsonResponse(
		w,
		redis.CreateResponse(
			nil,
			fmt.Sprintf("successfully deleted pack with ID [%s]", packID),
			nil,
			200,
		))
}
