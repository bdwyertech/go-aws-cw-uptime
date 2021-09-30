//go:build windows

package main

import (
	"time"

	"golang.org/x/sys/windows"
)

func GetUptime() time.Duration {
	return windows.DurationSinceBoot()
}
