package consulinfo

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func NewConsulInfo(configPath string) *ConsulInfo {
	// 单例模式
	if ConsulInfoPtr == nil {
		ConsulInfoPtr = &ConsulInfo{
			ConfigPath: configPath,
		}
	} else {
		ConsulInfoPtr.ConfigPath = configPath
	}
	return ConsulInfoPtr
}

func (consulInfoPtr *ConsulInfo) LoadConfig() {
	// 读取 YAML 文件
	data, err := os.ReadFile(consulInfoPtr.ConfigPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 解析 YAML 文件内容
	err = yaml.Unmarshal(data, &consulInfoPtr.Configs)
	if err != nil {
		fmt.Println("Error parsing YAML file:", err)
		return
	}
}

func (consulInfoPtr *ConsulInfo) GetConsulToken() string {
	return consulInfoPtr.Configs.Token
}

func (consulInfoPtr *ConsulInfo) GetServicePrefix() string {
	return consulInfoPtr.Configs.Prefix
}
