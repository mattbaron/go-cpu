package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/mattbaron/go-cpu/pcpu"
	"github.com/shirou/gopsutil/v3/process"
)

func Test(s1 string, s2 string) {
	i1, err1 := strconv.ParseInt(s1, 10, 32)
	if err1 == nil {
		fmt.Printf("i1=%d\n", i1)
	}

	i2, err2 := strconv.ParseInt(s2, 10, 32)
	if err2 == nil {
		fmt.Printf("i2=%d\n", i2)
	}

	if math.Abs(float64(i2)-float64(i1)) < 2 {
		fmt.Println("COOL")
	} else {
		fmt.Println("NOT COOL")
	}
}

func main() {
	pid64, err := strconv.ParseInt(os.Args[1], 10, 32)
	pid := int32(pid64)
	if err != nil {
		fmt.Printf("Invalid PID: %s\n", os.Args[1])
		os.Exit(2)
	}

	process, err := process.NewProcess(pid)

	if err != nil {
		fmt.Printf("Invalid PID: %d\n", pid)
		os.Exit(2)
	}

	//fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())

	collector := pcpu.NewCollector(process)
	collector.Collect()
	time.Sleep(time.Second * 5)
	collector.Collect()

	measurement, err := collector.LastMeasurement()
	if err == nil {
		fmt.Printf("Last CPUPercent() %v\n", measurement.CPUPercent)
	}

	cpuPercentInterval, err := collector.CPUPercentInterval()
	if err == nil {
		fmt.Printf("CPUPercentInterval() %v\n", cpuPercentInterval)
	}
}
