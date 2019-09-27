package sub

// Copyright xxx

// +build darwin

/*
このパッケージはGodocのサンプル用のパッケージどすえ

How to use this package

このパッケージはただのテストなのでなんの役にも立ちません
*/

// Repository is an interface for repository
type Repository interface {
	Get(string) string         // これはGet
	Set(string, string) string // これはSet
}

// DB is Data box
var DB map[string]string

// Sub is Data structure
type Sub struct {
	Data map[string]string
}

// NewRepository return Repositroy implementing object
func NewRepository() Repository {
	db := make(map[string]string)
	return &Sub{db}
}

// Get はただただデータを返します
func (s Sub) Get(in string) string {
	return s.Data[in]
}

// Set はただただデータをセットします
func (s Sub) Set(key, in string) string {
	s.Data[key] = in
	return in
}

// Init is init function
//
// Deprecated: This function is deprecated
func Init() {
	//TODO: 実装する
}
