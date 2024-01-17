package goschedtask

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TaskTimeout() {
	fmt.Println("Task FuncTimeout running")
	time.Sleep(3 * time.Second)
	fmt.Println("Task Timeout finish")
}

func TaskHttpReq(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Task HttpReq with %d \n", resp.StatusCode)
}

func Test(t *testing.T) {
	RegisterJob(TaskHttpReq, Second(5), "https://www.adsdgoogle.com")

	goschedtask := RunJobs()
	<-goschedtask
}
