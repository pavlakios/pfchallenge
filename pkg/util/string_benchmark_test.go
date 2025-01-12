package util

import "testing"

func BenchmarkRandString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RandString(10)
	}
}

func BenchmarkRandStringHex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RandStringHex(10)
	}
}
