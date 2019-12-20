package main

import (
	"testing"

)

func BenchmarkGoParalel(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		parallel2()
		
	}
}


func BenchmarkForLoop(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		forloop()
		
	}
}
