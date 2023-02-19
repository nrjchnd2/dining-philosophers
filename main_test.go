package main

import (
	"testing"
	"time"
)

func Test_Dine(t *testing.T) {
	eatingTime = 0 * time.Second
	thinkingTime = 0 * time.Second
	for i := 0; i < 10; i++ {
		orderDiningFinished = []string{}
		dine()
		if len(orderDiningFinished) != 5 {
			t.Errorf("Unexpected slice length received. Expected 5 but got %d", len(orderDiningFinished))
		}
	}
}

func Test_DineWithDelay(t *testing.T) {
	var delays = []struct {
		name  string
		delay time.Duration
	}{
		{name: "zero delay", delay: 0 * time.Second},
		{name: "quater second delay", delay: 250 * time.Millisecond},
		{name: "half second delay", delay: 500 * time.Millisecond},
	}

	for _, e := range delays {
		eatingTime = e.delay
		thinkingTime = e.delay
		orderDiningFinished = []string{}
		dine()
		if len(orderDiningFinished) != 5 {
			t.Errorf("Unexpected slice length received. Expected 5 but got %d", len(orderDiningFinished))
		}
	}
}
