package main

import (
	"fmt"

	"github.com/ichsanputr/goschedtask"
)

func main() {
	goschedtask.RegisterJob(HelloWorld, goschedtask.Second(4))
	<-goschedtask.RunJobs()
}

func HelloWorld() {
	fmt.Println("Hello World")
}
