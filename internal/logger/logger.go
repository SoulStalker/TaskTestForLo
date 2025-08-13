package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Event struct {
	Time   time.Time   `json:"time"`
	Action string      `json:"action"`
	Data   any `json:"data,omitempty"`
}

type AsyncLogger struct {
	ch     chan Event
	wg     sync.WaitGroup
	out    io.Writer
	closed chan struct{}
}

func NewAsyncLogger(buffer int, out io.Writer) *AsyncLogger {
	if out == nil {
		out = os.Stdout
	}
	l := &AsyncLogger{
		ch:     make(chan Event, buffer),
		out:    out,
		closed: make(chan struct{}),
	}
	l.wg.Add(1)
	go l.loop()
	return l
}

func (l *AsyncLogger) loop() {
	defer l.wg.Done()
	enc := json.NewEncoder(l.out)
	for ev := range l.ch {
		_ = enc.Encode(ev) // fire-and-forget
	}
	close(l.closed)
}

func (l *AsyncLogger) Log(action string, data any) {
	select {
	case l.ch <- Event{Time: time.Now().UTC(), Action: action, Data: data}:
	default:
		// drop on full channel to keep API responsive
		fmt.Fprintln(l.out, `{"time":"`+time.Now().UTC().Format(time.RFC3339Nano)+`","action":"log_drop"}`)
	}
}

func (l *AsyncLogger) Shutdown(ctx context.Context) error {
	close(l.ch)
	done := make(chan struct{})
	go func() { l.wg.Wait(); close(done) }()
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
