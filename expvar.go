package expvar

import (
	"encoding/json"
	"expvar"

	"github.com/gin-gonic/gin"
)

type Options struct {
	filters map[string]struct{}
}

type Option func(options *Options)

func WithFilters(filters ...string) Option {
	return func(options *Options) {
		options.filters = map[string]struct{}{}
		for _, filter := range filters {
			options.filters[filter] = struct{}{}
		}
	}
}

// Handler for gin framework
func Handler(opts ...Option) gin.HandlerFunc {
	options := Options{}
	for _, opt := range opts {
		opt(&options)
	}

	return func(c *gin.Context) {
		w := c.Writer
		c.Header("Content-Type", "application/json; charset=utf-8")
		result := map[string]interface{}{}
		expvar.Do(func(kv expvar.KeyValue) {
			if _, ok := options.filters[kv.Key]; ok {
				return
			}

			v := map[string]interface{}{}
			if err := json.Unmarshal([]byte(kv.Value.String()), &v); err == nil {
				result[kv.Key] = v
			} else {
				result[kv.Key] = kv.Value.String()
			}
		})
		data, _ := json.MarshalIndent(result, "", "    ")
		_, _ = w.Write(data)
		c.AbortWithStatus(200)
	}
}
