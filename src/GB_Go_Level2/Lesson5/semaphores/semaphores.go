// Package semaphores
// Task#3
// Протестируйте производительность операций чтения и записи на множестве действительных чисел,
// безопасность которого обеспечивается sync.Mutex и sync.RWMutex
// для разных вариантов использования: 10% запись, 90% чтение; 50% запись, 50% чтение; 90% запись, 10% чтение
package semaphores

import "sync"

// multiplicity struct for numbers
type multiplicity struct {
	mMap map[int]float64
	mu   sync.Mutex
	muRW sync.RWMutex
}

// NewMultiplicity init new multiplicity struct
func NewMultiplicity() *multiplicity {
	return &multiplicity{
		mMap: map[int]float64{},
	}
}

// Write to multiplicity with Mutex
func (m *multiplicity) Write(i int, f float64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mMap[i] = f
}

// WriteRW to multiplicity with RWMutex
func (m *multiplicity) WriteRW(i int, f float64) {
	m.muRW.Lock()
	defer m.muRW.Unlock()
	m.mMap[i] = f
}

// Read to multiplicity with Mutex
func (m *multiplicity) Read(i int) float64 {
	m.mu.Lock()
	defer m.mu.Unlock()
	res, ok := m.mMap[i]
	if !ok {
		return 0
	}
	return res
}

// ReadRW to multiplicity with RWMutex
func (m *multiplicity) ReadRW(i int) float64 {
	m.muRW.RLock()
	defer m.muRW.RUnlock()
	res, ok := m.mMap[i]
	if !ok {
		return 0
	}
	return res
}
