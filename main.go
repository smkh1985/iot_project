package main

import (
	"fmt"
	"time"

	"iotApp/config"
	"iotApp/publisher"
	"iotApp/sensors"
)

func main() {
	cfg := config.LoadConfig("config.yaml")
	pub := publisher.NewKafkaPublisher("kafka:9092", "iot-sensors")

	for _, sensorCfg := range cfg.Sensors {
		sensor := createSensor(sensorCfg)
		go func(s sensors.Sensor) {
			for {
				data := s.ReadData()
				if err := pub.Publish(data); err != nil {
					fmt.Println("❌ Publish error:", err)
				} else {
					fmt.Println("✅ Published:", data)
				}
				time.Sleep(s.Interval())
			}
		}(sensor)
	}

	select {} // block forever
}

func createSensor(cfg config.SensorConfig) sensors.Sensor {
	switch cfg.Type {
	case "temperature":
		return sensors.NewTemperatureSensor(cfg.ID, cfg.Interval)
	case "humidity":
		return sensors.NewHumiditySensor(cfg.ID, cfg.Interval)
	default:
		panic("Unsupported sensor type: " + cfg.Type)
	}
}
