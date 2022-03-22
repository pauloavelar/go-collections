package collections

import (
	"encoding/json"
	"sync"
)

type Set[T comparable] struct {
	mutex sync.RWMutex
	data  map[T]struct{}
}

func NewSet[T comparable](values ...T) *Set[T] {
	res := &Set[T]{data: make(map[T]struct{})}

	for _, each := range values {
		res.data[each] = struct{}{}
	}

	return res
}

func (s *Set[T]) Add(value ...T) *Set[T] {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.data == nil {
		s.data = make(map[T]struct{})
	}

	for _, each := range value {
		s.data[each] = struct{}{}
	}

	return s
}

func (s *Set[T]) Remove(value T) *Set[T] {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.data != nil {
		delete(s.data, value)
	}

	return s
}

func (s *Set[T]) Contains(value T) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var exists bool
	if s.data != nil {
		_, exists = s.data[value]
	}

	return exists
}

func (s Set[T]) MarshalJSON() ([]byte, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if len(s.data) == 0 {
		return nil, nil
	}

	res := make([]T, 0, len(s.data))
	for item := range s.data {
		res = append(res, item)
	}

	return json.Marshal(res)
}

func (s *Set[T]) UnmarshalJSON(data []byte) error {
	var items []T

	err := json.Unmarshal(data, &items)
	if err != nil {
		return err
	}

	res := make(map[T]struct{}, len(items))
	for _, item := range items {
		res[item] = struct{}{}
	}

	*s = Set[T]{data: res}

	return nil
}
