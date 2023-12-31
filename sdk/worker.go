package sdk

// Init 初始化并启动 rpc 服务
func Init() {
}

// Listener 监听 rpc 请求
func Listener() {
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
