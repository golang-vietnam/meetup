package main

import "fmt"

func main() {
	s1 := make([]int, 3, 4)
	copy(s1, []int{1, 2, 3})
	fmt.Println(len(s1), cap(s1), &s1[0])

	s1 = append(s1, 4)
	fmt.Println(len(s1), cap(s1), &s1[0])

	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), &s2[0])

	s1 = append(s1, 5)
	fmt.Println(len(s1), cap(s1), &s1[0])
}
