package database

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type MapStringString map[string]string

func (m MapStringString) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *MapStringString) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, m)
}

type Report struct {
	Name                       string          `gorm:"primaryKey;index"`
	Cluster                    string          `gorm:"primaryKey;index"`
	Type                       string          `gorm:"primaryKey;index"`
	NamespaceLabels            MapStringString `gorm:"type:json"`
	TeamLabel                  string
	DivisionLabel              string
	Namespace                  string
	Container                  string
	Start                      time.Time `gorm:"primaryKey;index"`
	End                        time.Time
	CpuCores                   float64
	CpuCoreRequestAverage      float64
	CpuCoreUsageAverage        float64
	CpuCoreHours               float64
	CpuCost                    float64
	CpuCostAdjustment          float64
	CpuEfficiency              float64
	GpuCount                   float64
	GpuRequestAverage          float64
	GpuUsageAverage            float64
	GpuHours                   float64
	GpuCost                    float64
	GpuCostAdjustment          float64
	GpuEfficiency              float64
	NetworkTransferBytes       float64
	NetworkReceiveBytes        float64
	NetworkCost                float64
	NetworkCrossZoneCost       float64
	NetworkCrossRegionCost     float64
	NetworkInternetCost        float64
	NetworkCostAdjustment      float64
	LoadBalancerCost           float64
	LoadBalancerCostAdjustment float64
	PvBytes                    float64
	PvByteHours                float64
	PvCost                     float64
	Pvs                        string
	PvCostAdjustment           float64
	RamBytes                   float64
	RamByteRequestAverage      float64
	RamByteUsageAverage        float64
	RamByteHours               float64
	RamCost                    float64
	RamCostAdjustment          float64
	RamEfficiency              float64
	ExternalCost               float64
	SharedCost                 float64
	TotalCost                  float64
	TotalEfficiency            float64
	CpuCoreUsageMax            float64
	RamByteUsageMax            float64
	CreatedAt                  time.Time `gorm:"autoCreateTime"`
	UpdatedAt                  time.Time `gorm:"autoUpdateTime"`
}
