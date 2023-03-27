package services

func NewServiceFuncMap() *ServiceFuncMap {
	if ServiceFuncMapPtr == nil {
		ServiceFuncMapPtr = &ServiceFuncMap{
			ServiceFuncInfo: make(map[string]struct {
				ServiceFunc func()
				Permission  string
			}),
		}
	}
	return ServiceFuncMapPtr
}

func (taskInfoPtr *TaskInfo) SetTaskInfo(sourceID []string, args []string, platformIpPort [2]string) {
	taskInfoPtr.SourceID = sourceID
	taskInfoPtr.Args = args
	taskInfoPtr.Platform.Ip = platformIpPort[0]
	taskInfoPtr.Platform.Port = platformIpPort[1]
}

func NewTaskInfo() *TaskInfo {
	return &TaskInfo{}
}

func (serviceFuncMapPtr *ServiceFuncMap) AddServiceFuncMap(command string, serviceFunc func(), permission string) {
	serviceFuncMapPtr.ServiceFuncInfo[command] = struct {
		ServiceFunc func()
		Permission  string
	}{
		ServiceFunc: serviceFunc,
		Permission:  permission,
	}
}

func (serviceFuncMapPtr *ServiceFuncMap) SetKeyword(keyword string) {
	serviceFuncMapPtr.Keyword = keyword
}

func (serviceFuncMapPtr *ServiceFuncMap) GetKeyword() string {
	return serviceFuncMapPtr.Keyword
}

func (serviceFuncMapPtr *ServiceFuncMap) GetServiceFunc(command string) func() {
	return serviceFuncMapPtr.ServiceFuncInfo[command].ServiceFunc
}

func (serviceFuncMapPtr *ServiceFuncMap) GetServicePermission(command string) string {
	return serviceFuncMapPtr.ServiceFuncInfo[command].Permission
}
