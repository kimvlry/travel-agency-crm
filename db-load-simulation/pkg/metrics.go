package simulation

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	startTime      = time.Now()
	queryTimings   = make(map[string]float64)
	queryTimingsMu sync.RWMutex
)

func UpdateQueryTiming(queryName string, duration float64) {
	queryTimingsMu.Lock()
	defer queryTimingsMu.Unlock()
	queryTimings[queryName] = duration
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime).Seconds()

	queryTimingsMu.RLock()
	defer queryTimingsMu.RUnlock()

	metrics := fmt.Sprintf(`
		# HELP db_load_uptime_seconds Application uptime in seconds
		# TYPE db_load_uptime_seconds counter
		db_load_uptime_seconds %.2f
	`, uptime)

	for query, timing := range queryTimings {
		metrics += fmt.Sprintf(`
			# HELP db_query_duration_seconds Query execution duration in seconds
			# TYPE db_query_duration_seconds gauge
			db_query_duration_seconds{query="%s"} %.4f
		`, query, timing)
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(metrics))
}

func SetupMetrics() {
	http.HandleFunc("/metrics", metricsHandler)
}
