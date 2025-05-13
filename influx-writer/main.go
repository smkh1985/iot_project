package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	influx "github.com/influxdata/influxdb1-client/v2"
	"github.com/segmentio/kafka-go"
)

// SensorData Kafka
type SensorData struct {
	SensorID  string    `json:"sensor_id"`
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	Value     float64   `json:"value"`
}

func main() {
	// Connect to InfluxDB
	influxClient, err := influx.NewHTTPClient(influx.HTTPConfig{
		Addr:     "http://influxdb:8086",
		Username: "admin",
		Password: "admin123",
	})
	if err != nil {
		log.Fatal("Error in Connecting to InfluxDB:", err)
	}
	
	_, err = influxClient.Query(influx.NewQuery("CREATE DATABASE sensors", "", ""))
	if err != nil {
		log.Println("❌ Error creating InfluxDB database:", err)
	} else {
		fmt.Println("✅ Ensured InfluxDB database exists")
	}
	defer influxClient.Close()

	// Connect to Kafka
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "iot-sensors",
		GroupID: "influx-writer-group",
	})
	defer reader.Close()

	fmt.Println("✅ influx-writer is ready to read from Kafka and write in InfluxDB")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("❌ Error in reading message", err)
			continue
		}

		var data SensorData
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			log.Println("❌ Error in JSON:", err)
			continue
		}

		// make point
		point, err := influx.NewPoint(
			"sensor_data", // Measurement
			map[string]string{
				"sensor_id": data.SensorID,
				"type":      data.Type,
			},
			map[string]interface{}{
				"value": data.Value,
			},
			data.Timestamp,
		)
		if err != nil {
			log.Println("❌ Error in making point", err)
			continue
		}

		// Send to Influx
		batch, _ := influx.NewBatchPoints(influx.BatchPointsConfig{
			Database:  "sensors",
			Precision: "s",
		})
		batch.AddPoint(point)

		if err := influxClient.Write(batch); err != nil {
			log.Println("❌ InfluxDB Error in writing :", err)
		} else {
			fmt.Println("✅ Saved :", data)
		}
	}
}

