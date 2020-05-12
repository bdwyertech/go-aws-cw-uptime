// +build !windows

package main

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/shirou/gopsutil/host"
)

func GetUptime() time.Duration {
	// Uptime in Seconds
	uptime, err := host.Uptime()
	if err != nil {
		log.Fatal(err)
	}

	return time.Duration(uptime) * time.Second
}
