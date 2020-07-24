package problems

import (
	"github.com/hultan/euler/tools"
	"sync"
	"time"
)

type Euler001 struct {
	Sum    tools.SafeCounter
}

func NewEuler001() *Euler001 {
	return new(Euler001)
}

func (e *Euler001) Start() {
	defer tools.TimeTrack(time.Now(), "Euler_001 goroutine")
	var wg sync.WaitGroup
	e.Sum.Reset()

	for i := 1; i < 10000000; i++ {
		wg.Add(1)
		go func(i int) {
			if i%3 == 0 || i%5 == 0 {
				e.Sum.Inc(i)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func (e *Euler001) Start2() {
	defer tools.TimeTrack(time.Now(), "Euler_001 simple")
	e.Sum.Reset()
	for i := 1; i < 10000000; i++ {
		if i%3 == 0 || i%5 == 0 {
			e.Sum.Inc(i)
		}
	}
}
