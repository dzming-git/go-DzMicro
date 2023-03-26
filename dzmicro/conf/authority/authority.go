package authority

type AuthorityInterface interface {
	LoadConfig()
	GetPermissionLevel(sourceID []string) (bool, int)
	GetPermissionByLevel(level int) (bool, string)
	CheckCommandPermission(command string, sourceID []string) (bool, bool)
}
