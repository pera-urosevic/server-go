package main

import (
	"os"
	"server/api"
	"server/charged"
	"server/env"
	"server/micron"
	"server/system"
)

func setup() {
	env.Load()
	system.Log("SERVER v5")
}

func main() {
	setup()

	startApi := os.Getenv("ENABLED_API")
	if startApi == "true" {
		go api.Start()
	}

	startMicron := os.Getenv("ENABLED_MICRON")
	if startMicron == "true" {
		go micron.Start()
	}

	startCharged := os.Getenv("ENABLED_CHARGED")
	if startCharged == "true" {
		go charged.Start()
	}

	select {}
}
