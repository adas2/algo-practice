package main

import "testing"

func TestParallalFuncs(t *testing.T) {
	// testChannels()
	rateLimiterMain()
}

func TestConcurrency(t *testing.T) {
	testConcurrency()
}

func TestProducerConsumer(t *testing.T) {
	// pcMain()
	// mainSPMC()
	mainMPMC()
}
