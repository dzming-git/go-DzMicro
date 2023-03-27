package services

// 单例模式
var ServiceFuncMapPtr *ServiceFuncMap

type ServiceFuncMapInterface interface {
	AddServiceFuncMap(command string, serviceFunc func(), permission string)
	SetKeyword(keyword string)
	GetKeyword() string
	GetServiceFunc(command string) func()
	GetServicePermission(command string) string
}

type NewInterface interface {
	NewServiceFuncMap() *ServiceFuncMap
	NewTaskInfo(sourceID []string, args []string, platformIpPort [2]string) *TaskInfo
}

type TaskInfo struct {
	SourceID []string
	Args     []string
	Platform struct {
		Ip   string
		Port string
	}
}

type ServiceFuncMap struct {
	Keyword         string
	ServiceFuncInfo map[string]struct {
		ServiceFunc func()
		Permission  string
	}
}
