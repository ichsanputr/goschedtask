package goschedtask

import (
	"errors"
	"reflect"
	"time"
)

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
	if err := JobParamsChecking(jobFunc, jobParams); err != nil {
		panic(err)
	}

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
	if err := JobParamsChecking(jobFunc, jobParams); err != nil {
		panic(err)
	}

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

func RegisterJobRunAt(jobFunc interface{}, timeRun time.Time, jobParams ...interface{}) {
	if err := JobParamsChecking(jobFunc, jobParams); err != nil {
		panic(err)
	}

	job := Job{
		JobFunc:     jobFunc,
		JobParams:   jobParams,
		RunLoop:     false,
		TimeRun:     timeRun,
		MustDeleted: false,
	}
	Jobs = append(Jobs, job)
}

func JobParamsChecking(jobFunc interface{}, jobParams []interface{}) error {
	if len(jobParams) != reflect.ValueOf(jobFunc).Type().NumIn() {
		return errors.New("the number of job function params is not the same as the function")
	}

	return nil
}
