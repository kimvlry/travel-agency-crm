package main

import (
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"os/exec"
)

func runBackup() {
	log.Println("⏳ Starting backup...")

	script := os.Getenv("BACKUP_SCRIPT_PATH")
	if script == "" {
		log.Fatalf("BACKUP_SCRIPT_PATH env variable not found")
	}
	cmd := exec.Command("/bin/sh", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("❌ Backup failed: %v", err)
	} else {
		log.Println("✅ Backup completed successfully")
	}
}

func main() {
	c := cron.New()

	schedule := os.Getenv("BACKUP_INTERVAL_CRON")
	if schedule == "" {
		schedule = "@daily"
	}

	_, err := c.AddFunc(schedule, runBackup)
	if err != nil {
		log.Fatalf("Failed to add cron function: %v", err)
	}

	log.Printf("Backup scheduler started with schedule: %s", schedule)

	c.Start()
	runBackup()

	select {}
}
