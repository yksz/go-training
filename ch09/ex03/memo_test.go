package memo_test

import (
	"strconv"
	"testing"
	"time"

	"./memo"
)

func TestGet(t *testing.T) {
	processingTime := 100 * time.Millisecond
	processingTimeStr := strconv.FormatInt(processingTime.Nanoseconds(), 10)

	m := memo.New(do)
	defer m.Close()

	// not cached
	if elapsed := elapsedTime(func() { m.Get(processingTimeStr, nil) }); elapsed < processingTime {
		t.Errorf("not cached: elapsed=%s", elapsed)
	}
	// cached
	if elapsed := elapsedTime(func() { m.Get(processingTimeStr, nil) }); elapsed > processingTime {
		t.Errorf("cached: elapsed=%s", elapsed)
	}
}

func TestGetWithDone(t *testing.T) {
	processingTime := 100 * time.Millisecond
	processingTimeStr := strconv.FormatInt(processingTime.Nanoseconds(), 10)

	m := memo.New(do)
	defer m.Close()

	done := make(chan struct{})
	go func() {
		done <- struct{}{}
	}()

	// not cached
	if elapsed := elapsedTime(func() { m.Get(processingTimeStr, nil) }); elapsed < processingTime {
		t.Errorf("not cached: elapsed=%s", elapsed)
	}
	// done
	if elapsed := elapsedTime(func() { m.Get(processingTimeStr, done) }); elapsed > processingTime {
		t.Errorf("done: elapsed=%s", elapsed)
	}
	// not cached with done
	if elapsed := elapsedTime(func() { m.Get(processingTimeStr, nil) }); elapsed < processingTime {
		t.Errorf("not cached with done: elapsed=%s", elapsed)
	}
}

func do(nanoseconds string) (interface{}, error) {
	ns, err := strconv.Atoi(nanoseconds)
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Duration(ns) * time.Nanosecond)
	return ns, nil
}

func elapsedTime(f func()) time.Duration {
	start := time.Now()
	f()
	return time.Since(start)
}
