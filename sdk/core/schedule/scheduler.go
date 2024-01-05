package schedule

import (
	"task-sitter-worker/sdk/model"
	"time"
)

type Scheduler struct {
	Trigger  *Trigger
	TaskBase *model.TaskBase
}

func AcquireScheduler(taskBase *model.TaskBase) (scheduler *Scheduler) {
	if trigger, ok := triggerMap[taskBase.TaskType]; ok {
		scheduler = &Scheduler{
			Trigger:  trigger,
			TaskBase: taskBase,
		}
	}
	return
}

func (s *Scheduler) CronTask() (err error) {
	if trigger := s.Trigger; trigger != nil {
		execTask(trigger)
	}
	return
}

func (s *Scheduler) RateTask() (err error) {
	if s.Trigger == nil {
		return
	}
	go func() {
		ticker := time.NewTicker(time.Duration(s.TaskBase.RateInterval) * time.Millisecond)
		select {
		case <-ticker.C:
			execTask(s.Trigger)
		}
	}()
	return
}

func execTask(trigger *Trigger) {
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
