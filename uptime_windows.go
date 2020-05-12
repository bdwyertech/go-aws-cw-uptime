// +build windows

package main

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/StackExchange/wmi"
)

// https://docs.microsoft.com/en-us/windows/win32/cimwin32prov/win32-operatingsystem
type Win32_OperatingSystem struct {
	FreePhysicalMemory     uint64
	FreeVirtualMemory      uint64
	NumberOfProcesses      uint64
	TotalVirtualMemorySize uint64
	TotalVisibleMemorySize uint64
	LastBootUpTime         time.Time
}

var LastBootUpTime time.Time

func lastBoot() time.Time {
	// Only retrieve LastBootUpTime once from WMI
	if LastBootUpTime.IsZero() {
		var dst []Win32_OperatingSystem
		q := wmi.CreateQuery(&dst, "")
		if err := wmi.Query(q, &dst); err != nil {
			log.Fatal(err)
		}
		LastBootUpTime = dst[0].LastBootUpTime
	}
	return LastBootUpTime
}

func GetUptime() time.Duration {
	return time.Since(lastBoot())
}
