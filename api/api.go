package api

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"net/http"
)

const baseURL = "http://capybara.development/"
const baseToken = "f30a3568ce54f2f40c13fae2197f0ec585ab15c0"

type CapybaraService struct {
	sling *sling.Sling
}

type CapybaraQuery struct {
	Token string `url:"token"`
}

type CapybaraBody struct {
	Host    *host.InfoStat         `json:"host"`
	Cpu     []cpu.InfoStat         `json:"cpu"`
	CpuTime []cpu.TimesStat        `json:"cpuTime"`
	Memory  *mem.VirtualMemoryStat `json:"memory"`
	Disk    *disk.UsageStat        `json:"disk"`
	Load    *load.AvgStat          `json:"load"`
}

type CapybaraResponse struct {
	Error bool `json:"error"`
}

func NewCapybaraService() *CapybaraService {
	queryParams := NewCapybaraQuery()

	return &CapybaraService{
		sling: sling.New().Base(baseURL).QueryStruct(queryParams).Add("Content-Type", "application/json").Add("Accept", "application/json"),
	}
}

func NewCapybaraQuery() *CapybaraQuery {
	return &CapybaraQuery{
		Token: baseToken,
	}
}

func SendEvent(body CapybaraBody) (*CapybaraResponse, *http.Response, error) {
	service := NewCapybaraService()
	cResponse := new(CapybaraResponse)

	resp, err := service.sling.Post("/api/event").BodyJSON(body).ReceiveSuccess(cResponse)
	if err != nil {
		panic(err)
	}

	fmt.Println(cResponse, resp, err)

	return cResponse, resp, err
}
