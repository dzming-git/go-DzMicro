package consulinfo

// 单例模式
var ConsulInfoPtr *ConsulInfo

type ConsulInfoInterface interface {
	LoadConfig()
	GetConsulToken() string
	GetServicePrefix() string
}

type NewInterface interface {
	NewConsulInfo(configPath string) *ConsulInfo
}

type ConsulInfo struct {
	ConfigPath string
	Configs    struct {
		Token  string `yaml:"TOKEN"`
		Prefix string `yaml:"PREFIX"`
	}
}
