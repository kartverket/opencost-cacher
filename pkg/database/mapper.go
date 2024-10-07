package database

import (
	"fmt"
	"kartverket/skip/opencost/pkg/opencost"
)

func MapToDatabaseObject(report opencost.ResourceResult, reportType string, cluster string) Report {
	team := report.Properties.NamespaceLabels["team"]
	division := report.Properties.NamespaceLabels["division"]
	namespace := report.Properties.Namespace
	if namespace == "" {
		namespace = "idle"
	}
	return Report{
		Name:                       report.Name,
		Cluster:                    cluster,
		Type:                       reportType,
		NamespaceLabels:            report.Properties.NamespaceLabels,
		TeamLabel:                  team,
		DivisionLabel:              division,
		Namespace:                  namespace,
		Container:                  report.Properties.Container,
		Start:                      report.Window.Start,
		End:                        report.Window.End,
		CpuCores:                   report.CpuCores,
		CpuCoreRequestAverage:      report.CpuCoreRequestAverage,
		CpuCoreUsageAverage:        report.CpuCoreUsageAverage,
		CpuCoreHours:               report.CpuCoreHours,
		CpuCost:                    report.CpuCost,
		CpuCostAdjustment:          report.CpuCostAdjustment,
		CpuEfficiency:              report.CpuEfficiency,
		GpuCount:                   report.GpuCount,
		GpuRequestAverage:          report.GpuRequestAverage,
		GpuUsageAverage:            report.GpuUsageAverage,
		GpuHours:                   report.GpuHours,
		GpuCost:                    report.GpuCost,
		GpuCostAdjustment:          report.GpuCostAdjustment,
		GpuEfficiency:              report.GpuEfficiency,
		NetworkTransferBytes:       report.NetworkTransferBytes,
		NetworkReceiveBytes:        report.NetworkReceiveBytes,
		NetworkCost:                report.NetworkCost,
		NetworkCrossZoneCost:       report.NetworkCrossZoneCost,
		NetworkCrossRegionCost:     report.NetworkCrossRegionCost,
		NetworkInternetCost:        report.NetworkInternetCost,
		NetworkCostAdjustment:      report.NetworkCostAdjustment,
		LoadBalancerCost:           report.LoadBalancerCost,
		LoadBalancerCostAdjustment: report.LoadBalancerCostAdjustment,
		PvBytes:                    report.PvBytes,
		PvByteHours:                report.PvByteHours,
		PvCost:                     report.PvCost,
		Pvs:                        fmt.Sprint(report.Pvs),
		PvCostAdjustment:           report.PvCostAdjustment,
		RamBytes:                   report.RamBytes,
		RamByteRequestAverage:      report.RamByteRequestAverage,
		RamByteUsageAverage:        report.RamByteUsageAverage,
		RamByteHours:               report.RamByteHours,
		RamCost:                    report.RamCost,
		RamCostAdjustment:          report.RamCostAdjustment,
		RamEfficiency:              report.RamEfficiency,
		ExternalCost:               report.ExternalCost,
		SharedCost:                 report.SharedCost,
		TotalCost:                  report.TotalCost,
		TotalEfficiency:            report.TotalEfficiency,
		CpuCoreUsageMax:            report.RawAllocationOnly.CpuCoreUsageMax,
		RamByteUsageMax:            report.RawAllocationOnly.RamByteUsageMax,
	}
}
