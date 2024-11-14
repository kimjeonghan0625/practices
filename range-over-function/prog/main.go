package main

import (
	"fmt"
	"iter"
	ds "range-over-function/data-structure"
	"slices"
)

func makeUnion[E comparable](s1, s2 *ds.Set[E]) *ds.Set[E] {
	s3 := ds.New[E]()
	for v := range s1.All() {
		s3.Add(v)
	}
	for v := range s2.All() {
		s3.Add(v)
	}

	return s3
}

// pull iterator로 zip을 구현해보자
// 여러 개의 슬라이스를 받아서 슬라이스의 각 요소들을 zip한 슬라이스를 하니씩 yield하는 이터레이터 반환
// func zip
func Zip[E comparable](s ...iter.Seq[E]) iter.Seq[[]E] {
	type PullIter[E comparable] struct {
		next func() (v E, ok bool)
		stop func()
	}
	pullIterSlice := make([]*PullIter[E], 0)

	// pull iter로의 변환
	for _, v := range s {
		next, stop := iter.Pull(v)
		pullIter := &PullIter[E]{next, stop}
		pullIterSlice = append(pullIterSlice, pullIter)
	}

	// pull iter를 통해 밸류를 꺼내와서 합치기
	return func(yield func([]E) bool) {
		for _, v := range pullIterSlice {
			defer v.stop()
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

func main() {
	s1 := []string{"hello", "jh", "apple"}
	s2 := []string{"go", "kim", "mango", "hihihi"}
	s3 := []string{"world", "andy", "pineapple", "wow"}
	// Zip(slices.Values(s1), slices.Values(s2))(
	// 	func (sl []string) bool {
	// 		fmt.Println(sl)
	// 		return true
	// 	},
	// )
	sv := slices.Values[[]string, string]
	for v := range Zip(slices.Values(s1), slices.Values(s2), sv(s3)) {
		fmt.Println(v)
	}
	fmt.Println(slices.Collect(Zip(slices.Values(s1), slices.Values(s2), sv(s3))))

	// s1 := ds.New[int]()
	// for i := range 5 {
	// 	s1.Add(i)
	// }
	// push
	// s1.Push(
	// 	func(i int) bool {
	// 		fmt.Println(i)
	// 		return true
	// 	},
	// )

	// pull
	// next, _ := s1.Pull()
	// for v, ok := next(); ok; v, ok = next() {
	// 	fmt.Println(v)
	// }

	// s1.All()(
	// 	func (i int) bool {
	// 		if i == 3 {
	// 			return true
	// 		}
	// 		fmt.Println(i)
	// 		return true
	// 	},
	// )

	// fmt.Println("//////")

	// for v := range s1.All() {
	// 	if v == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(v)
	// }

	// s2 := ds.New[int]()
	// for i := range 3 {
	// 	s2.Add(i + 8)
	// }

	// for _, v := range slices.Sorted(makeUnion(s1, s2).All()) {
	// 	fmt.Println(v)
	// }
}
