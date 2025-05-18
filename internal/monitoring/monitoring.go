package monitoring

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

var (
	queryDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "query_duration_seconds",
			Help: "Duration of SQL queries",
		},
		[]string{"query"},
	)
	queryRows = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "query_result_row_count",
			Help: "Number of rows returned by query",
		},
		[]string{"query"},
	)
)

const (
	topCitiesQuery = ` select c.name as city_name, count(*) as client_count
                       from clients cl
                       join cities c on cl.city_id = c.id
                       group by c.name
                       order by client_count desc `

	activeBookingsQuery = ` select
                            c.full_name as client_name,
                            t.title as tour_title,
                            b.status as booking_status
                            from bookings b
                            join booking_agreements a on b.id = a.booking_id
                                join clients c on a.client_id = c.id
                                    join tours t on t.id = b.tour_id
                            where b.status not in ('draft', 'cancellation_requested', 'cancelled')
                            order by b.created_at desc `
)

func init() {
	prometheus.MustRegister(queryDuration, queryRows)
}

func Start(db *sqlx.DB) {
	log.Printf("⏳ Starting monitoring server on :8080/monitoring ...")
	go startMetricsServer()
	go runMetricsLoop(db)
}

func startMetricsServer() {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func runMetricsLoop(db *sqlx.DB) {
	for {
		runQuery(db, "top cities", topCitiesQuery)
		runQuery(db, "active bookings", activeBookingsQuery)
		time.Sleep(15 * time.Second)
	}
}

func runQuery(db *sqlx.DB, title string, query string) {
	start := time.Now()
	rows, err := db.Query(query)
	duration := time.Since(start).Seconds()

	if err != nil {
		log.Printf("⚠️ query %s failed: %v", query, err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("rows.Close() failed: %v", err)
		}
	}(rows)

	count := 0
	for rows.Next() {
		count++
	}
	log.Printf("✅ query %s executed in %.3f seconds, rows: %d", title, duration, count)

	queryDuration.WithLabelValues(title).Observe(duration)
	queryRows.WithLabelValues(title).Set(float64(count))
}
