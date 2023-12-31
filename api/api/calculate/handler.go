package calculate

import (
	"fmt"
	"net/http"
	"re-parteners-tech-task/redis"
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
		redis.JsonResponse(
			w,
			redis.CreateResponse(
				nil,
				"please enter number with numeric values / digits [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]",
				err,
				405,
			))
		return
	}

	if iAmount <= 0 {
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

	vals, err := h.r.Calculate(iAmount)
	if redis.ErrorResponse(w, r, err, 400) {
		return
	}

	redis.JsonResponse(
		w,
		redis.CreateResponse(
			vals,
			fmt.Sprintf("successfully calculated for input [%d]", iAmount),
			nil,
			200,
		))
}
