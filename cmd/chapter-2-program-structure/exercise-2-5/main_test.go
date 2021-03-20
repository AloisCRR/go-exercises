package pop_count_v3

import (
	popcountv1 "github.com/AloisCRR/go-exercises/cmd/chapter-2-program-structure/exercise-2-3"
	popcountv2 "github.com/AloisCRR/go-exercises/cmd/chapter-2-program-structure/exercise-2-4"
	"testing"
)

func BenchmarkPopCountTableLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcountv1.PopCountTableLookup(12829474)
	}
}

func BenchmarkPopCountClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClear(12829474)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcountv1.PopCountLoop(12829474)
	}
}

func BenchmarkPopCountBitShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcountv2.PopCountBitShift(12829474)
	}
}
