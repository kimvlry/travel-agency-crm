package simulation

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

var (
	startTime = time.Now()
)

func getCPUUsage() float64 {
	var rusage runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&rusage)

	numCPU := runtime.NumCPU()
	numGoroutine := runtime.NumGoroutine()

	cpuUsage := float64(numGoroutine) / float64(numCPU) * 10.0
	if cpuUsage > 100 {
		cpuUsage = 100
	}

	return cpuUsage
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	cpuUsage := getCPUUsage()
	uptime := time.Since(startTime).Seconds()

	metrics := fmt.Sprintf(`
        # HELP db_load_cpu_usage CPU usage percentage
        # TYPE db_load_cpu_usage gauge
        db_load_cpu_usage %.2f
        
        # HELP db_load_uptime_seconds Application uptime in seconds
        # TYPE db_load_uptime_seconds counter
        db_load_uptime_seconds %.2f
    `, cpuUsage, uptime)

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(metrics))
}

func SetupMetrics() {
	http.HandleFunc("/metrics", metricsHandler)
}
