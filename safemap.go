package safemap

import (
	"sync"
)

type SafeMap struct {
	lock *sync.RWMutex
	data map[interface{}]interface{}
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		lock: new(sync.RWMutex),
		data: make(map[interface{}]interface{})}
}

func (m *SafeMap) Set(k interface{}, v interface{}) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	if val, ok := m.data[k]; !ok {
		m.data[k] = v
	} else if val != v {
		m.data[k] = v
	} else {
		return false
	}
	return true
}

func (m *SafeMap) Get(k interface{}) interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if val, ok := m.data[k]; ok {
		return val
	}
	return nil
}

func (m *SafeMap) Each(f func(k interface{}, v interface{})) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	for key, val := range m.data {
		f(key, val)
	}
}

func (m *SafeMap) Check(k interface{}) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if _, ok := m.data[k]; ok {
		return true
	}
	return false
}

func (m *SafeMap) Delete(k interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.data, k)
}
