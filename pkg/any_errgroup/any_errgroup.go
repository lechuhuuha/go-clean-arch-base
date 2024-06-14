package anyerrgroup

import (
	"context"
	"sync"
	"sync/atomic"
)

type ErrGroup struct {
	errOnce sync.Once

	ch    chan error
	total int32

	ctx    context.Context
	cancel func()
}

func New() *ErrGroup {
	ctx, cancel := context.WithCancel(context.Background())
	return &ErrGroup{
		ctx:    ctx,
		cancel: cancel,
		ch:     make(chan error),
	}
}

func (e *ErrGroup) Go(f func(ctx context.Context) error) {
	atomic.AddInt32(&e.total, 1)
	go func() {
		defer atomic.AddInt32(&e.total, -1)
		if err := f(e.ctx); err != nil {
			e.errOnce.Do(func() {
				if e.cancel != nil {
					e.cancel()
				}
				e.ch <- err
			})
		} else if atomic.LoadInt32(&e.total) == 1 {
			e.errOnce.Do(func() {
				if e.cancel != nil {
					e.cancel()
				}
				e.ch <- nil
			})
		}
	}()
}

func (e *ErrGroup) Wait() error {
	return <-e.ch
}
