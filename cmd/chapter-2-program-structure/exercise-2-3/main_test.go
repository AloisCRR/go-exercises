package pop_count_v1

import "testing"

func BenchmarkPopCountTableLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountTableLookup(10293458583234)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(10293458583234)
	}
}
