package pop_count_v2

import (
	pop_count_v1 "github.com/AloisCRR/go-exercises/cmd/chapter-2-program-structure/exercise-2-3"
	"testing"
)

func BenchmarkPopCountBitShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountBitShift(83292348397)
	}
}

func BenchmarkPopCountTableLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pop_count_v1.PopCountTableLookup(83292348397)
	}
}
