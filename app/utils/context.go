package utils

import (
	"github.com/gin-gonic/gin"
)

type Context struct {
	ID   int
	Role string
}

func GetContext(c *gin.Context) *Context {
	id, exist := c.Get("id")
	if !exist {
		id = 0
	}
	role, exist := c.Get("role")
	if !exist {
		role = ""
	}
	return &Context{ID: id.(int), Role: role.(string)}
}
