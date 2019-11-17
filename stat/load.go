package stat

import "github.com/shirou/gopsutil/load"

type Load struct {
	Average struct {
		Load1  float64 `json:"load1"`
		Load5  float64 `json:"load5"`
		Load15 float64 `json:"load15"`
	}

	Processes struct {
		ProcsTotal   int `json:"procsTotal"`
		ProcsRunning int `json:"procsRunning"`
		ProcsBlocked int `json:"procsBlocked"`
		Ctxt         int `json:"ctxt"`
	}
}

func NewLoad() Load {
	l := load.Avg()
	return Load{
		Average: struct {
			Load1  float64 `json:"load1"`
			Load5  float64 `json:"load5"`
			Load15 float64 `json:"load15"`
		}{},
		Processes: struct {
			ProcsTotal   int `json:"procsTotal"`
			ProcsRunning int `json:"procsRunning"`
			ProcsBlocked int `json:"procsBlocked"`
			Ctxt         int `json:"ctxt"`
		}{},
	}
}
