package main

// the DONE channel

// This pattern is utilized a lot in other patterns

// Goroutine is not garbage collected; hence, it is likely to be leaked.
// To avoid leaking, Goroutine should be cancelled whenever it is told to do.
// A parent Goroutine needs to send cancellation signal to its child via a read-only channel named "done".
// By convention, it is set as the 1st parameter.

func main() {

	// child goroutine
	doWork(<-done chan interface {}, other_params) <- terminated chan interface{} {
		terminated := make(chan interface{}) // to tell outer that it has finished
		defer close(terminated)

		for {
			select: {
			case:
				// do some work
			case <- done:
				return
			}
			// do some work
		}

		return terminated
	}

	// parent goroutine
	done := make(chan interface{})
	terminated := doWork(done, other_args)

	// do some work

	// tell child to stop
	close (done)

	// wait for child finish its work
	<- terminated

}
