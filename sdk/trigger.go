package sdk

var triggerMap map[string]TriggerInterface

type TaskBase struct {
	TaskType    string
	TaskContext string
}

type TriggerInterface interface {
	Pretreatment() error
	Executing() error
	Finish() error
	ErrorHandle(error)

	SetBase(*TaskBase)
	GetType() string
}

func RegisterTrigger(trigger TriggerInterface) {
	if _, ok := triggerMap[trigger.GetType()]; ok {
		return
	}
	triggerMap[trigger.GetType()] = trigger
}

func GetTrigger(taskType string) TriggerInterface {
	if trigger, ok := triggerMap[taskType]; ok {
		return trigger
	}
	return nil
}
