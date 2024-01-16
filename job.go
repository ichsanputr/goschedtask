package goschedtask

import "time"

type Job struct {
	JobFunc  interface{}
	RunLoop  bool
	Interval time.Duration
	TimeRun  time.Time
}

var (
	Jobs = []Job{}
)

func RegisterJob(job Job) {
	Jobs = append(Jobs, job)
}
