package calculate

import (
	"fmt"
	"math"
)

func calculator(target int, numbers []int) (result []string) {
	result = exact(target, numbers)

	if len(result) == 0 {
		if target < numbers[0] {
			return calculator(numbers[0], numbers)
		} else {
			return calculator(target+1, numbers)
		}
	}
	return result
}

func exact(target int, numbers []int) []string {
	var (
		results            = []int{}
		currentCombination = []int{}
		l                  = math.MaxInt64
	)
	currentCombination = findCombinationsHelper(target, numbers, 0, currentCombination, results, &l)
	// in case there isn't an exact match return nil
	if len(currentCombination) == 0 {
		return nil
	}
	// if the numbers add up exactly to target
	return stringify(currentCombination)
}

func stringify(a []int) (result []string) {
	m := map[int]int{}
	for _, i := range a {
		m[i]++
	}
	for v, k := range m {
		result = append(result, fmt.Sprintf("%d X %d", k, v))
	}

	return
}

func findCombinationsHelper(target int, numbers []int, startIndex int, currentCombination []int, shortestResult []int, shortestLength *int) []int {
	if target == 0 {
		// Found a valid combination
		if len(currentCombination) < *shortestLength {
			*shortestLength = len(currentCombination)
			shortestResult = currentCombination
		}
		return shortestResult
	}

	for i := startIndex; i < len(numbers); i++ {
		if target >= numbers[i] {
			// Include the current number in the combination
			currentCombination = append(currentCombination, numbers[i])
			// Recursively find combinations with the updated target
			shortestResult = findCombinationsHelper(target-numbers[i], numbers, i, currentCombination, shortestResult, shortestLength)
			// Remove the last element to backtrack
			currentCombination = currentCombination[:len(currentCombination)-1]
		}
	}

	return shortestResult
}
