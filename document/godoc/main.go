package main

import "fmt"

// Repository is an interface for repository
type Repository interface {
	Get(string) string
	Set(string, string) string
}

var db map[string]string

// sub is data structure
type sub struct {
	data map[string]string
}

// NewRepository return Repositroy implementing object
func NewRepository() Repository {
	db := make(map[string]string)
	return &sub{db}
}

func (s sub) Get(in string) string {
	return s.data[in]
}

func (s sub) Set(key, in string) string {
	s.data[key] = in
	return in
}

// main in main function
func main() {
	db := NewRepository()
	db.Set("1", "one")
	fmt.Println(db.Get("1"))
}
