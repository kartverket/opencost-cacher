package database

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Client struct {
	*gorm.DB
}

func NewClient(db *gorm.DB) *Client {
	return &Client{
		DB: db,
	}
}

func (c *Client) SaveReport(report *Report) error {
	return c.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(report).Error
}

func (c *Client) IsDateSaved(date time.Time, cluster string, reportType string) (bool, error) {
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	endOfDay := startOfDay.AddDate(0, 0, 1)

	if err := c.DB.Where("start >= ? and start < ? AND cluster = ? AND type = ?", startOfDay, endOfDay, cluster, reportType).First(&Report{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			fmt.Printf("Error querying database: %v\n", err)
			return false, err
		}
	}

	return true, nil
}

func (c *Client) GetReportsByClusterAndTypeAndDate(cluster string, reportType string, startDate time.Time, endDate time.Time) ([]Report, error) {
	var reports []Report

	if err := c.DB.
		Where("cluster = ? AND type = ? AND start > ? AND start < ?", cluster, reportType, startDate, endDate).
		Order("start desc").
		Find(&reports).Error; err != nil {
		fmt.Printf("Error querying database: %v\n", err)
		return nil, err
	}

	return reports, nil
}

func (c *Client) GetReportsForByTypeAndDate(reportType string, startDate time.Time, endDate time.Time) ([]Report, error) {
	var reports []Report

	if err := c.DB.
		Where("type = ? AND start > ? AND start < ?", reportType, startDate, endDate).
		Order("start desc").
		Find(&reports).Error; err != nil {
		fmt.Printf("Error querying database: %v\n", err)
		return nil, err
	}

	return reports, nil
}
