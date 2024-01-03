package worker

import "task-sitter-worker/sdk/model"

var triggerClientMap map[string]*TriggerClient

type TriggerClient struct {
	TaskType string
	TriggerInterface
}

type TriggerServer struct {
	Client   *TriggerClient
	TaskBase *model.TaskBase
}

type TriggerInterface interface {
	Pretreatment() error
	Executing() error
	Finish() error
	ErrorHandle(error)
}

func RegisterTriggerClient(trigger *TriggerClient) {
	if _, ok := triggerClientMap[trigger.TaskType]; ok {
		return
	}
	triggerClientMap[trigger.TaskType] = trigger
}

func AcquireTriggerServer(taskBase *model.TaskBase) (trigger *TriggerServer) {
	if client, ok := triggerClientMap[taskBase.TaskType]; ok {
		trigger = &TriggerServer{
			Client:   client,
			TaskBase: taskBase,
		}
	}
	return
}
