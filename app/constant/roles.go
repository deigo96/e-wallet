package constant

const (
	ROLE_UNKOWN = iota
	ROLE_SUPER_ADMIN
	ROLE_ADMIN
	ROLE_USER
)

var roles = map[int]string{
	ROLE_UNKOWN:      "Unknown",
	ROLE_SUPER_ADMIN: "Super Admin",
	ROLE_ADMIN:       "Admin",
	ROLE_USER:        "User",
}

func GetRoleName(role int) string {
	if name, ok := roles[role]; ok {
		return name
	}
	return roles[ROLE_UNKOWN]
}

func GetRoleValue(name string) int {
	for key, value := range roles {
		if value == name {
			return key
		}
	}
	return ROLE_UNKOWN
}

// func GetRoleName(role int) string {
// 	switch role {
// 	case ROLE_SUPER_ADMIN:
// 		return "Super Admin"
// 	case ROLE_ADMIN:
// 		return "Admin"
// 	case ROLE_USER:
// 		return "User"
// 	default:
// 		return "Unknown"
// 	}
// }
