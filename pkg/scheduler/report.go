package scheduler

import (
	"errors"
	"fmt"
	"kartverket/skip/opencost/pkg/database"
	"kartverket/skip/opencost/pkg/opencost"
	"time"
)

type Scheduler struct {
	dbClient    *database.Client
	reportType  string
	opencostURL string
	fullSync    bool
	cluster     string
}

func NewScheduler(dbClient *database.Client, reportType string, opencostURL string, fullSync bool, cluster string) *Scheduler {
	return &Scheduler{
		dbClient:    dbClient,
		reportType:  reportType,
		opencostURL: opencostURL,
		cluster:     cluster,
		fullSync:    fullSync,
	}
}

func getWindowFromDate(date time.Time) string {
	dateAsString := date.Format("2006-01-02")
	startDate := fmt.Sprintf("%sT00:00:00Z", dateAsString)
	endDate := fmt.Sprintf("%sT23:59:59Z", dateAsString)
	window := fmt.Sprintf("%s,%s", startDate, endDate)

	return window
}

func (s *Scheduler) saveLatestReport() {
	yesterday := time.Now().UTC().AddDate(0, 0, -1)
	window := getWindowFromDate(yesterday)
	s.saveReportForWindow(window)
}

func (s *Scheduler) saveReportForWindow(window string) error {
	opencostReport, err := opencost.GetReport(s.opencostURL, window, s.reportType)
	if err != nil {
		fmt.Println("Error when trying to get report:", err)
		return err
	}

	if len(opencostReport.Data[0]) == 0 {
		return errors.New("no data found, probably reached the end of the data")
	}

	for _, report := range opencostReport.Data[0] {
		report := database.MapToDatabaseObject(report, s.reportType, s.cluster)
		if err = s.dbClient.SaveReport(&report); err != nil {
			fmt.Println(fmt.Sprintf("Error when trying to save report %s: %w", report.Name, err))
		}
		fmt.Println(fmt.Sprintf("Report saved successfully: %s", report.Name))
	}

	return nil
}

func (s *Scheduler) SyncAllReports() {
	dateToCheck := time.Now().UTC()
	noData := false
	for !noData {
		dateToCheck = dateToCheck.AddDate(0, 0, -1)

		// check if previous report exists
		isSaved, err := s.dbClient.IsDateSaved(dateToCheck, s.cluster, s.reportType)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error when trying to check if report exists for date %s in sync all reports:%w", dateToCheck.Format("2006-01-02"), err))
			noData = true
			continue
		}
		if isSaved && !s.fullSync {
			fmt.Println(fmt.Sprintf("Report for date %s already saved", dateToCheck.Format("2006-01-02")))
			continue
		}

		fmt.Println(fmt.Sprintf("Syncing report for date: %s", dateToCheck.Format("2006-01-02")))
		window := getWindowFromDate(dateToCheck)
		if err := s.saveReportForWindow(window); err != nil {
			fmt.Println(fmt.Sprintf("Error when trying to save report for current date %s in sync all reports:%w", dateToCheck.Format("2006-01-02"), err))
			noData = true
		}
	}
	fmt.Println(fmt.Sprintf("All reports synced for cluster %s", s.cluster))
}
