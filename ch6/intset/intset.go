// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

// exercise 6.5
var (
	platformBits = 32 << (^uint(0) >> 63)
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	if s == nil {
		return false
	}
	word, bit := x/platformBits, uint(x%platformBits)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/platformBits, uint(x%platformBits)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < platformBits; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", platformBits*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

//exercise 6.1
func (i *IntSet) Len() int {
	if i == nil {
		return 0
	}
	counter := 0
	for w := range i.words {
		word := i.words[w]
		for i := 0; i < platformBits; i++ {
			counter += int((word >> uint(i)) % 2)
		}
	}
	return counter
}

func (i *IntSet) Remove(x int) {
	if !i.Has(x) {
		return
	}
	word, bit := x/platformBits, uint(x%platformBits)
	i.words[word] ^= (1 << bit)
}

func (i *IntSet) Clear() {
	if i == nil {
		return
	}
	i.words = []uint{}
}

func (i *IntSet) Copy() *IntSet {
	copySet := new(IntSet)
	copySet.words = make([]uint, len(i.words))
	copy(copySet.words, i.words)
	return copySet
}

// excercise 6.2
func (i *IntSet) AddAll(elements ...int) {
	for _, e := range elements {
		i.Add(e)
	}
}

func (i *IntSet) IntersectWith(other *IntSet) {
	numberOfWorlds := len(other.words)
	for w := 0; w < len(i.words); w++ {
		mask := uint(0)
		if w < numberOfWorlds {
			mask = other.words[w]
		}
		i.words[w] &= mask
	}
}

func (i *IntSet) DifferenceWith(other *IntSet) {
	otherNumOfWorld := len(other.words)
	for w := 0; w < len(i.words); w++ {
		mask := uint(0)
		if w < otherNumOfWorld {
			mask = other.words[w]
		}
		//A\B = A \cap !B
		i.words[w] &= ^mask
	}
}

func (i *IntSet) SymmetricDifferenceWith(other *IntSet) {
	otherNumOfWorlds := len(other.words)
	if otherNumOfWorlds > len(i.words) {
		newSlice := make([]uint, otherNumOfWorlds)
		copy(newSlice, i.words)
	}
	for w := 0; w < len(i.words); w++ {
		if w >= otherNumOfWorlds {
			break
		}
		i.words[w] ^= other.words[w]
	}
}

func (i *IntSet) Elems() []int {
	if i == nil {
		return []int{}
	}
	elems := []int{}
	for w := range i.words {
		word := i.words[w]
		for i := 0; i < platformBits; i++ {
			if (word>>uint(i))%2 == 1 {
				elems = append(elems, w*platformBits+i)
			}
		}
	}
	return elems

}
