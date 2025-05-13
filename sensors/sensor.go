package sensors

import "time"

type Sensor interface{
	ID() string
	ReadData() SensorStruct
	Type() string
	Interval() time.Duration
}


type SensorStruct struct{
	SensorID	string
	Type	string
	Value	float64
	Timestamp time.Time
}