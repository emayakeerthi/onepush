package utils

import (
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

var startTime time.Time

func CheckSystemHealth() map[string]any {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	systemInfo := map[string]any{
		"uptime":     time.Since(startTime).Truncate(time.Second).String(),
		"go_version": runtime.Version(),
		"goroutines": runtime.NumGoroutine(),
		"memory": gin.H{
			"alloc_mb":       float64(memStats.Alloc) / 1024 / 1024,
			"total_alloc_mb": float64(memStats.TotalAlloc) / 1024 / 1024,
			"sys_mb":         float64(memStats.Sys) / 1024 / 1024,
			"gc_cycles":      memStats.NumGC,
		},
	}

	// Determine system status dynamically
	status := "healthy"
	if runtime.NumGoroutine() > 1000 || float64(memStats.Alloc)/1024/1024 > 500 {
		status = "unhealthy"
	}

	return map[string]any{
		"status":    status,
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"version":   "1.0.0", // You can get this from build flags or env
		"service":   "onepush-server",
		"system":    systemInfo,
	}
}
