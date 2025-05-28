package simulation

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
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

func Start(db *sqlx.DB) {
	go runMetricsLoop(db)
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

	UpdateQueryTiming(title, duration)

	if err != nil {
		log.Printf("⚠️ query %s failed: %v", query, err)
		return
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
}
