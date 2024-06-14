package anyerrgroup

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	g := New()

	g.Go(func(ctx context.Context) error {
		time.Sleep(time.Millisecond)
		t.Log("1")
		return nil
	})

	g.Go(func(ctx context.Context) error {
		time.Sleep(time.Millisecond * 2)
		t.Log("2")
		return nil
	})

	err := g.Wait()
	assert.NoError(t, err)
	t.Log("Wait() returned no error")
	//
	g = New()

	g.Go(func(ctx context.Context) error {
		time.Sleep(time.Millisecond)
		t.Log("1")
		return errors.New("1")
	})

	g.Go(func(ctx context.Context) error {
		time.Sleep(time.Millisecond * 2)
		t.Log("2")
		return errors.New("2")
	})

	err = g.Wait()
	assert.Error(t, err)
	//
	g = New()

	g.Go(func(ctx context.Context) error {
		time.Sleep(time.Millisecond)
		t.Log("1")
		return errors.New("1")
	})

	g.Go(func(ctx context.Context) error {
		time.Sleep(time.Millisecond * 2)
		t.Log("2")
		return nil
	})

	err = g.Wait()
	assert.Error(t, err)
	//
	g = New()

	g.Go(func(ctx context.Context) error {
		time.Sleep(time.Millisecond)
		t.Log("1")
		return nil
	})

	g.Go(func(ctx context.Context) error {
		time.Sleep(time.Millisecond * 2)
		t.Log("2")
		return errors.New("2")
	})

	err = g.Wait()
	assert.Error(t, err)

	//
	g = New()

	g.Go(func(ctx context.Context) error {
		time.Sleep(time.Millisecond)
		t.Log("1")
		return nil
	})

	g.Go(func(ctx context.Context) error {
		time.Sleep(time.Millisecond * 2)
		t.Log("2")
		return errors.New("2")
	})

	g.Go(func(ctx context.Context) error {
		time.Sleep(time.Millisecond * 3)
		t.Log("3")
		return errors.New("3")
	})

	g.Go(func(ctx context.Context) error {
		time.Sleep(time.Millisecond * 3)
		t.Log("4")
		return errors.New("4")
	})

	err = g.Wait()
	assert.Error(t, err)
	// ensure all goroutines are finished
	time.Sleep(time.Millisecond * 4)
}
