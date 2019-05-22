package main

func (s Seq) Map(f func(string) string) Seq {
	var res = make([]string, len(s))
	for i, r := range s {
		res[i] = f(r)
	}
	return res
}
