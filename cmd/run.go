package cmd

import (
	"capybara-heartbeat/api"
	"github.com/shirou/gopsutil/cpu"
	_ "github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	_ "github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	_ "github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	_ "github.com/shirou/gopsutil/mem"
	"github.com/urfave/cli/v2"
	"github.com/whiteShtef/clockwork"
)

func Run(c *cli.Context) error {
	cw := clockwork.NewScheduler()
	cw.Schedule().Every(5).Seconds().Do(process)

	cw.Run()

	return nil
}

func process() {
	h, _ := host.Info()
	c, _ := cpu.Info()
	ct, _ := cpu.Times(false)
	m, _ := mem.VirtualMemory()
	d, _ := disk.Usage("/")
	l, _ := load.Avg()

	body := api.CapybaraBody{
		Host:    h,
		Cpu:     c,
		CpuTime: ct,
		Memory:  m,
		Disk:    d,
		Load:    l,
	}

	_, _, _ = api.SendEvent(body)
}
