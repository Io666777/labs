package main

import "fmt"

func Find[T any](slc []T, predicate func(T) bool) (T, bool) {
	for _, value := range slc {
		if predicate(value) {
			return value, true
		}
	}
	var znach T
	return znach, false
}

func MultiPredicateFilter[T any](slc []T, predicates ...func(T) bool) []T {
	var res []T

	for _, val := range slc {
		ok := true
		for _, p := range predicates {
			if !p(val) {
				ok = false
				break
			}
		}
		if ok {
			res = append(res, val)
		}
	}

	return res
}

func GroupAndAggregate[T any, K comparable](data []T, key func(T) K, agg func([]T) float64) map[K]float64 {
	m := make(map[K][]T)
	for _, v := range data {
		m[key(v)] = append(m[key(v)], v)
	}

	res := make(map[K]float64)
	for k, v := range m {
		res[k] = agg(v)
	}
	return res
}

type P struct {
	Name string
	Num  int
}

func main() {
	var a int = 1

	switch a {
	case 1:
		num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		res, ok := Find(num, func(n int) bool {
			return n > 5
		})
		if ok {
			fmt.Println("yes", res)
		} else {
			fmt.Println("no")
		}
		frut := []string{"apple", "banana", "car", "lemon"}
		frt, fok := Find(frut, func(s string) bool {
			return len(s) > 0 && s[0] == 'a'
		})
		if fok {
			fmt.Println("yes - a", frt)
		} else {
			fmt.Println("no")
		}

	case 2:
		pr := []P{
			{"car", 1},
			{"cat", 2},
			{"dog", 3},
		}

		chep := MultiPredicateFilter(pr,
			func(p P) bool { return p.Num < 2 },
			func(p P) bool { return p.Name[0] == 'd' },
		)

		fmt.Println("res - ", chep)

	case 3:
		pr := []P{
			{"dada", 2},
			{"hehe", 3},
			{"dada", 5},
			{"hehe", 10},
		}

		res := GroupAndAggregate(
			pr,
			func(p P) string { return p.Name },
			func(lines []P) float64 {
				sum := 0
				for _, s := range lines {
					sum += s.Num
				}
				return float64(sum)
			},
		)
		fmt.Println(res)
	}
}
