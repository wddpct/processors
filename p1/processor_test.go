package p1

import (
	"testing"
)

func BenchmarkProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Process()
	}
}
