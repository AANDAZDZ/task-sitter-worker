package main

import "task-sitter-worker/sdk/worker"

type BasicTask struct {
}

func (t *BasicTask) Pretreatment() (err error) {
	return
}

func (t *BasicTask) Executing() (err error) {
	return
}

func (t *BasicTask) Finish() (err error) {
	return
}

func (t *BasicTask) ErrorHandle(err error) {
	return
}

func main() {
	basic := &BasicTask{}
	worker.RegisterTriggerClient(&worker.TriggerClient{
		TaskType:         "basic",
		TriggerInterface: basic,
	})
}
