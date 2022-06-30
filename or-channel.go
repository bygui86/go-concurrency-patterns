package main

// The OR channel

// This pattern aims to combine multiple "done" channels into one "agg_done".
// If one of the "done" channels is signaled, then the whole "agg_done" channel is also closed.
// Useful when we do not know in advanced how many "done" channels we are going to have at runtime.

func main() {

	// return agg_done channel
	var or func(channels ... <-chan interface{}) <- chan interface{}

	or = func(channels ...<-chan interface{}) <-chan interface{} {
		// base cases
		switch len(channels) {
			case 0: return nil
			case 1: return channels[0]
		}

		orDone := make(chan interface{})

		go func() {
			defer close(orDone)

			switch len(channels) {
				case 2: 
					select {
						case <- channels[0]:
						case <- channels[1]:
					}
				default:
					select {
						case <- channels[0]:
						case <- channels[1]:
						case <- channels[2]:
						// Following line makes the upper & lower recursive function depends on each other like a tree.
						// The upper injects its own "orDone" channel into the lower.
						// Then the lower also return its own "orDone" to the upper.
						// If any orDone channel closes, the upper & lower both are notified.
						case <- or(append(channels[3:], orDone)...):
					}

			}

		}
		return orDone
	}

}
