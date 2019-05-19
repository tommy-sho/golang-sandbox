package main

type Seq []string

func (s Seq) ForEach(f func(string) string) Seq {
	var res = make([]string, len(s))
	for i, r := range s {
		res[i] = f(r)
	}
	return res
}
