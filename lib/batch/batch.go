package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	mu := sync.Mutex{}
	wg := new(sync.WaitGroup)
	res = make([]user, 0, n)

	ch := make(chan struct{}, pool)

	i := int64(0)

	for i = 0; i < n; i++ {
		wg.Add(1)
		ch <- struct{}{}

		go func(any int64) {
			user := getOne(any)
			mu.Lock()
			res = append(res, user)
			mu.Unlock()
			<-ch
			defer wg.Done()
		}(i)
	}

	wg.Wait()

	return
}
