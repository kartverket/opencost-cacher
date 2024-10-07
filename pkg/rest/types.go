package rest

type NamespaceCost struct {
	Name            string                   `json:"name"`
	CpuCost         float64                  `json:"cpuCost"`
	MemoryCost      float64                  `json:"memoryCost"`
	PVCost          float64                  `json:"pVCost"`
	TotalCost       float64                  `json:"totalCost"`
	TotalEfficiency float64                  `json:"totalEfficiency"`
	Team            string                   `json:"team"`
	Division        string                   `json:"division"`
	Labels          map[string]string        `json:"labels"`
	Containers      map[string]ContainerCost `json:"containers"`
}

type ContainerCost struct {
	Name            string  `json:"name"`
	CpuCost         float64 `json:"cpuCost"`
	MemoryCost      float64 `json:"memoryCost"`
	PVCost          float64 `json:"pVCost"`
	TotalCost       float64 `json:"totalCost"`
	TotalEfficiency float64 `json:"totalEfficiency"`
}
