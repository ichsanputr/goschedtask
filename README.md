## Goschedtask: A Golang package for scheduling Task.

Goschedtask is a Golang job scheduling package which lets you run go functions periodically based on intervals or running at a certain time. This package is inspired by GoCron version 1.

Install it

```
go get github.com/ichsanputr/goschedtask
```
Basic usage

```go
package main

import (
	"fmt"
	"github.com/ichsanputr/goschedtask"
)

func TaskTimeout(msg string) {
	fmt.Println(msg)
}

func TaskHttpReq(url string) {
	resp, err := http.Get(url)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("Task HttpReq with %d \n", resp.StatusCode)
}

func Test(t *testing.T) {
  // run job on certain time
  goschedtask.RegisterJobRunAt(TaskHttpReq, time.Now().Add(time.Second*2), "https://www.google.com")

  // run job periodic based on interval
  goschedtask.RegisterJob(TaskTimeout, Second(2), "Task FuncTimeout running")

  // run job once
  goschedtask.RegisterJobRunOnce(TaskTimeout, Second(4), "Task Hai")

  // run all jobs
  <-goschedtask.RunJobs()
}
```

This code will block the code from running because it is waiting for the value from channel stopped. To stop the scheduler from running you can enter the value to the channel scheduler in a goroutine.

```go
scheduler := goschedtask.RunJobs()

go func() {
  time.Sleep(4 * time.Second)

  // stopping scheduler after 4 second
  scheduler <- true
}()

<-scheduler
```

You can also run other code while the scheduler is running, like this.

```go
scheduler := goschedtask.RunJobs()
SomeFunc()
<-scheduler
```


