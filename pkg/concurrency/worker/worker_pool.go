package worker

import (
	"sync"

	"go.uber.org/zap"
)

type pool struct {
	workerCount int
	logger      *zap.SugaredLogger
	jobs        chan *job
}

type worker struct {
	id     int
	logger *zap.SugaredLogger
	jobs   <-chan *job
	wg     *sync.WaitGroup
}

type job struct {
	id   int
	work WorkFunc
}

type WorkFunc func() error

func NewPool(workerCount int, logger *zap.SugaredLogger) *pool {
	return &pool{
		workerCount: workerCount,
		logger:      logger,
		jobs:        make(chan *job),
	}
}

func (p *pool) Start() {
	var wg sync.WaitGroup

	for i := 1; i <= p.workerCount; i++ {
		wg.Add(1)
		worker := newWorker(i, p.logger, p.jobs, &wg)
		p.logger.Infof("Worker %v has been spawned.", worker.id)
		go worker.Work()
	}

	p.logger.Infof("Worker Pool started with %v workers.", p.workerCount)
	wg.Wait()
}

func (p *pool) Stop() {
	close(p.jobs)
}

func (p *pool) SubmitJob(job *job) {
	p.logger.Infof("Submitting job %v.", job.id)
	p.jobs <- job
}

func newWorker(id int, logger *zap.SugaredLogger, jobs <-chan *job, wg *sync.WaitGroup) *worker {
	return &worker{
		id:     id,
		logger: logger,
		jobs:   jobs,
		wg:     wg,
	}
}

func (w *worker) Work() {
	defer func(wg *sync.WaitGroup) {
		w.logger.Infof("Completing worker %v.", w.id)
		wg.Done()
	}(w.wg)

	for job := range w.jobs {
		w.logger.Infof("Worker %v picked up job %v.", w.id, job.id)
		err := job.Complete()
		if err != nil {
			w.logger.Errorf("Worker %v encountered an error with a job %v. %v+", w.id, job.id, err)
		}
	}
}

func NewJob(id int, workFunc WorkFunc) *job {
	return &job{
		id:   id,
		work: workFunc,
	}
}

func (j *job) Complete() error {
	err := j.work()
	if err != nil {
		return err
	}
	return nil
}
