package pcpu

import (
	"errors"

	"github.com/shirou/gopsutil/v3/process"
)

type Collector struct {
	Process      *process.Process
	Measurements []*Measurement
}

func NewCollector(p *process.Process) *Collector {
	return &Collector{
		Process:      p,
		Measurements: make([]*Measurement, 0),
	}
}

func (c *Collector) Collect() (*Measurement, error) {
	measurement, err := NewMeasurement(c.Process)
	if err == nil {
		c.Measurements = append(c.Measurements, measurement)
		return measurement, nil
	} else {
		return nil, err
	}
}

func (c *Collector) CPUPercentInterval() (float64, error) {
	if len(c.Measurements) < 2 {
		return -1, errors.New("more than 2 measurments required for interval CPU percentage")
	}

	m1, _ := c.FirstMasurement()
	m2, _ := c.LastMeasurement()

	measurementPeriod := (m2.MeasurementTime.Sub(m1.MeasurementTime).Seconds()) * float64(numCPU)
	totalTimeDiff := m2.TotalTime() - m1.TotalTime()
	cpuPercent := ((totalTimeDiff / measurementPeriod) * 100) * float64(numCPU)

	return cpuPercent, nil
}

func (c *Collector) FirstMasurement() (*Measurement, error) {
	if len(c.Measurements) == 0 {
		return nil, errors.New("not enough measurements")
	}

	return c.Measurements[0], nil
}

func (c *Collector) LastMeasurement() (*Measurement, error) {
	if len(c.Measurements) == 0 {
		return nil, errors.New("not enough measurements")
	}

	return c.Measurements[len(c.Measurements)-1], nil
}
