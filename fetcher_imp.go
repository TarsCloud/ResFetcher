package main

import (
	"context"

	"github.com/TarsCloud/ResFetcher/TarsTestToolKit"
	"github.com/TarsCloud/ResFetcher/services/cpu"
	"github.com/TarsCloud/ResFetcher/services/memory"
)

// fetcherImp servant implementation
type fetcherImp struct {
}

// Init servant init
func (imp *fetcherImp) Init() error {
	//initialize servant here:
	//...
	return nil
}

// Destroy servant destory
func (imp *fetcherImp) Destroy() {
	//destroy servant here:
	//...
}

// FetchResInfo 获取系统资源
func (imp *fetcherImp) FetchResInfo(tarsCtx context.Context) (ret TarsTestToolKit.ResResp, err error) {
	memInfo, err := memory.Stat()
	if err != nil {
		return ret, err
	}

	cores, err := cpu.Stat()
	if err != nil {
		return ret, err
	}

	ret.MemInfo = *memInfo
	ret.Cores = cores

	return ret, nil
}
