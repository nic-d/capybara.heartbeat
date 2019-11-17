package stat

import "github.com/shirou/gopsutil/cpu"

type Cpu struct {
	Time struct {
		CPU       string  `json:"cpu"`
		User      float64 `json:"user"`
		System    float64 `json:"system"`
		Idle      float64 `json:"idle"`
		Nice      float64 `json:"nice"`
		Iowait    float64 `json:"iowait"`
		Irq       float64 `json:"irq"`
		Softirq   float64 `json:"softirq"`
		Steal     float64 `json:"steal"`
		Guest     float64 `json:"guest"`
		GuestNice float64 `json:"guestNice"`
	}

	Info struct {
		CPU        int32    `json:"cpu"`
		VendorID   string   `json:"vendorId"`
		Family     string   `json:"family"`
		Model      string   `json:"model"`
		Stepping   int32    `json:"stepping"`
		PhysicalID string   `json:"physicalId"`
		CoreID     string   `json:"coreId"`
		Cores      int32    `json:"cores"`
		ModelName  string   `json:"modelName"`
		Mhz        float64  `json:"mhz"`
		CacheSize  int32    `json:"cacheSize"`
		Flags      []string `json:"flags"`
		Microcode  string   `json:"microcode"`
	}
}

func NewCpu() Cpu {
	c, _ := cpu.Info()

	return Cpu{
		CPU:        0,
		VendorID:   "",
		Family:     "",
		Model:      "",
		Stepping:   0,
		PhysicalID: "",
		CoreID:     "",
		Cores:      0,
		ModelName:  "",
		Mhz:        0,
		CacheSize:  0,
		Flags:      nil,
		Microcode:  nil,
	}
}
