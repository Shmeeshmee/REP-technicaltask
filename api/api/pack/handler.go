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

func (h *Handler) Set(w http.ResponseWriter, r *http.Request) {
	var (
		p Pack
		_ = json.NewDecoder(r.Body).Decode(&p)
	)

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

	// 	if redis.ErrorResponse(w, r, err, 400) {
	// 	return
	// }

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
