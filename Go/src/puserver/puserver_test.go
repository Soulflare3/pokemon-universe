package main

import (
	"testing"
)

func TestTimeService(t *testing.T) {
	timeService := NewTimeService()
	timeService.calculateTimeFromSystem()
	
	t.Logf("INFO: Current time is %v @ %d:%d\n", Days[timeService.Day], timeService.Hour, timeService.Minutes)
}