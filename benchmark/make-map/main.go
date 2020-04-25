package main

func main() {}

func mapWithMake(in []int) {
	s := make(map[int]int, len(in))
	for n := range in {
		s[in[n]] = in[n]
	}
}

func mapWithNoMake(in []int) {
	s := map[int]int{}
	for n := range in {
		s[in[n]] = in[n]
	}
}
