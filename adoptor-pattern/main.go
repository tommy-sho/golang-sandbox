package main

import "fmt"

type Option string

type Handler struct {
	URL     string
	Owner   string
	Options []Option
}

type CustomOptionFunc func(*Handler) *Handler

func CustomOption(h *Handler, options ...CustomOptionFunc) {
	for _, o := range options {
		h = o(h)
	}
}

func AddOption(s string) CustomOptionFunc {
	return CustomOptionFunc(func(h *Handler) *Handler {
		option := Option(s)
		h.Options = append(h.Options, option)
		return h
	})
}

func main() {
	handler := &Handler{
		URL:   "google.com",
		Owner: "tomioka",
	}

	fmt.Println(handler)

	firstOption := AddOption("first")
	secondOption := AddOption("second")
	CustomOption(handler, firstOption, secondOption)
	fmt.Println(handler)
}
