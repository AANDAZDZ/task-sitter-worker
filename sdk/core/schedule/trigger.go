package schedule

var triggerMap map[string]*Trigger

type Trigger struct {
	TaskType string
	TriggerInterface
}

type TriggerInterface interface {
	Pretreatment() error
	Executing() error
	Finish() error
	ErrorHandle(error)
}

func RegisterTrigger(trigger *Trigger) {
	if _, ok := triggerMap[trigger.TaskType]; ok {
		return
	}
	triggerMap[trigger.TaskType] = trigger
}
