Port scanning using a worker pool is an effective way to manage concurrent network requests without exhausting system resources. The pattern utilizes goroutines and channels to distribute scanning tasks among a fixed number of workers. 

`main.go` is a worker pool pattern usinga buffered channel + a Waitgroup. 

### The process:
- We create a job queue (`ports` channel);
- You start a fixed number of workers (100 goroutines);
- You queue up 1024 jobs (numbers 1-1024);
- Each worker pulls jobs, prints them, and calls `wg.Done()`;
- `main` waits until all jobs are done;
- So worker pool pattern is about having controlled concurrency instead of 1024 goroutines at once.