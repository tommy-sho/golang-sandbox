package main

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
)

//Calculate until 10.
func Calculate(x, y uint8) (uint8, error) {
	res := x + y
	if 10 > res {
		return res, error(nil)
	}

	return 0, errors.New("over 10")
}

func RandCalculate() (uint8, error) {
	x := uint8(rand.Int63n(10))
	y := uint8(rand.Int63n(10))
	fmt.Println(fmt.Sprintf("num is %d, %d", x, y))
	return Calculate(x, 10)
}

func main() {
	res, err := RandCalculate()
	fmt.Println(reflect.TypeOf(err))
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
