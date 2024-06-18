package utils

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestSingleFlightLruCache(t *testing.T) {
	cache := NewSingleFlightLruCache[string, string](5)

	t.Run("single flight", func(t *testing.T) {
		var counter int32
		var wg sync.WaitGroup
		result := make([]string, 10)
		for i := 0; i < 10; i++ {
			i := i
			wg.Add(1)
			go func() {
				defer wg.Done()

				val, _ := cache.Get(fmt.Sprintf("%v", i%3), func(key string) (string, error) {
					time.Sleep(time.Second)
					atomic.AddInt32(&counter, 1)
					return key, nil
				})
				result[i] = val
			}()
		}
		wg.Wait()
		if counter != 10/3 {
			t.Fatalf("single flight not working, %v", counter)
		}
		for i, item := range result {
			if item != fmt.Sprintf("%v", i%3) {
				t.Fatal("result not match")
			}
		}
	})

	t.Run("error on single flight", func(t *testing.T) {
		result := make([]error, 10)
		var wg sync.WaitGroup
		var counter int32
		for i := 0; i < 10; i++ {
			i := i
			wg.Add(1)
			go func() {
				defer wg.Done()
				_, err := cache.Get("111", func(key string) (string, error) {
					atomic.AddInt32(&counter, 1)
					time.Sleep(100 * time.Millisecond)
					return "", fmt.Errorf("my error")
				})
				result[i] = err
			}()
		}
		wg.Wait()
		if counter != 1 {
			t.Fatalf("counter unexpected: %v", counter)
		}
		for _, item := range result {
			if item.Error() != "my error" {
				t.Fatal("error unexpected")
			}
		}

		// check the error should not be cached
		_, err := cache.Get("111", func(key string) (string, error) {
			return "", fmt.Errorf("new error")
		})
		if err.Error() != "new error" {
			t.Fatal("result unexpected")
		}
	})

}
