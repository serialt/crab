package crab

import (
	"testing"

	"github.com/serialt/crab/internal"
)

func TestContainSubSlice(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestContainSubSlice")
	assert.Equal(true, ContainSubSlice([]string{"a", "a", "b", "c"}, []string{"a", "a"}))
	assert.Equal(false, ContainSubSlice([]string{"a", "a", "b", "c"}, []string{"a", "d"}))
	assert.Equal(true, ContainSubSlice([]int{1, 2, 3}, []int{1, 2}))
	assert.Equal(false, ContainSubSlice([]int{1, 2, 3}, []int{0, 1}))
}

func TestSliceChunk(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestChunk")

	arr := []string{"a", "b", "c", "d", "e"}

	assert.Equal([][]string{}, SliceChunk(arr, -1))

	assert.Equal([][]string{}, SliceChunk(arr, 0))

	r1 := [][]string{{"a"}, {"b"}, {"c"}, {"d"}, {"e"}}
	assert.Equal(r1, SliceChunk(arr, 1))

	r2 := [][]string{{"a", "b"}, {"c", "d"}, {"e"}}
	assert.Equal(r2, SliceChunk(arr, 2))

	r3 := [][]string{{"a", "b", "c"}, {"d", "e"}}
	assert.Equal(r3, SliceChunk(arr, 3))

	r4 := [][]string{{"a", "b", "c", "d"}, {"e"}}
	assert.Equal(r4, SliceChunk(arr, 4))

	r5 := [][]string{{"a", "b", "c", "d", "e"}}
	assert.Equal(r5, SliceChunk(arr, 5))

	r6 := [][]string{{"a", "b", "c", "d", "e"}}
	assert.Equal(r6, SliceChunk(arr, 6))
}

func TestUpdateAt(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestUpdateAt")

	assert.Equal([]string{"a", "b", "c"}, SliceUpdateAt([]string{"a", "b", "c"}, -1, "1"))
	assert.Equal([]string{"1", "b", "c"}, SliceUpdateAt([]string{"a", "b", "c"}, 0, "1"))
	assert.Equal([]string{"a", "b", "2"}, SliceUpdateAt([]string{"a", "b", "c"}, 2, "2"))
	assert.Equal([]string{"a", "b", "c"}, SliceUpdateAt([]string{"a", "b", "c"}, 3, "2"))
}

func TestSliceUnique(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestUnique")

	assert.Equal([]int{1, 2, 3}, SliceUnique([]int{1, 2, 2, 3}))
	assert.Equal([]string{"a", "b", "c"}, SliceUnique([]string{"a", "a", "b", "c"}))
}

func TestSliceUnion(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestUnion")

	s1 := []int{1, 3, 4, 6}
	s2 := []int{1, 2, 5, 6}
	s3 := []int{0, 4, 5, 7}

	assert.Equal([]int{1, 3, 4, 6, 2, 5, 0, 7}, SliceUnion(s1, s2, s3))
	assert.Equal([]int{1, 3, 4, 6, 2, 5}, SliceUnion(s1, s2))
	assert.Equal([]int{1, 3, 4, 6}, SliceUnion(s1))
}

func TestMerge(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestMerge")

	s1 := []int{1, 2, 3, 4}
	s2 := []int{2, 3, 4, 5}
	s3 := []int{4, 5, 6}

	assert.Equal([]int{1, 2, 3, 4, 2, 3, 4, 5, 4, 5, 6}, SliceMerge(s1, s2, s3))
	assert.Equal([]int{1, 2, 3, 4, 2, 3, 4, 5}, SliceMerge(s1, s2))
	assert.Equal([]int{2, 3, 4, 5, 4, 5, 6}, SliceMerge(s2, s3))
}

func TestToSlice(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestToSlice")

	str1 := "a"
	str2 := "b"
	assert.Equal([]string{"a"}, ToSlice(str1))
	assert.Equal([]string{"a", "b"}, ToSlice(str1, str2))
}

func TestSliceDiff(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestSliceDiff")
	slice1 := []int64{1, 2, 3}
	slice2 := []int64{3, 2, 5, 6}
	result := []int64{1}

	assert.Equal(result, SliceDiff(slice1, slice2))
}
