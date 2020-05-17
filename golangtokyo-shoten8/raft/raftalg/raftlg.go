package raftalg

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/coreos/etcd/raft/raftpb"

	"github.com/coreos/etcd/wal/walpb"

	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/rafthttp"
	"github.com/coreos/etcd/wal"
)

type RaftAlg struct {
	commitC         chan string
	doneRestoreLogC chan struct{}

	node        raft.Node
	raftStorage *raft.MemoryStorage
	wal         *wal.WAL
	transport   *rafthttp.Transport

	id     int
	peers  []string
	waldir string
}

func New(id int, peers []string) *RaftAlg {
	return &RaftAlg{
		commitC:         make(chan string),
		doneRestoreLogC: make(chan struct{}),
		raftStorage:     raft.NewMemoryStorage(),
		id:              id,
		peers:           peers,
		waldir:          fmt.Sprintf("kvraft-%d", id),
	}
}

func (r *RaftAlg) Run(ctx context.Context) error {
	c := &raft.Config{
		ID:              uint64(r.id),
		ElectionTick:    10,
		HeartbeatTick:   1,
		Storage:         r.raftStorage,
		MaxSizePerMsg:   1024 * 1024,
		MaxInflightMsgs: 256,
		Logger: &raft.DefaultLogger{
			Logger: log.New(os.Stderr, "[Raft-debug]", 0),
		},
	}

	rpeers := make([]raft.Peer, len(r.peers))
	for i := range rpeers {
		rpeers[i] = raft.Peer{ID: uint64(i + 1)}
	}

	oldwal := wal.Exist(r.waldir)

	w, err := r.replayWAL(ctx)
	if err != nil {
		return err
	}
	r.wal = w

	if oldwal {
		r.node = raft.RestartNode(c)
	} else {
		r.node = raft.StartNode(c, rpeers)
	}

	return nil
}

func (r *RaftAlg) replayWAL(ctx context.Context) (*wal.WAL, error) {
	if !wal.Exist(r.waldir) {
		_ = os.Mkdir(r.waldir, 0750)
		w, _ := wal.Create(r.waldir, nil)
		w.Close()
	}
	w, _ := wal.Open(r.waldir, walpb.Snapshot{})
	_, _, ents, _ := w.ReadAll()
	select {
	case r.doneRestoreLogC <- struct{}{}:
	case <-time.After(10 * time.Second):
		return nil, errors.New(
			"timeout(10s) receiving done restore channel")
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	_ = r.publishEntries(ctx, ents)
	return w, nil
}

func (r *RaftAlg) publishEntries(ctx context.Context, ents []raftpb.Entry) error {
	for i := range ents {
		if ents[i].Type != raftpb.EntryNormal || len(ents[i].Data) == 0 {
			continue
		}

		s := ents[i].Data

		select {
		case r.commitC <- string(s):
		case <-time.After(10 * time.Second):
			return errors.New("timeout 10sec committed channel")
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}
