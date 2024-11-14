package main

import (
	"fmt"
	"slices"
)

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
