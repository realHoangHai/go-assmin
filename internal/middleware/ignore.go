package middleware

import "github.com/gin-gonic/gin"

type IgnoreFunc func(*gin.Context) bool

func (h *Handler) IgnorePathPrefixes(prefixes ...string) IgnoreFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)
		for _, prefix := range prefixes {
			if pl := len(prefixes); pathLen >= pl && path[:pl] == prefix {
				return true
			}
		}
		return false
	}
}

func (h *Handler) UnignorePathPrefixes(prefixes ...string) IgnoreFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)
		for _, prefix := range prefixes {
			if pl := len(prefixes); pathLen >= pl && path[:pl] == prefix {
				return false
			}
		}
		return true
	}
}

func IgnoreHandler(c *gin.Context, ignores ...IgnoreFunc) bool {
	for _, ignore := range ignores {
		if ignore(c) {
			return true
		}
	}
	return false
}
