package cpu

import (
	"github.com/TarsCloud/ResFetcher/fetcher"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

// Stat 获取cpu stat信息
func Stat() ([]fetcher.CoreInfo, error) {
	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		return nil, err
	}

	cpus := make([]fetcher.CoreInfo, 0)
	for _, s := range stat.CPUStats {
		cpus = append(cpus, fetcher.CoreInfo{
			Total: int64(s.User + s.Nice + s.System + s.Idle + s.IOWait + s.IRQ + s.SoftIRQ + s.Steal + s.Guest + s.GuestNice),
			Idle:  int64(s.Idle),
		})
	}

	return cpus, nil
}
