package event

import (
	"sync"

	"github.com/iqrahadian/paperid-assesment/model/param"
)

var (
	jobChan     = make(chan param.DisburseParam)
	Mutex       = &sync.RWMutex{}
	AccountLock = make(map[string]bool)
)

func PublishEvent(job param.DisburseParam) {
	jobChan <- job
}

func StartConsumer(wg *sync.WaitGroup) {
	for w := 1; w < 10; w++ {
		wg.Add(1)
		go runConsumer(wg)
	}
}

func runConsumer(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case job := <-jobChan:
			if LockAccount(job.SourceAccountID) {
				executeDisburse(job)
			} else {
				jobChan <- job
			}
		}
	}
}

func LockAccount(key string) bool {
	defer Mutex.Unlock()
	Mutex.Lock()

	if _, ok := AccountLock[key]; ok {
		return false
	}

	AccountLock[key] = true

	return true
}

func UnlockAccount(v string) {
	Mutex.Lock()
	delete(AccountLock, v)
	Mutex.Unlock()
}
