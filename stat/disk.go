package stat

import "github.com/shirou/gopsutil/disk"

type Disk struct {
	UsageStat struct {
		Path              string  `json:"path"`
		Fstype            string  `json:"fstype"`
		Total             uint64  `json:"total"`
		Free              uint64  `json:"free"`
		Used              uint64  `json:"used"`
		UsedPercent       float64 `json:"usedPercent"`
		InodesTotal       uint64  `json:"inodesTotal"`
		InodesUsed        uint64  `json:"inodesUsed"`
		InodesFree        uint64  `json:"inodesFree"`
		InodesUsedPercent float64 `json:"inodesUsedPercent"`
	}

	PartitionStat struct {
		Device     string `json:"device"`
		Mountpoint string `json:"mountpoint"`
		Fstype     string `json:"fstype"`
		Opts       string `json:"opts"`
	}

	IOCountersStat struct {
		ReadCount        uint64 `json:"readCount"`
		MergedReadCount  uint64 `json:"mergedReadCount"`
		WriteCount       uint64 `json:"writeCount"`
		MergedWriteCount uint64 `json:"mergedWriteCount"`
		ReadBytes        uint64 `json:"readBytes"`
		WriteBytes       uint64 `json:"writeBytes"`
		ReadTime         uint64 `json:"readTime"`
		WriteTime        uint64 `json:"writeTime"`
		IopsInProgress   uint64 `json:"iopsInProgress"`
		IoTime           uint64 `json:"ioTime"`
		WeightedIO       uint64 `json:"weightedIO"`
		Name             string `json:"name"`
		SerialNumber     string `json:"serialNumber"`
		Label            string `json:"label"`
	}
}

func NewDisk() Disk {
	d := disk.Usage("/")

	return Disk{
		UsageStat: struct {
			Path              string  `json:"path"`
			Fstype            string  `json:"fstype"`
			Total             uint64  `json:"total"`
			Free              uint64  `json:"free"`
			Used              uint64  `json:"used"`
			UsedPercent       float64 `json:"usedPercent"`
			InodesTotal       uint64  `json:"inodesTotal"`
			InodesUsed        uint64  `json:"inodesUsed"`
			InodesFree        uint64  `json:"inodesFree"`
			InodesUsedPercent float64 `json:"inodesUsedPercent"`
		}{},
		PartitionStat: struct {
			Device     string `json:"device"`
			Mountpoint string `json:"mountpoint"`
			Fstype     string `json:"fstype"`
			Opts       string `json:"opts"`
		}{},
		IOCountersStat: struct {
			ReadCount        uint64 `json:"readCount"`
			MergedReadCount  uint64 `json:"mergedReadCount"`
			WriteCount       uint64 `json:"writeCount"`
			MergedWriteCount uint64 `json:"mergedWriteCount"`
			ReadBytes        uint64 `json:"readBytes"`
			WriteBytes       uint64 `json:"writeBytes"`
			ReadTime         uint64 `json:"readTime"`
			WriteTime        uint64 `json:"writeTime"`
			IopsInProgress   uint64 `json:"iopsInProgress"`
			IoTime           uint64 `json:"ioTime"`
			WeightedIO       uint64 `json:"weightedIO"`
			Name             string `json:"name"`
			SerialNumber     string `json:"serialNumber"`
			Label            string `json:"label"`
		}{},
	}
}
