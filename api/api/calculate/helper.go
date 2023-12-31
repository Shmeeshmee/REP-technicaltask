package calculate

import (
	"fmt"
	"sort"
)

func calculator(target int, numbers []int) Response {
	if len(numbers) == 0 {
		return Response{}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	m := inc(numbers, &target)

	resStr := []string{}
	for k, v := range m {
		if v > 0 {
			resStr = append(resStr, fmt.Sprintf("%d X %d", v, k))
		}
	}

	return Response{
		Packs: resStr,
		Total: target,
	}
}

func inc(a []int, input *int) map[int]int {
	for {
		res := arrs(a, *input)
		if *input < res[len(res)-1].value {
			input = &res[len(res)-1].value
		}

		if m := combine(*input, res, a); m != nil {
			return m
		}

		*input++
	}
}

type packValue struct {
	value int
	final int
	len   int
	br    int
}

func arrs(arr []int, t int) (res []packValue) {
	for i, a := range arr {
		var (
			br = 0
			q  = 0
		)
		if i != 0 {
			if a%2 == 0 {
				q = 5
			} else {
				q = 10
			}
			res[0].br -= q * a
			res[0].br -= q
		} else {
			br = t / a * a
			q = t / a
		}
		res = append(res, packValue{
			value: a,
			final: q * a,
			len:   q,
			br:    br,
		},
		)

	}

	return res
}

func combine(target int, res []packValue, keys []int) map[int]int {
	arr := res[0]
	if arr.final > target {
		arr.final = target / arr.value * arr.value
	}
	for e := arr.final; e >= 0; e -= arr.value {
		if e > target {
			continue
		}
		if arr.br > e {
			break
		}
		// step one check
		// if target-e == 0 && e-arr.value < 0 {
		if target-e == 0 {
			return map[int]int{
				keys[0]: e / arr.value,
			}
		}
		// step two check
		// skip recombining if it's the last arr element
		// break if target is less that the element value
		if len(res) == 1 {
			if target < e {
				break
			}
			continue
		}
		if m := combine((target - e), res[1:], keys[1:]); m != nil {
			m[keys[0]] = e / arr.value
			return m
		}
	}
	return nil
}
