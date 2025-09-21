package sequence

func Sequence(args ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)

		var start, stop, step int
		switch len(args) {
		case 1:
			start, stop, step = 0, args[0], 1
		case 2:
			start, stop, step = args[0], args[1], 1
		case 3:
			start, stop, step = args[0], args[1], args[2]
		default:
			return
		}

		if step <= 0 && start <= stop {
			return
		}

		for i := start; i < stop; i += step {
			ch <- i
		}
	}()
	return ch
}
