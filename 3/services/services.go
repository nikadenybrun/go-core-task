package services

import "sync"

type Map struct {
	data map[string]int
	mu   sync.RWMutex
}

func NewMap() *Map {
	return &Map{
		data: make(map[string]int),
	}
}

func (m *Map) Add(key string, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *Map) Remove(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}

func (m *Map) Copy() map[string]int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	copiedMap := make(map[string]int)
	for k, v := range m.data {
		copiedMap[k] = v
	}
	return copiedMap
}

func (m *Map) Exists(key string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	_, Ok := m.data[key]
	return Ok
}

func (m *Map) Get(key string) (int, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, Ok := m.data[key]
	return value, Ok
}
