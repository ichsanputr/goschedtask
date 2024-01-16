package goschedtask

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	RegisterJob(Job{
		JobFunc: func() {
			fmt.Println("Task first")
		},
		RunLoop:  true,
		Interval: Second(4),
	})

	RegisterJob(Job{
		JobFunc: func() {
			time.Sleep(2 * time.Second)
			fmt.Println("Task second")
		},
		RunLoop:  true,
		Interval: Second(8),
	})

	goschedtask := RunJobs()
	<-goschedtask
}
