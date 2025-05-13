package sensors

import (
	"math/rand"
	"time"
)

type TemperatureSensor struct{
	sensorId string
	interval time.Duration
}

func NewTemperatureSensor(id string, interval time.Duration) *TemperatureSensor{
	return &TemperatureSensor{
		sensorId: id, 
		interval: interval,
	}
}

func (ts *TemperatureSensor) ID () string{
	return ts.sensorId
}

func (ts *TemperatureSensor) Type () string{
	return "temperature"
}

func (ts *TemperatureSensor) Interval () time.Duration{
	return ts.interval
}

func (ts *TemperatureSensor) ReadData() SensorStruct{
	sensor_data := SensorStruct{
		SensorID: ts.ID(),
		Type: ts.Type(),
		Timestamp: time.Now().UTC(),
		Value: 20 + rand.Float64()*10,
	}
	return sensor_data
}

