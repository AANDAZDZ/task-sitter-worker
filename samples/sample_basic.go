package main

import (
	"task-sitter-worker/sdk/core/schedule"
)

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
	schedule.RegisterTrigger(&schedule.Trigger{
		TaskType:         "basic",
		TriggerInterface: basic,
	})
}
