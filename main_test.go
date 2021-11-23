package main

import (
	"runtime"
	"testing"
)

func BenchmarkCaller(b *testing.B) {
	for i:=0;i<b.N;i++{
		Caller(1)
	}
}

func BenchmarkRuntime(b *testing.B) {
	for i:=0;i<b.N;i++{
		runtime.Caller(1)
	}
}
