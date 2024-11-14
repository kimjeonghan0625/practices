package main

import (
	ds "range-over-function/data-structure"
)

func MakeUnion[E comparable](s1, s2 *ds.Set[E]) *ds.Set[E] {
	s3 := ds.New[E]()
	for v := range s1.All() {
		s3.Add(v)
	}
	for v := range s2.All() {
		s3.Add(v)
	}

	return s3
}