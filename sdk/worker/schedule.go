package worker

func ScheduleTask() {
	req := &TaskBase{}
	trigger := GetTrigger(req.TaskType)
	trigger.SetBase(req)
	err := trigger.Pretreatment()
	if err != nil {
		trigger.ErrorHandle(err)
	}
	err = trigger.Executing()
	if err != nil {
		trigger.ErrorHandle(err)
	}
	err = trigger.Finish()
	if err != nil {
		trigger.ErrorHandle(err)
	}
}
