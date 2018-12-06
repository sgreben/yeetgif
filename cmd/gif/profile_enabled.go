// +build profile

package main

import (
	"log"
	"os"
	"runtime/pprof"
)

var (
	cpuProfile *os.File
	memProfile *os.File
)

func startProfile() {
	var err error
	cpuProfile, err := os.Create("yeetgif-cpu.profile")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(cpuProfile)
	memProfile, err := os.Create("yeetgif-mem.profile")
	if err != nil {
		log.Fatal(err)
	}
	pprof.WriteHeapProfile(memProfile)
}

func stopProfile() {
	pprof.StopCPUProfile()
	cpuProfile.Close()
	memProfile.Close()
}
