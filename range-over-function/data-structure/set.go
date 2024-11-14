package datastructure

import "iter"

type Set[E comparable] struct {
	m map[E]struct{}
}

func New[E comparable]() *Set[E] {
	return &Set[E]{
		make(map[E]struct{}),
	}
}

func (s *Set[E]) Add(val E) {
	s.m[val] = struct{}{}
}

func (s *Set[E]) Contains(val E) bool {
	_, ok := s.m[val]
	return ok
}

// Union returns the union of two sets.
func Union[E comparable](s1, s2 *Set[E]) *Set[E] {
	r := New[E]()
	// Note for/range over internal Set field m.
	// We are looping over the maps in s1 and s2.
	for v := range s1.m {
		r.Add(v)
	}
	for v := range s2.m {
		r.Add(v)
	}
	return r
}

// 그런데 만약에 union같은 함수가 이 패키지 바깥에서 정의된다면?
// Set타입의 unexported field에 접근할 수 없어서 위와 같은 방식으로 순회할 수 없다.
// 하지만 어떤 container가 포함하고 있는 필드를 외부에서 순회할 필요가 있는 경우는 많으므로,
// 타입은 사용자가 패키지 밖에서도 비공개 필드를 순회할 수 있는 메서드를 제공해야 한다.

// func (s *Set[E]) Push(f func(E) bool) {
// 	for v := range s.m {
// 		if !f(v) {
// 			return
// 		}
// 	}
// }

// func (s *Set[E]) Pull() (func() (E, bool), func()) {
// 	ch1 := make(chan E)
// 	ch2 := make(chan bool)

// 	go func() {
// 		defer close(ch1)
// 		for v := range s.m {
// 			select {
// 			case ch1 <- v:
// 			case <-ch2:
// 				return
// 			}
// 		}
// 	}()

// 	next := func() (E, bool) {
// 		v, ok := <-ch1
// 		return v, ok
// 	}

// 	stop := func() {
// 		close(ch2)
// 	}

// 	return next, stop
// }

func (s *Set[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range s.m {
			if !yield(v) {
				return
			}
		}
	}
}
