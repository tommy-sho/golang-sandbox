package main

import (
	"testing"
)

func benchmarkList(len int, t *testing.B) {

	list := makeList(len)
	t.ResetTimer()
	// Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
	for i := 0; i < t.N; i++ {
		for _, i := range list {
			if i == 10 {
				_ = i
			}
		}
	}
}

func benchmarkMap(len int, t *testing.B) {
	list := makeMap(len)
	t.ResetTimer()
	// Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
	for i := 0; i < t.N; i++ {
		for v, k := range list {
			if v == 10 {
				_ = k
			}
		}
	}
}

func BenchmarkList1(b *testing.B)   { benchmarkList(1, b) }
func BenchmarkList5(b *testing.B)   { benchmarkList(5, b) }
func BenchmarkList10(b *testing.B)  { benchmarkList(10, b) }
func BenchmarkList100(b *testing.B) { benchmarkList(100, b) }
func BenchmarkMap1(b *testing.B)    { benchmarkMap(1, b) }
func BenchmarkMap5(b *testing.B)    { benchmarkMap(5, b) }
func BenchmarkMap10(b *testing.B)   { benchmarkMap(10, b) }
func BenchmarkMap100(b *testing.B)  { benchmarkMap(100, b) }
