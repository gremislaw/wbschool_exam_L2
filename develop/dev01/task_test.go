package main

import (
	"testing"
	"time"
)

func TestGetTime_Success(t *testing.T) {
	actualTime, err := GetTime()
	if err != nil {
		t.Fatalf("ошибка, получено %v", err)
	}

	now := time.Now()
	if actualTime.Before(now.Add(-time.Hour)) || actualTime.After(now.Add(time.Hour)) {
		t.Errorf("ошибка, ожидалось верное время как у нас, получено %v", actualTime)
	}
}
