package gorush

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
)

// PrintGoRushVersion provide print server engine
func PrintGoRushVersion() {
	fmt.Printf(`GoRush %s, Compiler: %s %s, Copyright (C) 2016 Bo-Yi Wu, Inc.`,
		Version,
		runtime.Compiler,
		runtime.Version())
}

// VersionMiddleware : add version on header.
func VersionMiddleware() gin.HandlerFunc {
	// Set out header value for each response
	return func(c *gin.Context) {
		c.Writer.Header().Set("Server-Version", "GoRush/"+Version)
		c.Next()
	}
}
