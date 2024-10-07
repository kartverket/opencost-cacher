package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"kartverket/skip/opencost/pkg/config"
	"kartverket/skip/opencost/pkg/database"
	"kartverket/skip/opencost/pkg/rest"
	"kartverket/skip/opencost/pkg/scheduler"
	"os"
	"time"
)

func dailySync(s *scheduler.Scheduler) {
	for {
		now := time.Now()
		nextRunTime := time.Date(now.Year(), now.Month(), now.Day(), 03, 0, 0, 0, now.Location())
		if now.After(nextRunTime) {
			nextRunTime = nextRunTime.Add(24 * time.Hour)
			fmt.Println(fmt.Sprintf("Next run time is tomorrow at %s", nextRunTime.Format("2006-01-02 15:04:05")))
		}

		durationUntilNextRun := nextRunTime.Sub(now)
		time.Sleep(durationUntilNextRun)

		s.SyncAllReports()
	}
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("error loading config: %v", err))
		os.Exit(1)
	}

	var db *gorm.DB
	if cfg.LocalDB {
		db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	} else {
		db, err = gorm.Open(postgres.Open(cfg.DatabaseConfig), &gorm.Config{})
	}

	if err != nil {
		fmt.Println(fmt.Errorf("error connecting to database: %v", err))
		os.Exit(1)
	}

	db.AutoMigrate(&database.Report{})
	dbClient := database.NewClient(db)

	for cluster, opencostUrl := range cfg.OpenCostURLs {
		fmt.Println(fmt.Sprintf("Starting scheduler for cluster %s with URL %s", cluster, opencostUrl))

		instanceScheduler := scheduler.NewScheduler(dbClient, "container,namespace", opencostUrl, cfg.FullSync, cluster)

		go instanceScheduler.SyncAllReports()

		go func() {
			dailySync(instanceScheduler)
		}()
	}

	r := gin.Default()
	corsConfig := cors.Config{
		AllowOrigins:  []string{"http://localhost:3000"},
		AllowMethods:  []string{"GET"},
		AllowHeaders:  []string{"Origin", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
	}
	r.Use(cors.New(corsConfig))

	handler := rest.NewRestHandler(dbClient, "container,namespace")
	r.GET("/reports", handler.HandleGET)
	r.Run()
}
