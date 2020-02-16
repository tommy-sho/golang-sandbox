package main

import (
	"errors"
	"fmt"

	"golang.org/x/xerrors"
)

var (
	HogeError error = errors.New("hogehoge error")
	XerrorErr error = xerrors.New("xerror")
)

func main() {
	e := errors.New("first error")
	wrapErr := fmt.Errorf("faield to hogehoge: %w", e)
	fmt.Printf("%+v\n", wrapErr)

	errs := fmt.Errorf("this is error: %v", HogeError)
	fmt.Printf("%+v\n", errs)

	werr := xerrors.Errorf("failed to hogehoge: %w", XerrorErr)
	fmt.Printf("%#v\n", werr)

}
