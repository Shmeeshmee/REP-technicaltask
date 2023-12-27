package calculate

import (
	"fmt"
	"math"
	"sort"
)

// func init() {
// 	redis.Set("init", PackKey, 1)
// }

// func incrID() {
// 	redis.Incr("init", PackKey)
// }

// func setPack(value float64) {
// 	ID := redis.Get("init", PackKey)
// 	redis.Set(PackKey, ID, value)
// }

// Within the constraints of Rule 1 above, send out no more items than necessary to
// fulfil the order.

func doAll(target int, numbers []int) (result []string) {
	result = two(target, numbers)
	if len(result) == 0 {
		if target < numbers[0] {
			return doAll(numbers[0], numbers)
		} else {
			return doAll(target+1, numbers)
		}
	}
	return result
}

func two(target int, numbers []int) []string {
	sort.Ints(numbers)
	var results [][]int
	var currentCombination []int
	l := math.MaxInt64

	findCombinationsHelper(target, numbers, 0, currentCombination, &results)

	for _, result := range results {
		if len(result) < l {
			l = len(result)
			currentCombination = result
		}
	}

	if len(currentCombination) == 0 {
		return nil
	}

	return stringify(currentCombination)
}

func stringify(a []int) (result []string) {
	m := map[int]int{}
	for _, i := range a {
		m[i]++
	}

	for k, v := range m {
		result = append(result, fmt.Sprintf("%d X %d", k, v))
	}

	return
}

func findCombinationsHelper(target int, numbers []int, startIndex int, currentCombination []int, result *[][]int) {
	if target == 0 {
		// Found a valid combination
		*result = append(*result, append([]int{}, currentCombination...))
		return
	}

	for i := startIndex; i < len(numbers); i++ {
		if target >= numbers[i] {
			// Include the current number in the combination
			currentCombination = append(currentCombination, numbers[i])

			// Recursively find combinations with the updated target
			findCombinationsHelper(target-numbers[i], numbers, i, currentCombination, result)

			// Remove the last element to backtrack
			currentCombination = currentCombination[:len(currentCombination)-1]
		}
	}
}

