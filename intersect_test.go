package main

import "testing"

var x = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
var y = []int{2, 3, 5, 7, 11, 6, 12, 13, 14, 15, 16, 17, 18}

func BenchmarkIntersect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Intersect(x, y)
	}
}

func BenchmarkIntersection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Intersection(x, y)
	}
}
