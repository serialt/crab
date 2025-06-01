package crab

import (
	"slices"
)

// ContainSubSlice check if the slice contain a given subslice or not.
// Play: https://go.dev/play/p/bcuQ3UT6Sev
func ContainSubSlice[T comparable](slice, subSlice []T) bool {
	for _, v := range subSlice {

		if !slices.Contains(slice, v) {
			return false
		}
	}

	return true
}

// Chunk creates a slice of elements split into groups the length of size.
// Play: https://go.dev/play/p/b4Pou5j2L_C
func SliceChunk[T any](slice []T, size int) [][]T {
	result := [][]T{}

	if len(slice) == 0 || size <= 0 {
		return result
	}

	for _, item := range slice {
		l := len(result)
		if l == 0 || len(result[l-1]) == size {
			result = append(result, []T{})
			l++
		}

		result[l-1] = append(result[l-1], item)
	}

	return result
}

// Equal checks if two slices are equal: the same length and all elements' order and value are equal.
// Play: https://go.dev/play/p/WcRQJ37ifPa
func Equal[T comparable](slice1, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

// UpdateAt update the slice element at index.
// Play: https://go.dev/play/p/f3mh2KloWVm
func SliceUpdateAt[T any](slice []T, index int, value T) []T {
	size := len(slice)

	if index < 0 || index >= size {
		return slice
	}
	slice = append(slice[:index], append([]T{value}, slice[index+1:]...)...)

	return slice
}

// Unique remove duplicate elements in slice.
// Play: https://go.dev/play/p/AXw0R3ZTE6a
func SliceUnique[T comparable](slice []T) []T {
	result := []T{}

	for i := 0; i < len(slice); i++ {
		v := slice[i]
		skip := true
		for j := range result {
			if v == result[j] {
				skip = false
				break
			}
		}
		if skip {
			result = append(result, v)
		}
	}

	return result
}

// Union creates a slice of unique elements, in order, from all given slices.
// Play: https://go.dev/play/p/hfXV1iRIZOf
func SliceUnion[T comparable](slices ...[]T) []T {
	result := []T{}
	contain := map[T]struct{}{}

	for _, slice := range slices {
		for _, item := range slice {
			if _, ok := contain[item]; !ok {
				contain[item] = struct{}{}
				result = append(result, item)
			}
		}
	}

	return result
}

// Merge all given slices into one slice.
// Play: https://go.dev/play/p/lbjFp784r9N
func SliceMerge[T any](slices ...[]T) []T {
	result := make([]T, 0)

	for _, v := range slices {
		result = append(result, v...)
	}

	return result
}

// ToSlice returns a slices of a variable parameter transformation.
// Play: https://go.dev/play/p/YzbzVq5kscN
func ToSlice[T any](items ...T) []T {
	result := make([]T, len(items))
	copy(result, items)

	return result
}

// // Random get a random item of slice, return idx=-1 when slice is empty
// // Play: https://go.dev/play/p/UzpGQptWppw
// func Random[T any](slice []T) (val T, idx int) {
// 	if len(slice) == 0 {
// 		return val, -1
// 	}

// 	idx = random.RandInt(0, len(slice))
// 	return slice[idx], idx
// }

func SliceDiff[T comparable](slice, comparedSlice []T) []T {
	result := []T{}
	for _, v := range slice {
		if !slices.Contains(comparedSlice, v) {
			result = append(result, v)
		}
	}
	return result
}
