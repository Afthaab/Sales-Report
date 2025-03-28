package script

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Afthaab/Sales-Report-Lumel/internal/loader"
	"github.com/Afthaab/Sales-Report-Lumel/internal/util"
	"github.com/robfig/cron/v3"
)

func StartCronJob(loader loader.LoaderInterface) {
	c := cron.New()
	_, err := c.AddFunc(util.CRON_JOB_SCHEDULER, func() {
		log.Println("starting scheduled data refresh...")
		err := RunCSVLoader(loader)
		if err != nil {
			LogRefreshEvent("FAILURE", err.Error())
			log.Println("scheduled data refresh failed:", err)
		} else {
			LogRefreshEvent("SUCCESS", "scheduled data refresh completed successfully.")
			log.Println("scheduled data refresh completed successfully.")
		}
	})
	if err != nil {
		log.Fatalf("failed to schedule cron job: %v", err)
	}
	c.Start()
	log.Println("cron job scheduled: Daily data refresh at 1 minute.")
}

func LogRefreshEvent(status string, details string) {
	logDir := util.LOGS_PATH
	logFile := filepath.Join(logDir, "refresh.log")

	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Println("error creating log directory:", err)
		return
	}

	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("error opening log file:", err)
		return
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	logger.Printf("[%s] %s\n", status, details)
}
