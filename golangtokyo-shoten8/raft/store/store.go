package store

import "sync"

type Store struct {
	mu      sync.Mutex
	kvStore map[string]string
}

func New() *Store {
	return &Store{
		kvStore: make(map[string]string),
	}
}

func (s *Store) Lookup(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	v, ok := s.kvStore[key]
	return v, ok
}

func (s *Store) Save(key, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.kvStore[key] = value
	return nil
}
