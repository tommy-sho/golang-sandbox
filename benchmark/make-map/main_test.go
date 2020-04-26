package main

import "testing"

var (
	one      = createInput(1)
	ten      = createInput(10)
	hundred  = createInput(100)
	thousand = createInput(1000)
)

func BenchmarkNoMake(b *testing.B) {
	b.ReportAllocs()
	mapWithNoMake(one)
}
func BenchmarkNoMake10(b *testing.B) {
	b.ReportAllocs()
	mapWithNoMake(ten)
}
func BenchmarkNoMake100(b *testing.B) {
	b.ReportAllocs()
	mapWithNoMake(hundred)
}
func BenchmarkNoMake1000(b *testing.B) {
	b.ReportAllocs()
	mapWithNoMake(thousand)
}
func BenchmarkMake(b *testing.B) {
	b.ReportAllocs()
	mapWithMake(one)
}
func BenchmarkMake10(b *testing.B) {
	b.ReportAllocs()
	mapWithMake(ten)
}
func BenchmarkMake100(b *testing.B) {
	b.ReportAllocs()
	mapWithMake(hundred)
}
func BenchmarkMake1000(b *testing.B) {
	b.ReportAllocs()
	mapWithMake(thousand)
}

func TestMake(t *testing.T) {
	mapWithNoMake(hundred)
}
func createInput(in int) []int {
	s := make([]int, in)
	for i := 0; i < in; i++ {
		s[i] = i
	}

	return s
}
