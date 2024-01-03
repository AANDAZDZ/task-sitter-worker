package worker

import "task-sitter-worker/sdk/model"

func ScheduleTask(req *model.ScheduleTaskReq) (resp *model.ScheduleTaskResp) {
	taskBase := transferTaskBase(req)
	trigger := AcquireTriggerServer(taskBase)
	err := trigger.Client.Pretreatment()
	if err != nil {
		trigger.Client.ErrorHandle(err)
	}
	err = trigger.Client.Executing()
	if err != nil {
		trigger.Client.ErrorHandle(err)
	}
	err = trigger.Client.Finish()
	if err != nil {
		trigger.Client.ErrorHandle(err)
	}
	return
}

func transferTaskBase(req *model.ScheduleTaskReq) *model.TaskBase {
	return nil
}
