package service

import (
	"task-sitter-worker/sdk/core/schedule"
	"task-sitter-worker/sdk/model"
)

const (
	CronTask = 0
	RateTask = 1
)

func ScheduleTask(req *model.ScheduleTaskReq) (resp *model.ScheduleTaskResp) {
	taskBase := transferTaskBase(req)
	scheduler := schedule.AcquireScheduler(taskBase)
	var err error
	switch req.ScheduleType {
	case CronTask:
		err = scheduler.CronTask()
	case RateTask:
		err = scheduler.RateTask()
	}
	if err != nil {
	}
	return
}

func transferTaskBase(req *model.ScheduleTaskReq) *model.TaskBase {
	return nil
}
