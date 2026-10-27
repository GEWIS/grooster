package httputil

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// PathUint parses the named path parameter as an unsigned ID.
func PathUint(c *gin.Context, key string) (uint, error) {
	id, err := strconv.ParseUint(c.Param(key), 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}

// MustPathUint parses the named path parameter as an unsigned ID. When the
// parameter is invalid it writes a 400 response and returns false.
func MustPathUint(c *gin.Context, key string) (uint, bool) {
	id, err := PathUint(c, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid " + key})
		return 0, false
	}

	return id, true
}
