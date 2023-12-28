package calculate

import (
	"re-parteners-tech-task/redis"
	"sort"
	"strconv"
)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Calculate(amount int) ([]string, error) {
	packs := redis.GetValues("pack")
	var (
		val []int
	)
	for _, v := range packs {
		tmp, _ := strconv.Atoi(v)
		val = append(val, tmp)
	}

	sort.Ints(val)

	return calculator(amount, val), nil
}
