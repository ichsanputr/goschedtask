package goschedtask

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
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
	RegisterJobRunAt(TaskHttpReq, time.Now().Add(time.Second*2), "https://www.google.com")
	RegisterJob(TaskTimeout, Second(2), "Task FuncTimeout running", "kakas")
	RegisterJob(TaskTimeout, Second(2), "Task wkwk")

	<-RunJobs()
}
