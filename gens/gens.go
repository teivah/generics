// Package gens contains generics utilities for slice.
package gens

import (
	"constraints"

	genc "github.com/teivah/generics/constraints"
)

// Acc returns an accumulated result.
func Acc[T any, O constraints.Ordered](s []T, acc func(T) O) O {
	var sum O
	for _, v := range s {
		sum += acc(v)
	}
	return sum
}

// Contains returns whether a slice contains a specific value.
func Contains[C comparable](s []C, value C) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}

// Count returns the number of elements that matches a given value.
func Count[C comparable](s []C, value C) int {
	i := 0
	for _, v := range s {
		if v == value {
			i++
		}
	}
	return i
}

// Dedup deduplicates a slice of elements.
func Dedup[T any, C comparable](s []T, idGetter func(T) C) []T {
	res := make([]T, 0, len(s))

	set := make(map[C]struct{}, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		id := idGetter(s[i])
		if _, exists := set[id]; exists {
			continue
		}
		set[id] = struct{}{}
		res = append(res, s[i])
	}

	return res
}

// Filter filters a slice based on a filter function.
func Filter[T any](s []T, filter func(T) bool) []T {
	res := make([]T, 0, len(s))
	for _, v := range s {
		if !filter(v) {
			res = append(res, v)
		}
	}
	return res
}

// Join joins two slices.
// If the two slices don't have the same length, it will have the length of the max slice.
func Join[T any](s1, s2 []T, join func(t1, t2 T) T) []T {
	max := len(s1)
	if len(s2) > max {
		max = len(s2)
	}

	res := make([]T, max)
	for i := 0; i < len(s1); i++ {
		if i < len(s2) {
			res[i] = join(s1[i], s2[i])
		} else {
			res[i] = s1[i]
		}
	}
	for i := len(s1); i < len(s2); i++ {
		res[i] = s2[i]
	}
	return res
}

// Max returns the maximum.
func Max[O constraints.Ordered](s []O, min O) O {
	for _, v := range s {
		if v > min {
			min = v
		}
	}
	return min
}

// Min returns the minimum.
func Min[O constraints.Ordered](s []O, max O) O {
	for _, v := range s {
		if v < max {
			max = v
		}
	}
	return max
}

// Reduce reduces a slice of elements based on a reduce function.
func Reduce[T any](s []T, reduce func(current T, agg *T)) T {
	var res T
	for _, v := range s {
		reduce(v, &res)
	}
	return res
}

// Send sends all the elements of a slice to a given channel.
func Send[T any](s []T, ch chan<- T, withClosure bool) {
	for _, v := range s {
		ch <- v
	}
	if withClosure {
		close(ch)
	}
}

// Sub returns a subslice based on a function.
func Sub[T any](s []T, until func(T) bool) []T {
	for i, v := range s {
		if until(v) {
			return s[:i]
		}
	}
	return s
}

// Sum sum all the elements of a slice.
func Sum[N genc.Number](s []N) N {
	var sum N
	for _, v := range s {
		sum += v
	}
	return sum
}

// ToMap converts a slice to a map based on a key getter function.
func ToMap[T any, C comparable](s []T, keyGetter func(T) C) map[C]T {
	m := make(map[C]T, len(s))
	for _, v := range s {
		m[keyGetter(v)] = v
	}
	return m
}
