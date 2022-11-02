package request

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	errors "github.com/realHoangHai/go-assmin/internal/common/errors"
	"strconv"
	"strings"
)

// GetToken return to authorize token extracted from header
func GetToken(c *gin.Context) string {
	var token string
	header := c.GetHeader("Authorization")
	prefix := "Bearer "
	if header != "" && strings.HasPrefix(header, prefix) {
		token = header[len(prefix):]
	}
	return token
}

// ParseUUID return the value of the url param in uuid type
func ParseUUID(c *gin.Context, key string) uuid.UUID {
	uid, err := uuid.Parse(c.Param(key))
	if err != nil {
		return uuid.Nil
	}
	return uid
}

// ParseID return the value of the url param in uint64 type
func ParseID(c *gin.Context, key string) uint64 {
	id, err := strconv.ParseUint(c.Param(key), 10, 64)
	if err != nil {
		return 0
	}
	return id
}

// ParseJson will parse body json data to struct
func ParseJson(c *gin.Context, v interface{}) error {
	if err := c.ShouldBindJSON(v); err != nil {
		return errors.ErrInvalidRequest(err, fmt.Sprintf("parse request json failed: %s", err.Error()))
	}
	return nil
}

// ParseForm will parse body form data to struct
func ParseForm(c *gin.Context, v interface{}) error {
	if err := c.ShouldBindWith(v, binding.Form); err != nil {
		return errors.ErrInvalidRequest(err, fmt.Sprintf("parse request form failed: %s", err.Error()))
	}
	return nil
}
