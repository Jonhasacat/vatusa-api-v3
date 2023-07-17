package async

import "sync"

type WorkerFunc func(offset int)

func SpawnWorkers(numWorkers int, workerFunc WorkerFunc) {
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(offset int) {
			workerFunc(offset)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}
