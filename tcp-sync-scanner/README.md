Port scanning using a worker pool is an effective way to manage concurrent network requests without exhausting system resources. The pattern utilizes goroutines and channels to distribute scanning tasks among a fixed number of workers. 

`main.go` is a worker pool pattern usinga buffered channel + a Waitgroup. 