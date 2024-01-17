package goschedtask

import (
	"reflect"
	"time"

	"golang.org/x/exp/slices"
)

type Scheduler struct {
	Jobs []Job
}

var (
	mainScheduler = NewScheduler()
)

func (Sched Scheduler) RunJobs() chan bool {
	mainScheduler.Jobs = Jobs
	mainScheduler.SetFirstTimeRun()

	tick := time.NewTicker(time.Second * 1)
	stopped := make(chan bool)

	go func() {
		for {
			select {
			case <-tick.C:
				mainScheduler.RunPendingJobs()
			case <-stopped:
				return
			}
		}
	}()

	return stopped
}

func (Sched *Scheduler) RunPendingJobs() {
	for i, j := range Sched.Jobs {
		if j.MustDeleted {
			Sched.Jobs = slices.Delete(Sched.Jobs, i, i+1)
			continue
		}

		if ShouldRun(j.TimeRun) {
			go RunJobWithParam(j.JobFunc, j.JobParams)
			Sched.Jobs[i].TimeRun = Sched.Jobs[i].TimeRun.Add(Sched.Jobs[i].Interval)

			if !j.RunLoop {
				Sched.Jobs[i].MustDeleted = true
			}
		}
	}

}

func (Sched Scheduler) SetFirstTimeRun() {
	for i, j := range Sched.Jobs {
		if j.Interval.Seconds() != 0 {
			Sched.Jobs[i].TimeRun = time.Now().Add(j.Interval)
		}
	}
}

func NewScheduler() Scheduler {
	return Scheduler{
		Jobs: []Job{},
	}
}

func ShouldRun(t time.Time) bool {
	if time.Now().After(t) {
		return true
	}

	return false
}

func RunJobWithParam(jobFunc interface{}, params []interface{}) {
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	reflect.ValueOf(jobFunc).Call(in)
}

func RunJobs() chan bool {
	return mainScheduler.RunJobs()
}
