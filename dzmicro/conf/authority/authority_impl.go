package authority

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	GlobalPermissionFirst bool                        `yaml:"GLOBAL_PERMISSION_FIRST"`
	PermissionLevel       map[string]int              `yaml:"PERMISSION_LEVEL"`
	Authorites            map[interface{}]interface{} `yaml:"AUTHORITIES"`
}

type ConfigConvert struct {
	GlobalPermissionFirst bool                   `yaml:"GLOBAL_PERMISSION_FIRST"`
	PermissionLevel       map[string]int         `yaml:"PERMISSION_LEVEL"`
	Authorites            map[string]interface{} `yaml:"AUTHORITIES"`
}

type Authority struct {
	ConfigPath string
	Configs    ConfigConvert
}

func NewAuthority(configPath string) *Authority {
	return &Authority{
		ConfigPath: configPath,
	}
}

// 转换函数，将 map[interface{}]interface{} 类型数据转换为 map[string]interface{} 类型
func convert(m map[interface{}]interface{}) interface{} {
	// 创建一个空的 map[string]interface{} 类型的变量 newM
	newM := make(map[string]interface{})
	// 遍历 m 中的数据
	for k, v := range m {
		// 如果 v 是 map[interface{}]interface{} 类型，则递归调用 convert 函数进行转换并将结果放入 newM 中
		if child, ok := v.(map[interface{}]interface{}); ok {
			newM[k.(string)] = convert(child)
		} else {
			// 否则直接将键和值放入 newM 中
			newM[k.(string)] = v
		}
	}
	return newM
}

func (authorityPtr *Authority) LoadConfig() {
	// 读取 YAML 文件
	data, err := os.ReadFile(authorityPtr.ConfigPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 解析 YAML 文件内容
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing YAML file:", err)
		return
	}
	authorityPtr.Configs.GlobalPermissionFirst = config.GlobalPermissionFirst
	authorityPtr.Configs.PermissionLevel = config.PermissionLevel
	authorityPtr.Configs.Authorites = convert(config.Authorites).(map[string]interface{})
}

func (authorityPtr *Authority) GetPermissionLevel(sourceID []string) (bool, int) {
	permissionLevelNotNone := false
	var permissionLevel int
	authorities := authorityPtr.Configs.Authorites
	idIndex := 0
	maxIndex := len(sourceID)
	backtrackingTimes := 0

	type BacktrackingNode struct {
		idIndex           int
		authorities       map[string]interface{}
		backtrackingTimes int
	}
	NewBacktrackingNode := func(idIndex int, authorities map[string]interface{}, backtrackingTimes int) *BacktrackingNode {
		return &BacktrackingNode{
			idIndex:           idIndex,
			authorities:       authorities,
			backtrackingTimes: backtrackingTimes,
		}
	}

	BacktrackingPathStack := make([]*BacktrackingNode, 0, 10)

	for {
		xID := sourceID[idIndex]
		if _, ok := authorities["GLOBAL"]; ok && backtrackingTimes <= 0 {
			if _, ok := authorities[xID]; !ok && authorityPtr.Configs.GlobalPermissionFirst {
				// 如果x_id没查询到，并且global优先为设定，则不使用global权限
				// global优先未设定时，只有参与配置的id可以使用global权限
				break
			}
			backtrackingNodePtr := NewBacktrackingNode(idIndex, authorities, backtrackingTimes)
			BacktrackingPathStack = append(BacktrackingPathStack, backtrackingNodePtr)
			authorities, _ = authorities["GLOBAL"].(map[string]interface{})
		} else if _, ok := authorities[xID]; ok && backtrackingTimes <= 1 {
			backtrackingNodePtr := NewBacktrackingNode(idIndex, authorities, backtrackingTimes)
			BacktrackingPathStack = append(BacktrackingPathStack, backtrackingNodePtr)
			authorities, _ = authorities[xID].(map[string]interface{})
		} else if _, ok := authorities["DEFAULT"]; ok && backtrackingTimes <= 2 {
			backtrackingNodePtr := NewBacktrackingNode(idIndex, authorities, backtrackingTimes)
			BacktrackingPathStack = append(BacktrackingPathStack, backtrackingNodePtr)
			authorities, _ = authorities["DEFAULT"].(map[string]interface{})
		} else if len(BacktrackingPathStack) > 0 {
			// 找不到权限信息，回溯
			backtrackingNodePtr := BacktrackingPathStack[len(BacktrackingPathStack)-1]
			idIndex = backtrackingNodePtr.idIndex
			authorities = backtrackingNodePtr.authorities
			backtrackingTimes = backtrackingNodePtr.backtrackingTimes + 1
			// 出栈
			BacktrackingPathStack = BacktrackingPathStack[:len(BacktrackingPathStack)-1]
			continue
		}
		idIndex++
		backtrackingTimes = 0
		if _, ok := authorities["PERMISSION"]; ok || idIndex >= maxIndex {
			permissionLevel, _ = authorities["PERMISSION"].(int)
			permissionLevelNotNone = true
			break
		}
	}

	return permissionLevelNotNone, permissionLevel
}

func (authorityPtr *Authority) GetPermissionByLevel(level int) (bool, string) {
	for permission, l := range authorityPtr.Configs.PermissionLevel {
		if l == level {
			return true, permission
		}
	}
	return false, ""
}

func (authorityPtr *Authority) CheckCommandPermission(command string, sourceID []string) (bool, bool) {
	permissionLevelNotNone, permissionLevel := authorityPtr.GetPermissionLevel(sourceID)
	permissionNeed := "USER" //TODO func_dict.get_permission(command)
	if permissionLevelNeed, ok := authorityPtr.Configs.PermissionLevel[permissionNeed]; ok {
		if permissionLevelNeed == -3 {
			// 只准内部调用，不对用户开放
			return false, false
		}
		if permissionLevelNotNone {
			if permissionLevel == -2 {
				// 最高权限
				return true, true
			} else if permissionLevel == -1 {
				// 禁止权限
				return true, false
			} else {
				// 一般权限
				return true, permissionLevel >= permissionLevelNeed
			}
		} else {
			// 该群没有被配置
			return false, false
		}
	} else {
		fmt.Println("func_dict中权限配置错误")
		return false, false
	}

}
