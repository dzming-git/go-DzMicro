package routeinfo

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func NewRouteInfo(configPath string) *RouteInfo {
	// 单例模式
	if RouteInfoPtr == nil {
		RouteInfoPtr = &RouteInfo{
			ConfigPath: configPath,
		}
	} else {
		RouteInfoPtr.ConfigPath = configPath
	}
	return RouteInfoPtr
}

func (routeInfoPtr *RouteInfo) LoadConfig() {
	// 读取 YAML 文件
	data, err := os.ReadFile(routeInfoPtr.ConfigPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 解析 YAML 文件内容
	err = yaml.Unmarshal(data, &routeInfoPtr.Configs)
	if err != nil {
		fmt.Println("Error parsing YAML file:", err)
		return
	}
}

func (routeInfoPtr *RouteInfo) GetServiceName() string {
	return routeInfoPtr.Configs.Service.Name
}

func (routeInfoPtr *RouteInfo) GetServiceIp() string {
	return routeInfoPtr.Configs.Service.Ip
}

func (routeInfoPtr *RouteInfo) GetServicePort() string {
	return routeInfoPtr.Configs.Service.Port
}

func (routeInfoPtr *RouteInfo) GetServiceTags() []string {
	return routeInfoPtr.Configs.Service.Tags
}
