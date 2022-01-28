package gens

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type foo struct {
	id    int
	value int
}

func (f foo) getID() int {
	return f.id
}

func (f foo) getValue() int {
	return f.value
}

func TestAcc(t *testing.T) {
	got := Acc(
		[]foo{
			{value: 1},
			{value: 2},
			{value: 3},
		}, foo.getValue,
	)
	assert.Equal(t, 6, got)
}

func TestContains(t *testing.T) {
	type args struct {
		s     []int
		value int
	}
	tests := map[string]struct {
		args     args
		expected bool
	}{
		"true": {
			args: args{
				s:     []int{1, 2, 3},
				value: 2,
			},
			expected: true,
		},
		"false": {
			args: args{
				s:     []int{1, 2, 3},
				value: 4,
			},
			expected: false,
		},
	}
	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				assert.Equal(t, tt.expected, Contains(tt.args.s, tt.args.value))
			},
		)
	}
}

func TestCount(t *testing.T) {
	got := Count(
		[]int{1, 2, 3, 1}, 1,
	)
	assert.Equal(t, 2, got)
}

func TestDedup(t *testing.T) {
	got := Dedup(
		[]int{2, 1, 3, 2}, func(i int) int {
			return i
		},
	)
	assert.Equal(t, []int{2, 3, 1}, got)
}

func TestFilter(t *testing.T) {
	got := Filter(
		[]int{1, 2, 3}, func(i int) bool {
			return i == 2
		},
	)
	assert.Equal(t, []int{1, 3}, got)
}

func TestJoin(t *testing.T) {
	join := func(i int, i2 int) int {
		return i + i2
	}

	type args struct {
		s1 []int
		s2 []int
	}
	tests := map[string]struct {
		args     args
		expected []int
	}{
		"equal length": {
			args: args{
				s1: []int{1, 2, 3},
				s2: []int{10, 20, 30},
			},
			expected: []int{11, 22, 33},
		},
		"s1 < s2": {
			args: args{
				s1: []int{1, 2},
				s2: []int{10, 20, 30},
			},
			expected: []int{11, 22, 30},
		},
		"s1 > s2": {
			args: args{
				s1: []int{1, 2, 3},
				s2: []int{10, 20},
			},
			expected: []int{11, 22, 3},
		},
	}
	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				assert.Equal(t, tt.expected, Join(tt.args.s1, tt.args.s2, join))
			},
		)
	}
}

func TestMax(t *testing.T) {
	assert.Equal(t, 3, Max([]int{1, 2, 3}, math.MinInt64))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 1, Min([]int{1, 2, 3}, math.MaxInt64))
}

func TestReduce(t *testing.T) {
	got := Reduce(
		[]foo{
			{value: 1},
			{value: 2},
			{value: 3},
		}, func(current foo, agg *foo) {
			agg.value += current.value
		},
	)
	assert.Equal(t, foo{value: 6}, got)
}

func TestSend(t *testing.T) {
	ch := make(chan int, 3)
	var ch2 chan<- int = ch // Temporary fix to overcome IntelliJ limitation
	Send([]int{1, 2, 3}, ch2, true)
	sum := 0
	for v := range ch {
		sum += v
	}
	assert.Equal(t, 6, sum)
}

func TestSub(t *testing.T) {
	type args struct {
		s     []int
		until func(int) bool
	}
	tests := map[string]struct {
		args     args
		expected []int
	}{
		"sub": {
			args: args{
				s: []int{1, 2, 3},
				until: func(i int) bool {
					return i == 2
				},
			},
			expected: []int{1},
		},
		"all": {
			args: args{
				s: []int{1, 2, 3},
				until: func(i int) bool {
					return i == 0
				},
			},
			expected: []int{1, 2, 3},
		},
	}
	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				assert.Equal(t, tt.expected, Sub(tt.args.s, tt.args.until))
			},
		)
	}
}

func TestSum(t *testing.T) {
	got := Sum([]int{1, 2, 3})
	assert.Equal(t, 6, got)
}

func TestToMap(t *testing.T) {
	got := ToMap(
		[]foo{
			{
				id:    1,
				value: 1,
			},
			{
				id:    2,
				value: 2,
			},
		}, foo.getID,
	)
	assert.Equal(
		t, map[int]foo{
			1: {
				id:    1,
				value: 1,
			},
			2: {
				id:    2,
				value: 2,
			},
		}, got,
	)
}
