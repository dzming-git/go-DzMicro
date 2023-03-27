package authority

// 单例模式
var AuthorityPtr *Authority

type AuthorityInterface interface {
	LoadConfig()
	GetPermissionLevel(sourceID []string) (bool, int)
	GetPermissionByLevel(level int) (bool, string)
	CheckCommandPermission(command string, sourceID []string) (bool, bool)
}

type NewInterface interface {
	NewAuthority(configPath string) *Authority
}

type Config struct {
	GlobalPermissionFirst bool                        `yaml:"GLOBAL_PERMISSION_FIRST"`
	PermissionLevel       map[string]int              `yaml:"PERMISSION_LEVEL"`
	Authorites            map[interface{}]interface{} `yaml:"AUTHORITIES"`
}

type ConfigConvert struct {
	GlobalPermissionFirst bool
	PermissionLevel       map[string]int
	Authorites            map[string]interface{}
}

type Authority struct {
	ConfigPath string
	Configs    ConfigConvert
}
