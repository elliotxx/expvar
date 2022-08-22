package expvar

import (
	"encoding/json"
	"expvar"

	"github.com/gin-gonic/gin"
)

// Handler for gin framework
func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := c.Writer
		c.Header("Content-Type", "application/json; charset=utf-8")
		result := map[string]interface{}{}
		expvar.Do(func(kv expvar.KeyValue) {
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
