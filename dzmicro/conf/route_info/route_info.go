package routeinfo

// 单例模式
var RouteInfoPtr *RouteInfo

type RouteInfoInterface interface {
	LoadConfig()
	GetServiceName() string
	GetServiceIp() string
	GetServicePort() string
	GetServiceTags() []string
}

type NewInterface interface {
	NewRouteInfo(configPath string) *RouteInfo
}

type RouteInfo struct {
	ConfigPath string
	Configs    struct {
		Service struct {
			Name string   `yaml:"name"`
			Ip   string   `yaml:"ip"`
			Port string   `yaml:"port"`
			Tags []string `yaml:"tags"`
		} `yaml:"service"`
	}
}
