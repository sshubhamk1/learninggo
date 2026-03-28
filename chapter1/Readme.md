# Chapter 1 findings

- Go scheduler: Orchestrates concurrent functions called goroutines.
- Garbage collector: Runs in the background and automatically reclaims memory that is guaranteed to no longer be in use.
- Tony Hoare’s 1978 article “Communicating Sequential Processes.”
- The go keyword schedules the sequential function to concurrent function.
- **In traditional programming languages, thread pools are often used to manage the high execution cost of running multiple threads. In Go, the story is different. Goroutines are incredibly lightweight, eliminating the need for a thread pool. This efficiency allows you to run millions of goroutines on a regular machine, opening a world of possibilities.**
- go routines switches in tens of nanoseconds while kernel threads does it in microseconds.
- Go routines size start around 2KB while traditional threads starts with 8KB.
- goroutines are inside runqueues(queue with heaps)
- If local runqueue is empty steal from the other runqueue
- It also have a global runqueue (lower priority queue less check on it)
- It uses composition instead of inheritance.