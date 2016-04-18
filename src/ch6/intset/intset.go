/*
Package intset provides a set of integers based on a bit vector

Exercise 6.5 - Convert uint64 to uint
*/

package intset

import (
    "bytes"
    "fmt"
)

var bitsize = (32 << (^uint(0) >> 63))

type IntSet struct {
    words []uint
}

func (s *IntSet) Has(x int) bool {
    word, bit := x/bitsize, uint(x%bitsize)
    return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
    word, bit := x/bitsize, uint(x%bitsize)
    for word >= len(s.words) {
        s.words = append(s.words, 0)
    }
    s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
    for i, tword := range t.words {
        if i < len(s.words) {
            s.words[i] |= tword
        } else {
            s.words = append(s.words, tword)
        }
    }
}

func (s *IntSet) String() string {
    var buf bytes.Buffer
    buf.WriteByte('{')
    for i, word := range s.words {
        if word == 0 {
            continue
        }
        for j := 0; j < bitsize; j++ {
            if word&(1<<uint(j)) != 0 {
                if buf.Len() > len("{") {
                    buf.WriteByte(' ')
                }
                fmt.Fprintf(&buf, "%d", bitsize*i+j)
            }
        }
    }
    buf.WriteByte('}')
    return buf.String()
}

// Exercise 6.1
func (s *IntSet) Len() int {
    var total int
    for _, word := range s.words {
        total += countOnes(word)
    }
    return total
}

func (s *IntSet) Remove(x int) {
    if s.Has(x) {
        word, bit := x/bitsize, uint(x%bitsize)
        s.words[word] &= ^(1 << bit)
    }
}

func (s *IntSet) Clear() {
    s.words = []uint{}
}

func (s *IntSet) Copy() *IntSet {
    var res IntSet
    res.UnionWith(s)
    return &res
}

func countOnes(x uint) (sum int) {
    for ;x != 0; x &= (x-1) {
        sum++
    }
    return
}

// Exercise 6.2
func (s *IntSet) AddAll(xs...int) {
    add := s.Add
    for _, x := range xs {
        add(x)
    }
}

// Exercise 6.3
func (s *IntSet) IntersectWith(t *IntSet) {
        for i, tword := range t.words {
        if i < len(s.words) {
            s.words[i] &= tword
        } else {
            return
        }
    }
}

func (s *IntSet) DifferenceWith(t *IntSet) {
    for i, tword := range t.words {
        if i < len(s.words) {
            s.words[i] ^= (tword & s.words[i])
        } else {
            return
        }
    }
}

func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
    for i, tword := range t.words {
        if i < len(s.words) {
            s.words[i] ^= tword
        } else {
            s.words = append(s.words, tword)
        }
    }
}

// Exercise 6.4
func (s *IntSet) Elem() []int {
    var elements []int
    for i, word := range s.words {
        if word == 0 {
            continue
        }
        for j := 0; j < bitsize; j++ {
            if word&(1<<uint(j)) != 0 {
                 elements = append(elements, bitsize*i+j)
            }
        }
    }
    return elements
}
