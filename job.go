package goschedtask

import "time"

type Job struct {
	JobFunc     interface{}
	JobParams   []interface{}
	RunLoop     bool
	Interval    time.Duration
	TimeRun     time.Time
	MustDeleted bool
}

var (
	Jobs = []Job{}
)

func RegisterJob(jobFunc interface{}, interval time.Duration, jobParams ...interface{}) {
	job := Job{
		JobFunc:     jobFunc,
		JobParams:   jobParams,
		RunLoop:     true,
		Interval:    interval,
		TimeRun:     time.Time{},
		MustDeleted: false,
	}
	Jobs = append(Jobs, job)
}

func RegisterJobRunOnce(jobFunc interface{}, interval time.Duration, jobParams ...interface{}) {
	job := Job{
		JobFunc:     jobFunc,
		JobParams:   jobParams,
		RunLoop:     false,
		Interval:    interval,
		TimeRun:     time.Time{},
		MustDeleted: false,
	}
	Jobs = append(Jobs, job)
}
