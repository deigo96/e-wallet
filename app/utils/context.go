package utils

import (
	"github.com/deigo96/e-wallet.git/app/constant"
	"github.com/gin-gonic/gin"
)

type Context struct {
	ID   int
	Role string
}

func GetContext(c *gin.Context) *Context {
	id, exist := c.Get("id")
	if !exist {
		id = float64(0)
	}
	role, exist := c.Get("role")
	if !exist {
		role = ""
	}
	floatID := id.(float64)

	return &Context{ID: int(floatID), Role: role.(string)}
}

func GetID(c *gin.Context) int {
	return GetContext(c).ID
}

func IsAdmin(role string) bool {
	return role == constant.GetRoleName(constant.ROLE_ADMIN) ||
		role == constant.GetRoleName(constant.ROLE_SUPER_ADMIN)
}
