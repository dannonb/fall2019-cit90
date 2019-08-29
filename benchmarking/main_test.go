package main



func BenchmarkCompLit(b *testing.B) {
	for i := 0, i < b.N; i++ {
		compLit()
	}
}