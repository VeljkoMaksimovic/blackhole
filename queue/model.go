package queue

import (
	"context"
	"errors"
	"github.com/c12s/blackhole/model"
	pb "github.com/c12s/blackhole/pb"
	storage "github.com/c12s/blackhole/storage"
)

type TokenBucket struct {
	Capacity     int64
	Tokens       int64
	FillInterval *model.FillInterval
	Notify       chan bool
	Reset        chan bool
	Delay        chan bool
	Attempt      int
	TRetry       *model.Retry
}

type TaskQueue struct {
	Name   string
	Queue  storage.DB
	Bucket *TokenBucket
	Pool   *WorkerPool
}

type Worker struct {
	ID   string
	Kill chan bool
}

type WorkerPool struct {
	MaxQueued     int
	MaxWorkers    int
	Pipe          chan *pb.Task
	ActiveWorkers map[string]*Worker
	done          chan string
	active        chan string
	Workers       map[string]*Worker
}

type BlackHole struct {
	Queues map[string]*TaskQueue
}

func (bh *BlackHole) GetTK(name string) (*TaskQueue, error) {
	if tk, ok := bh.Queues[name]; ok {
		return tk, nil
	}
	return nil, errors.New("Queue not exists!")
}

func newPool(ctx context.Context, maxqueued, maxworkers int) *WorkerPool {
	return &WorkerPool{
		MaxQueued:     maxqueued,
		MaxWorkers:    maxworkers,
		Pipe:          make(chan *pb.Task),
		ActiveWorkers: map[string]*Worker{},
		done:          make(chan string),
		active:        make(chan string),
		Workers:       map[string]*Worker{},
	}
}

func newQueue(name string, tb *TokenBucket, wp *WorkerPool, db storage.DB) *TaskQueue {
	return &TaskQueue{
		Name:   name,
		Queue:  db,
		Bucket: tb,
		Pool:   wp,
	}
}

func newBucket(capacity, tokens int64, interval *model.FillInterval, retry *model.Retry) *TokenBucket {
	return &TokenBucket{
		Capacity:     capacity,
		Tokens:       tokens,
		FillInterval: interval,
		Notify:       make(chan bool),
		Delay:        make(chan bool),
		Reset:        make(chan bool),
		Attempt:      1,
		TRetry:       retry,
	}
}

func New(ctx context.Context, db storage.DB, options []*model.TaskOption) *BlackHole {
	q := map[string]*TaskQueue{}
	for _, opt := range options {
		tb := newBucket(opt.Capacity, opt.Tokens, opt.FillRate, opt.TRetry)
		wp := newPool(ctx, opt.MaxQueued, opt.MaxWorkers)
		tq := newQueue(opt.Name, tb, wp, db)

		q[opt.Name] = tq
		tq.StartQueue(ctx)
	}

	return &BlackHole{
		Queues: q,
	}
}
