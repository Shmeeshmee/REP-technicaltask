package pack

import (
	"re-parteners-tech-task/redis"
	"strconv"
)

func init() {
	redis.Set("init", PackKey, 1)
}

func incrID() {
	redis.Incr("init", PackKey)
}

func setPack(value int) {
	ID := redis.Get("init", PackKey)
	redis.Set(PackKey, ID, value)
}

func exists(size int) bool {
	for _, v := range redis.GetValues("pack") {
		if v == strconv.Itoa(size) {
			return true
		}
	}
	return false
}
