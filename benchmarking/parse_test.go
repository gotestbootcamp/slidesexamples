package parse

import "testing"

func BenchmarkParse(b *testing.B) {
	b.Run("withReader", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ParseWithReader("testdata/basic.json")
		}
	})
	b.Run("withMarshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Parse("testdata/basic.json")
		}
	})
}
