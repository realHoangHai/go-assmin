package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/realHoangHai/go-assmin/config"
	"github.com/realHoangHai/go-assmin/internal/common/errors"
	"github.com/spf13/cast"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/redis"
	"strings"
)

func (h *Handler) LimitIP(pattern string, ignorePaths ...IgnoreFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if IgnoreHandler(c, ignorePaths...) {
			c.Next()
			return
		}

		key := getIP(c)
		if ok := h.limit(c, key, pattern); !ok {
			return
		}
		c.Next()
	}
}

func (h *Handler) LimitIPCustom(pattern string, ignorePaths ...IgnoreFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if IgnoreHandler(c, ignorePaths...) {
			c.Next()
			return
		}

		key := getRouteWithIP(c)
		if ok := h.limit(c, key, pattern); !ok {
			return
		}
		c.Next()
	}
}

func getIP(c *gin.Context) string {
	return c.ClientIP()
}

func getRouteWithIP(c *gin.Context) string {
	return routeToKeyString(c.ClientIP() + c.FullPath())
}

func (h *Handler) limit(c *gin.Context, key, pattern string) bool {
	rate, err := h.checkRate(c, key, pattern)
	if err != nil {
		panic(err)
	}
	c.Header("X-RateLimit-Limit", cast.ToString(rate.Limit))
	c.Header("X-RateLimit-Remaining", cast.ToString(rate.Remaining))
	c.Header("X-RateLimit-Reset", cast.ToString(rate.Reset))
	if rate.Reached {
		panic(errors.ErrTooManyRequests)
	}
	return true
}

func (h *Handler) checkRate(c *gin.Context, key, formatted string) (limiter.Context, error) {
	var ctx limiter.Context
	rate, err := limiter.NewRateFromFormatted(formatted)
	if err != nil {
		return ctx, err
	}
	store, err := redis.NewStoreWithOptions(h.redis, limiter.StoreOptions{
		Prefix: fmt.Sprintf("%s:limiter", config.C.Core.Application),
	})
	if err != nil {
		return ctx, err
	}
	limit := limiter.New(store, rate)
	if c.GetBool("limit-once-" + key) {
		return limit.Peek(c, key)
	} else {
		c.Set("limit-once-"+key, true)
		return limit.Get(c, key)
	}
}

func routeToKeyString(routeName string) string {
	routeName = strings.ReplaceAll(routeName, "/", "-")
	routeName = strings.ReplaceAll(routeName, ":", "_")
	return routeName
}
