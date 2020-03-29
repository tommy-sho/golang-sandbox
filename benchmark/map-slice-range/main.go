package main

func makeList(len int) []int {
	res := make([]int, 0, len)
	for i := 0; i < len; i++ {
		res = append(res, i)
	}

	return res
}

func makeMap(len int) map[int]int {
	res := make(map[int]int, len)
	for i := 0; i < len; i++ {
		res[i] = i
	}

	return res
}
