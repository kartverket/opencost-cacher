package opencost

import (
	"time"
)

type Response struct {
	Code   int                         `json:"code"`
	Status string                      `json:"status"`
	Data   []map[string]ResourceResult `json:"data"`
}

type ResourceProperties struct {
	Cluster         string            `json:"cluster"`
	Node            string            `json:"node"`
	Container       string            `json:"container"`
	Namespace       string            `json:"namespace"`
	Pod             string            `json:"pod"`
	ProviderID      string            `json:"providerID"`
	NamespaceLabels map[string]string `json:"namespaceLabels"`
}

type SearchWindow struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type ResourceResult struct {
	Name                       string             `json:"name"`
	Properties                 ResourceProperties `json:"properties"`
	Window                     SearchWindow       `json:"window"`
	Start                      time.Time          `json:"start"`
	End                        time.Time          `json:"end"`
	Minutes                    float64            `json:"minutes"`
	CpuCores                   float64            `json:"cpuCores"`
	CpuCoreRequestAverage      float64            `json:"cpuCoreRequestAverage"`
	CpuCoreUsageAverage        float64            `json:"cpuCoreUsageAverage"`
	CpuCoreHours               float64            `json:"cpuCoreHours"`
	CpuCost                    float64            `json:"cpuCost"`
	CpuCostAdjustment          float64            `json:"cpuCostAdjustment"`
	CpuEfficiency              float64            `json:"cpuEfficiency"`
	GpuCount                   float64            `json:"gpuCount"`
	GpuRequestAverage          float64            `json:"gpuRequestAverage"`
	GpuUsageAverage            float64            `json:"gpuUsageAverage"`
	GpuHours                   float64            `json:"gpuHours"`
	GpuCost                    float64            `json:"gpuCost"`
	GpuCostAdjustment          float64            `json:"gpuCostAdjustment"`
	GpuEfficiency              float64            `json:"gpuEfficiency"`
	NetworkTransferBytes       float64            `json:"networkTransferBytes"`
	NetworkReceiveBytes        float64            `json:"networkReceiveBytes"`
	NetworkCost                float64            `json:"networkCost"`
	NetworkCrossZoneCost       float64            `json:"networkCrossZoneCost"`
	NetworkCrossRegionCost     float64            `json:"networkCrossRegionCost"`
	NetworkInternetCost        float64            `json:"networkInternetCost"`
	NetworkCostAdjustment      float64            `json:"networkCostAdjustment"`
	LoadBalancerCost           float64            `json:"loadBalancerCost"`
	LoadBalancerCostAdjustment float64            `json:"loadBalancerCostAdjustment"`
	PvBytes                    float64            `json:"pvBytes"`
	PvByteHours                float64            `json:"pvByteHours"`
	PvCost                     float64            `json:"pvCost"`
	Pvs                        interface{}        `json:"pvs"`
	PvCostAdjustment           float64            `json:"pvCostAdjustment"`
	RamBytes                   float64            `json:"ramBytes"`
	RamByteRequestAverage      float64            `json:"ramByteRequestAverage"`
	RamByteUsageAverage        float64            `json:"ramByteUsageAverage"`
	RamByteHours               float64            `json:"ramByteHours"`
	RamCost                    float64            `json:"ramCost"`
	RamCostAdjustment          float64            `json:"ramCostAdjustment"`
	RamEfficiency              float64            `json:"ramEfficiency"`
	ExternalCost               float64            `json:"externalCost"`
	SharedCost                 float64            `json:"sharedCost"`
	TotalCost                  float64            `json:"totalCost"`
	TotalEfficiency            float64            `json:"totalEfficiency"`
	RawAllocationOnly          struct {
		CpuCoreUsageMax float64 `json:"cpuCoreUsageMax"`
		RamByteUsageMax float64 `json:"ramByteUsageMax"`
	} `json:"rawAllocationOnly"`
	LbAllocations interface{} `json:"lbAllocations"`
}
