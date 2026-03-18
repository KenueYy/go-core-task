package main

type CustomWaitGroup struct {
	sem chan struct{}
}

func NewCustomWaitGroup(size int) *CustomWaitGroup {
	if size <= 0 {
		size = 1
	}
	return &CustomWaitGroup{
		sem: make(chan struct{}, size),
	}
}

func (w *CustomWaitGroup) Add() {
	w.sem <- struct{}{}
}

func (w *CustomWaitGroup) Done() {
	<-w.sem
}

func (w *CustomWaitGroup) Wait() {
	for i := 0; i < cap(w.sem); i++ {
		select {
		case <-w.sem:
		default:
			return
		}
	}
}
