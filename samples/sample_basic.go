package main

import "task-sitter-worker/sdk"

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

func (t *BasicTask) SetBase(task *sdk.TaskBase) {
	return
}

func (t *BasicTask) GetType() string {
	return ""
}

func main() {
	basic := &BasicTask{}
	sdk.RegisterTrigger(basic)
}
