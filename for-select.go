package main

// the FOR-SELECT

// This pattern is typically used to read data from multiple channels.

var c1, c2 <-chan int

func main() {

	// Either loop infinitely or range over something
	for {
		// All cases are considered simultaneously and have equal chance to be selected.
		// If none of the cases are ready to run, the entire select statement blocks.
		select {
		case <-c1:
			// do some work with int sent to channel c1
		case <-c2:
			// do some work with int sent to channel c2
		default:
			// auto run if other cases are not ready
		}
		// do some further work
	}

}
