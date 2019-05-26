package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func captureOutput(f func() error) (string, error) {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	os.Stdout = w

	funcErr := f()

	os.Stdout = stdout
	err = w.Close()
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		panic(err)
	}

	return buf.String(), funcErr
}

func TestForEach(t *testing.T) {
	testFunc := func(s string) error {
		_, err := fmt.Print(s + "1")
		if err != nil {
			return err
		}
		return nil
	}

	expect := []struct {
		name           string
		list           Seq
		f              func(string) error
		expectedResult string
		expectedErr    error
	}{
		{
			name:           "Failed",
			list:           []string{""},
			f:              testFunc,
			expectedResult: "1",
			expectedErr:    nil,
		},
		{
			name:           "success",
			list:           []string{"1", "2", "3", "4", "5"},
			f:              testFunc,
			expectedResult: "1121314151",
			expectedErr:    nil,
		},
	}
	for _, c := range expect {
		output, err := captureOutput(func() error {
			return c.list.ForEach(c.f)
		})
		assert.Equal(t, c.expectedResult, output)
		assert.Equal(t, c.expectedErr, err)
	}
}
