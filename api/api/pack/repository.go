package pack

import (
	"fmt"
	"re-parteners-tech-task/redis"
	"strconv"
)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetByKey(packKey string) (*Pack, error) {
	s := redis.Get(PackKey, packKey)

	f, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("cannot parse float")
		return nil, err

	}
	return &Pack{
		Size: f,
	}, nil
}

func (r *Repository) Set(p Pack) {
	setPack(p.Size)
	incrID()
}

func (r *Repository) GetAll() map[string]string {
	return redis.GetValues(PackKey)
}

func (r *Repository) DeleteByID(packID string) error {
	fmt.Println("entered DeleteByID pack repository")
	val, err := redis.Delete(PackKey, packID)
	fmt.Println(val)
	return err
}
