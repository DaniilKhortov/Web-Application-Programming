package queueapp

import "testing"

func TestDummy(t *testing.T) {
	t.Log("test visible")
}
func BenchmarkSimulateQueueConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimulateQueueConcat(1000)
	}
}

func BenchmarkSimulateQueueConcatOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimulateQueueConcatOptimized(1000)
	}
}
