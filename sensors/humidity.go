package sensors

import (
	"math/rand"
	"time"
)

type HumiditySensor struct{
	sensorId string
	interval time.Duration
}

func NewHumiditySensor(id string, interval time.Duration) *HumiditySensor{
	return &HumiditySensor{
		sensorId: id, 
		interval: interval,
	}
}

func (ts *HumiditySensor) ID () string{
	return ts.sensorId
}

func (ts *HumiditySensor) Type () string{
	return "humidity"
}

func (ts *HumiditySensor) Interval () time.Duration{
	return ts.interval
}

func (ts *HumiditySensor) ReadData() SensorStruct{
	sensor_data := SensorStruct{
		SensorID: ts.ID(),
		Type: ts.Type(),
		Timestamp: time.Now().UTC(),
		Value: 30 + rand.Float64()*50,
	}
	return sensor_data
}
