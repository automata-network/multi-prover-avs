package utils

import (
	"sync"

	"github.com/chzyer/logex"
	lru "github.com/hashicorp/golang-lru/v2"
)

type SingleFlightLruCache[K comparable, V any] struct {
	cache        *lru.Cache[K, V]
	singleFlight map[K]*singleFlightProc[V]
	mutex        sync.Mutex
}

type singleFlightProc[V any] struct {
	val    V
	err    error
	mutex  sync.Mutex
	notify chan struct{}
}

func NewSingleFlightLruCache[K comparable, V any](cap int) *SingleFlightLruCache[K, V] {
	cache, err := lru.New[K, V](cap)
	if err != nil {
		panic(err)
	}
	return &SingleFlightLruCache[K, V]{
		cache:        cache,
		singleFlight: make(map[K]*singleFlightProc[V], 8),
	}
}

func (s *SingleFlightLruCache[K, V]) Get(key K, getter func(key K) (V, error)) (V, error) {
	var err error
	val, ok := s.cache.Get(key)
	if ok {
		return val, nil
	}
	s.mutex.Lock()
	val, ok = s.cache.Get(key)
	if ok {
		return val, nil
	}
	proc, ok := s.singleFlight[key]
	if !ok {
		proc = &singleFlightProc[V]{notify: make(chan struct{})}
		s.singleFlight[key] = proc
	}
	s.mutex.Unlock()

	if !ok {
		val, err = getter(key)

		s.mutex.Lock()
		if err == nil {
			s.cache.Add(key, val)
		}
		delete(s.singleFlight, key)
		s.mutex.Unlock()

		proc.mutex.Lock()
		proc.val = val
		proc.err = err
		close(proc.notify)
		proc.mutex.Unlock()
	} else {
		// waiting for the single flight finished
		<-proc.notify

		proc.mutex.Lock()
		val, err = proc.val, proc.err
		proc.mutex.Unlock()
	}
	if err != nil {
		return val, logex.Trace(err)
	}
	return val, nil
}
