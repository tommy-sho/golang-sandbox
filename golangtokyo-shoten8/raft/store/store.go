package store

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"sync"
	"time"
)

type Store struct {
	mu      sync.RWMutex
	kvStore map[string]string
	Raft
}

func New(raft Raft) *Store {
	return &Store{
		kvStore: make(map[string]string),
		Raft:    raft,
	}
}

func (s *Store) Lookup(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	v, ok := s.kvStore[key]
	return v, ok
}

func (s *Store) Save(key, value string) error {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(kv{key, value})
	if err != nil {
		return err
	}

	if err = s.Propose(buf.Bytes()); err != nil {
		return nil
	}

	return nil
}

func (s *Store) RunCommitReader(ctx context.Context) error {
	select {
	case <-s.Raft.DoneReplayWAL():
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(10 * time.Second):
		return errors.New("timeout 10sec receiving done replay channel")
	}

	for {
		select {
		case data := <-s.Raft.Commit():
			var kvdata kv
			dec := gob.NewDecoder(bytes.NewBuffer([]byte(data)))
			if err := dec.Decode(&kvdata); err != nil {
				return err
			}

			s.mu.Lock()
			s.kvStore[kvdata.Key] = kvdata.Value
			s.mu.Unlock()

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

type Raft interface {
	Propose(prop []byte) error
	Commit() <-chan string
	DoneReplayWAL() <-chan struct{}
}

type kv struct {
	Key   string
	Value string
}
