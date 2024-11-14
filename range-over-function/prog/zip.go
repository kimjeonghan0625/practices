package main

import (
	"iter"
)

type PullIter[E comparable] struct {
	next func() (v E, ok bool)
	stop func()
}

func Zip[E comparable](s ...iter.Seq[E]) iter.Seq[[]E] {
	return func(yield func([]E) bool) {
		pullIterSlice := make([]*PullIter[E], 0)
		for _, v := range s {
			next, stop := iter.Pull(v)
			pullIter := &PullIter[E]{next, stop}
			defer pullIter.stop()
			pullIterSlice = append(pullIterSlice, pullIter)
		}
		for {
			tmp := make([]E, 0)
			for _, p := range pullIterSlice {
				v, ok := p.next()
				if !ok {
					return
				}
				tmp = append(tmp, v)
			}
			if !yield(tmp) {
				return
			}
		}
	}
}