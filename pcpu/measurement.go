package pcpu

import (
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/process"
)

var numCPU int = runtime.NumCPU()

type Measurement struct {
	Process         *process.Process
	MeasurementTime time.Time
	Times           cpu.TimesStat
	CPUPercent      float64
}

func NewMeasurement(p *process.Process) (*Measurement, error) {
	m := Measurement{
		Process:         p,
		MeasurementTime: time.Now(),
	}

	cpuPercent, err := p.CPUPercent()
	if err == nil {
		m.CPUPercent = cpuPercent
	} else {
		return nil, err
	}

	cpuTime, err := p.Times()
	if err == nil {
		m.Times = *cpuTime
	} else {
		return nil, err
	}

	return &m, nil
}

func (m *Measurement) TotalTime() float64 {
	total := m.Times.User + m.Times.System + m.Times.Nice + m.Times.Iowait + m.Times.Irq + m.Times.Softirq + m.Times.Steal + m.Times.Idle
	return total
}

func (m *Measurement) ActiveTime() float64 {
	return m.TotalTime() - m.Times.Idle
}

func (m *Measurement) ProcessName() string {
	name, err := m.Process.Name()
	if err == nil {
		return name
	}
	return "unknown"
}

func (m *Measurement) Dump() {
	fmt.Printf("comm=%s, pid=%d, utime=%f, cpu_usage=%f\n", m.ProcessName(), m.Process.Pid, m.Times.User, m.CPUPercent)
}
