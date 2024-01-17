package goschedtask

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

func TaskTimeout() {
	fmt.Println("Task FuncTimeout running")
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
	RegisterJobRunAt(TaskHttpReq, time.Now().Add(time.Second*10), "https://www.google.com")
	RegisterJob(TaskTimeout, Second(2))

	<-RunJobs()
}
