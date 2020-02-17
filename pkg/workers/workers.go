package workers

import (
	"time"
	"reflect"
)

type Job struct {
	ID       int64
	ticker   *time.Ticker
	done     chan bool
}

type Worker struct {
	jobs []*Job
}

func NewWorker() *Worker {

	return &Worker{[]*Job{}}
}


func (w *Worker) NewJob(id int64, interval int64) *Job {
	j := &Job{
		ID:       id,
		ticker:   time.NewTicker(time.Duration(interval) * time.Second),
		done:     make(chan bool),
	}

	_ = append(w.jobs, j)
	return j
}

func (w *Worker) RemoveJob(id int64) bool {
	for i, j := range w.jobs {
		if j.ID == id {
			w.jobs = append(w.jobs[:i], w.jobs[i+1])
			return true
		}
	}
	return false
}

func (w *Worker) FindJob(id int64) *Job {
	for _, j := range w.jobs {
		if j.ID == id {
			return j
		}
	}
	return nil
}

func (j *Job) Run(fn interface{}, args ...interface{}) {
	
	v := reflect.ValueOf(fn)
    rargs := make([]reflect.Value, len(args))
    for i, a := range args {
        rargs[i] = reflect.ValueOf(a)
    }
	
	go func() {
		for {
			select {
			case <-j.ticker.C:
				v.Call(rargs)
			case <-j.done:
				return
			}
		}
	}()
}

func (j *Job) Stop() {
	close(j.done)
	j.ticker.Stop()
}
