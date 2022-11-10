package wait_for_goroutine_finish

func WaitForGoroutineFinish(funcs ...func()) {
	count := len(funcs)
	wg := sync.WaitGroup{}
	wg.Add(count)

	for _, f := range funcs {
		fc := f
		go func() {
			fc()
			wg.Done()
		}()
	}

	wg.Wait()
}
