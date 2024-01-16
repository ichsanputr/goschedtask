package goschedtask

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	RegisterJob(func() {
		fmt.Println("Task third")
	}, Second(3))

	RegisterJobRunOnce(func() {
		fmt.Println("Task first")
	}, Second(3))

	RegisterJobRunOnce(func() {
		fmt.Println("Task second")
	}, Second(10))

	goschedtask := RunJobs()
	<-goschedtask
}
