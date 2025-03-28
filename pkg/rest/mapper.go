package rest

import (
	"kartverket/skip/opencost/pkg/database"
	"maps"
	"slices"
)

// used for container,namespace aggregation
func mapDatabaseReportsToNamespaceCosts(reports []database.Report, daysBetween int) []NamespaceCost {
	var namespaces = make(map[string]NamespaceCost)
	for _, report := range reports {
		namespace, exists := namespaces[report.Namespace]
		if !exists {
			namespace = NamespaceCost{
				Name:            report.Namespace,
				Team:            report.TeamLabel,
				Division:        report.DivisionLabel,
				Labels:          report.NamespaceLabels,
				Containers:      map[string]ContainerCost{},
				CpuCost:         0.0,
				MemoryCost:      0.0,
				PVCost:          0.0,
				TotalCost:       0.0,
				TotalEfficiency: 0.0,
			}
		}

		if namespace.Name == "idle" {
			report.TotalCost *= -1.0
			report.PvCost *= -1.0
			report.CpuCost *= -1.0
			report.RamCost *= -1.0
		}

		namespace.CpuCost += report.CpuCost
		namespace.MemoryCost += report.RamCost
		namespace.PVCost += report.PvCost
		namespace.TotalCost += report.TotalCost

		container, exists := namespace.Containers[report.Container]
		if !exists {
			container = ContainerCost{
				Name:            report.Container,
				CpuCost:         0.0,
				MemoryCost:      0.0,
				PVCost:          0.0,
				TotalCost:       0.0,
				TotalEfficiency: 0.0,
			}
		}

		container.CpuCost += report.CpuCost
		container.MemoryCost += report.RamCost
		container.PVCost += report.PvCost
		container.TotalCost += report.TotalCost
		container.TotalEfficiency += report.TotalEfficiency * 100 / float64(daysBetween)

		namespace.Containers[report.Container] = container
		namespace.TotalEfficiency = getNamespaceEfficiency(namespace)
		namespaces[report.Namespace] = namespace
	}

	var namespaceCosts []NamespaceCost = slices.Collect(maps.Values(namespaces))

	if namespaceCosts == nil {
		namespaceCosts = make([]NamespaceCost, 0)
	}

	return namespaceCosts
}

func getNamespaceEfficiency(namespace NamespaceCost) float64 {
	var totalEfficiency float64 = 0.0
	for _, container := range namespace.Containers {
		totalEfficiency += container.TotalEfficiency
	}
	totalEfficiency /= float64(len(namespace.Containers))
	return totalEfficiency
}
