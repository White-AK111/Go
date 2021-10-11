package semaphores

import (
	"math/rand"
	"testing"
	"time"
)

// Benchmark for 10% write and 90% read, using Mutex
func BenchmarkMultiplicity_W10_R90(b *testing.B) {
	var m = NewMultiplicity()
	rand.Seed(time.Now().UnixNano())

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 0; i < 1000; i++ {
					m.Write(i, rand.Float64())
				}
				for i := 0; i < 9000; i++ {
					_ = m.Read(rand.Intn(1000))
				}
			}
		})
	})
}

// Benchmark for 50% write and 50% read, using Mutex
func BenchmarkMultiplicity_W50_R50(b *testing.B) {
	var m = NewMultiplicity()
	rand.Seed(time.Now().UnixNano())

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 0; i < 5000; i++ {
					m.Write(i, rand.Float64())
				}
				for i := 0; i < 5000; i++ {
					_ = m.Read(rand.Intn(5000))
				}
			}
		})
	})
}

// Benchmark for 90% write and 10% read, using Mutex
func BenchmarkMultiplicity_W90_R10(b *testing.B) {
	var m = NewMultiplicity()
	rand.Seed(time.Now().UnixNano())

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 0; i < 9000; i++ {
					m.Write(i, rand.Float64())
				}
				for i := 0; i < 1000; i++ {
					_ = m.Read(rand.Intn(9000))
				}
			}
		})
	})
}

// Benchmark for 10% write and 90% read, using RWMutex
func BenchmarkMultiplicity_W10_R90_RW(b *testing.B) {
	var m = NewMultiplicity()
	rand.Seed(time.Now().UnixNano())

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 0; i < 1000; i++ {
					m.Write(i, rand.Float64())
				}
				for i := 0; i < 9000; i++ {
					_ = m.Read(rand.Intn(1000))
				}
			}
		})
	})
}

// Benchmark for 50% write and 50% read, using RWMutex
func BenchmarkMultiplicity_W50_R50_RW(b *testing.B) {
	var m = NewMultiplicity()
	rand.Seed(time.Now().UnixNano())

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 0; i < 5000; i++ {
					m.Write(i, rand.Float64())
				}
				for i := 0; i < 5000; i++ {
					_ = m.Read(rand.Intn(5000))
				}
			}
		})
	})
}

// Benchmark for 90% write and 10% read, using RWMutex
func BenchmarkMultiplicity_W90_R10_RW(b *testing.B) {
	var m = NewMultiplicity()
	rand.Seed(time.Now().UnixNano())

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 0; i < 9000; i++ {
					m.Write(i, rand.Float64())
				}
				for i := 0; i < 1000; i++ {
					_ = m.Read(rand.Intn(9000))
				}
			}
		})
	})
}
