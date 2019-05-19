package main

import (
	"fmt"
	"testing"
)

func TestForEach(t *testing.T) {
	list := Seq([]string{"1", "2", "3", "4"})
	f := func(s string) string {
		return s + "s"
	}
	fmt.Println(list.ForEach(f))
	expect := []struct {
		name           string
		list           Seq
		f              func(string) string
		expectedResult Seq
	}{
		{
			name: "Eailed",
			list: []string{"1", "2", "3", "4", "5"},
			f: func(s string) string {
				return s + "1"
			},
			expectedResult: []string{"11", "21", "31", "41", "51"},
		},
		{
			name: "success",
			list: []string{"1", "2", "3", "4", "5"},
			f: func(s string) string {
				return "1" + s
			},
			expectedResult: []string{"11", "12", "13", "14", "15"},
		},
	}
	for _, c := range expect {
		res := c.list.ForEach(c.f)
		for i := range res {
			if res[i] != c.expectedResult[i] {
				t.Errorf("failed")
			}
		}
	}
}
