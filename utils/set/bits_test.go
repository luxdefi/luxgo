// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package set

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Bits_New(t *testing.T) {
	tests := []struct {
		name   string
		bits   []int
		length int
	}{
		{
			name:   "empty",
			bits:   []int{},
			length: 0,
		},
		{
			name:   "populated",
			bits:   []int{0, 9, 99, 999, 9999},
			length: 10_000,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
<<<<<<< HEAD
			require := require.New(t)
			b := NewBits(test.bits...)

			for _, bit := range test.bits {
				require.True(b.Contains(bit))
			}

			require.Equal(test.length, b.Len())
=======
			r := require.New(t)
			b := NewBits(test.bits...)

			for _, bit := range test.bits {
				r.True(b.Contains(bit))
			}

			r.Equal(test.length, b.Len())
>>>>>>> 483d9bd18 (Move bit sets to the set package (#2365))
		})
	}
}

func Test_Bits_AddRemove(t *testing.T) {
	tests := []struct {
		name             string
		toAdd            []int
		toRemove         []int
		expectedElements []int
		expectedLen      int
	}{
		{
			name:             "empty sets",
			toAdd:            []int{},
			toRemove:         []int{},
			expectedElements: []int{}, // []
			expectedLen:      0,
		},
		{
			name:             "add only",
			toAdd:            []int{0, 1, 2},
			toRemove:         []int{},
			expectedElements: []int{0, 1, 2}, // [1, 1, 1]
			expectedLen:      3,
		},
		{
			name:             "remove left-most",
			toAdd:            []int{0, 1, 2},
			toRemove:         []int{0},
			expectedElements: []int{1, 2}, // [1, 1, 0]
			expectedLen:      3,
		},
		{
			name:             "remove middle",
			toAdd:            []int{0, 1, 2},
			toRemove:         []int{1},
			expectedElements: []int{2, 0}, // [1, 0, 1]
			expectedLen:      3,
		},
		{
			name:             "remove right-most",
			toAdd:            []int{0, 1, 2},
			toRemove:         []int{2},
			expectedElements: []int{0, 1}, // [1, 1]
			expectedLen:      2,
		},
		{
			name:             "remove all",
			toAdd:            []int{0, 1, 2},
			toRemove:         []int{0, 1, 2},
			expectedElements: []int{}, // [1, 1, 1]
			expectedLen:      0,
		},
		{
			name:             "remove reverse-order",
			toAdd:            []int{0, 1, 2},
			toRemove:         []int{2, 1, 0},
			expectedElements: []int{}, // []
			expectedLen:      0,
		},
		{
			name:             "remove non-existent elements",
			toAdd:            []int{0, 1, 2},
			toRemove:         []int{3, 4, 5},
			expectedElements: []int{0, 1, 2}, // [1, 1, 1]
			expectedLen:      3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
<<<<<<< HEAD
			require := require.New(t)
=======
			r := require.New(t)
>>>>>>> 483d9bd18 (Move bit sets to the set package (#2365))
			b := NewBits()

			for _, add := range test.toAdd {
				b.Add(add)
			}

			for _, remove := range test.toRemove {
				b.Remove(remove)
			}

			for _, element := range test.expectedElements {
<<<<<<< HEAD
				require.True(b.Contains(element))
			}

			require.Equal(test.expectedLen, b.Len())
=======
				r.True(b.Contains(element))
			}

			r.Equal(test.expectedLen, b.Len())
>>>>>>> 483d9bd18 (Move bit sets to the set package (#2365))
		})
	}
}

func Test_Bits_Union(t *testing.T) {
	tests := []struct {
		name        string
		left        []int
		right       []int
		expected    []int
		expectedLen int
	}{
		{
			name:        "empty sets",
			left:        []int{},
			right:       []int{},
			expected:    []int{}, // []
			expectedLen: 0,
		},
		{
			name:        "left and right are same",
			left:        []int{2, 1, 0},
			right:       []int{2, 1, 0},
			expected:    []int{2, 1, 0}, // [1, 1, 1]
			expectedLen: 3,
		},
		{
			name:        "left and no right",
			left:        []int{2, 1, 0},
			right:       []int{},
			expected:    []int{2, 1, 0}, // [1, 1, 1]
			expectedLen: 3,
		},
		{
			name:        "right and no left",
			left:        []int{},
			right:       []int{2, 1, 0},
			expected:    []int{2, 1, 0}, // [1, 1, 1]
			expectedLen: 3,
		},
		{
			name:        "left and right overlap",
			left:        []int{2, 1},
			right:       []int{1, 0},
			expected:    []int{2, 1, 0}, // [1, 1, 1]
			expectedLen: 3,
		},
		{
			name:        "left and right overlap different sizes",
			left:        []int{5, 3, 1},
			right:       []int{8, 6, 4, 2, 0},
			expected:    []int{8, 6, 5, 4, 3, 2, 1, 0}, // [1, 0, 1, 1, 1, 1, 1, 1, 1]
			expectedLen: 9,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
<<<<<<< HEAD
			require := require.New(t)
=======
			r := require.New(t)
>>>>>>> 483d9bd18 (Move bit sets to the set package (#2365))
			b := NewBits()

			for _, add := range test.left {
				b.Add(add)
			}
			for _, add := range test.right {
				b.Add(add)
			}

			for _, element := range test.expected {
<<<<<<< HEAD
				require.True(b.Contains(element))
			}

			require.Equal(test.expectedLen, b.Len())
=======
				r.True(b.Contains(element))
			}

			r.Equal(test.expectedLen, b.Len())
>>>>>>> 483d9bd18 (Move bit sets to the set package (#2365))
		})
	}
}

func Test_Bits_Intersection(t *testing.T) {
	tests := []struct {
		name        string
		left        []int
		right       []int
		expected    []int
		expectedLen int
	}{
		{
			name:        "empty sets",
			left:        []int{},
			right:       []int{},
			expected:    []int{}, // []
			expectedLen: 0,
		},
		{
			name:        "left and right are same",
			left:        []int{2, 1, 0},
			right:       []int{2, 1, 0},
			expected:    []int{2, 1, 0}, // [1, 1, 1]
			expectedLen: 3,
		},
		{
			name:        "left and no right",
			left:        []int{2, 1, 0},
			right:       []int{},
			expected:    []int{}, // []
			expectedLen: 0,
		},
		{
			name:        "right and no left",
			left:        []int{},
			right:       []int{2, 1, 0},
			expected:    []int{}, // []
			expectedLen: 0,
		},
		{
			name:        "left and right overlap",
			left:        []int{2, 1},
			right:       []int{1, 0},
			expected:    []int{1}, // [1, 0]
			expectedLen: 2,
		},
		{
			name:        "left and right overlap different sizes",
			left:        []int{5, 3, 1},
			right:       []int{8, 6, 4, 2, 0},
			expected:    []int{}, // []
			expectedLen: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
<<<<<<< HEAD
			require := require.New(t)
=======
			r := require.New(t)
>>>>>>> 483d9bd18 (Move bit sets to the set package (#2365))
			left := NewBits()
			right := NewBits()
			for _, add := range test.left {
				left.Add(add)
			}
			for _, add := range test.right {
				right.Add(add)
			}

			left.Intersection(right)

			expected := NewBits()
			for _, element := range test.expected {
				expected.Add(element)
			}

<<<<<<< HEAD
			require.ElementsMatch(left.bits.Bits(), expected.bits.Bits())
=======
			r.ElementsMatch(left.bits.Bits(), expected.bits.Bits())
>>>>>>> 483d9bd18 (Move bit sets to the set package (#2365))
		})
	}
}

func Test_Bits_Difference(t *testing.T) {
	tests := []struct {
		name        string
		left        []int
		right       []int
		expected    []int
		expectedLen int
	}{
		{
			name:        "empty sets",
			left:        []int{},
			right:       []int{},
			expected:    []int{}, // []
			expectedLen: 0,
		},
		{
			name:        "left and right are same",
			left:        []int{2, 1, 0},
			right:       []int{2, 1, 0},
			expected:    []int{}, // []
			expectedLen: 0,
		},
		{
			name:        "left and no right",
			left:        []int{2, 1, 0},
			right:       []int{},
			expected:    []int{2, 1, 0}, // [1, 1, 1]
			expectedLen: 3,
		},
		{
			name:        "right and no left",
			left:        []int{},
			right:       []int{2, 1, 0},
			expected:    []int{}, // []
			expectedLen: 3,
		},
		{
			name:        "left and right overlap",
			left:        []int{2, 1},
			right:       []int{1, 0},
			expected:    []int{2}, // [1, 0, 0]
			expectedLen: 3,
		},
		{
			name:        "left and right overlap different sizes",
			left:        []int{5, 3, 1},
			right:       []int{8, 6, 4, 2, 0},
			expected:    []int{5, 3, 1}, // [1, 0, 1, 0, 1, 0]
			expectedLen: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
<<<<<<< HEAD
			require := require.New(t)
=======
			r := require.New(t)
>>>>>>> 483d9bd18 (Move bit sets to the set package (#2365))
			left := NewBits()
			right := NewBits()
			for _, add := range test.left {
				left.Add(add)
			}
			for _, add := range test.right {
				right.Add(add)
			}

			left.Difference(right)

			expected := NewBits()
			for _, element := range test.expected {
				expected.Add(element)
			}

<<<<<<< HEAD
			require.ElementsMatch(left.bits.Bits(), expected.bits.Bits())
=======
			r.ElementsMatch(left.bits.Bits(), expected.bits.Bits())
>>>>>>> 483d9bd18 (Move bit sets to the set package (#2365))
		})
	}
}

func Test_Bits_Clear(t *testing.T) {
	tests := []struct {
		name   string
		bitset []int
	}{
		{
			name:   "empty",
			bitset: []int{}, // []
		},
		{
			name:   "populated",
			bitset: []int{5, 4, 3, 2, 1}, // [1, 1, 1, 1, 1]
		},
		{
			name:   "populated - big",
			bitset: []int{255}, // [1, 0...]
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
<<<<<<< HEAD
			require := require.New(t)
=======
			r := require.New(t)
>>>>>>> 483d9bd18 (Move bit sets to the set package (#2365))
			b := NewBits()

			for bit := range test.bitset {
				b.Add(bit)
			}

			b.Clear()

<<<<<<< HEAD
			require.Zero(b.Len())
=======
			r.Zero(b.Len())
>>>>>>> 483d9bd18 (Move bit sets to the set package (#2365))
		})
	}
}

func Test_Bits_String(t *testing.T) {
	tests := []struct {
		name     string
		bitset   []int
		expected string
	}{
		{
			name:     "empty",
			bitset:   []int{},
			expected: "", // []
		},
		{
			name:     "populated",
			bitset:   []int{7, 6, 5, 4, 3, 2, 1, 0}, // [1, 1, 1, 1, 1, 1, 1, 1]
			expected: "ff",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
<<<<<<< HEAD
			require := require.New(t)
=======
			r := require.New(t)
>>>>>>> 483d9bd18 (Move bit sets to the set package (#2365))
			b := NewBits()

			for _, bit := range test.bitset {
				b.Add(bit)
			}

<<<<<<< HEAD
			require.Equal(test.expected, b.String())
=======
			r.Equal(test.expected, b.String())
>>>>>>> 483d9bd18 (Move bit sets to the set package (#2365))
		})
	}
}

func Test_Bits_HammingWeight(t *testing.T) {
	tests := []struct {
		name     string
		bitset   []int
		expected int
	}{
		{
			name:     "empty",
			bitset:   []int{}, // []
			expected: 0,
		},
		{
			name:     "populated - more than one word",
			bitset:   []int{255}, // [1, 0...]
			expected: 1,
		},
		{
			name:     "populated - all ones",
			bitset:   []int{5, 4, 3, 2, 1, 0}, // [1, 1, 1, 1, 1, 1]
			expected: 6,
		},
		{
			name:     "populated - trailing zeroes",
			bitset:   []int{5, 4, 3}, // [1, 1, 1, 0, 0, 0]
			expected: 3,
		},
		{
			name:     "populated - interwoven 1",
			bitset:   []int{4, 2, 0}, // [1, 0, 1, 0, 1]
			expected: 3,
		},
		{
			name:     "populated - interwoven 2",
			bitset:   []int{3, 1}, // [1, 0, 1, 0]
			expected: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
<<<<<<< HEAD
			require := require.New(t)
=======
			r := require.New(t)
>>>>>>> 483d9bd18 (Move bit sets to the set package (#2365))
			b := NewBits()

			for _, bit := range test.bitset {
				b.Add(bit)
			}

<<<<<<< HEAD
			require.Equal(test.expected, b.HammingWeight())
=======
			r.Equal(test.expected, b.HammingWeight())
>>>>>>> 483d9bd18 (Move bit sets to the set package (#2365))
		})
	}
}

func Test_Bits_Bytes(t *testing.T) {
	require := require.New(t)

	type test struct {
		name string
		elts []int
	}

	tests := []test{
		{
			name: "empty",
			elts: []int{},
		},
		{
			name: "single; element > 63",
			elts: []int{1337},
		},
		{
			name: "multiple",
			elts: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBits(tt.elts...)
			bytes := b.Bytes()
			fromBytes := BitsFromBytes(bytes)

			require.Equal(len(tt.elts), fromBytes.HammingWeight())
			for _, elt := range tt.elts {
				require.True(fromBytes.Contains(elt))
			}
		})
	}
}
