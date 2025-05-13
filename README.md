# 🛰️ IoT Sensor Simulator System

This project simulates multiple IoT sensors (e.g., temperature and humidity) using **Go**, sends their data to **Kafka**, stores it in **InfluxDB**, and visualizes and monitors it via **Grafana** — all containerized using **Docker Compose**.

---

## 🧱 System Architecture

- **Sensor Simulator (Go)** – Simulates multiple sensors concurrently and publishes readings to Kafka.
- **Kafka** – Message broker for real-time data.
- **Kafka Consumer (Go)** – Reads messages from Kafka and writes them to InfluxDB.
- **InfluxDB** – Time-series database for storing sensor data.
- **Grafana** – Visualizes sensor data and supports alerting.
- **Docker Compose** – Orchestrates all components for local or cloud deployment.

![System Diagram](./docs/system-diagram.png)

---

## 🚀 Getting Started

### 📦 Prerequisites

- Docker
- Docker Compose

### 🔧 Build & Run

```bash
docker compose up --build -d
```

This will spin up all services:
- 3 simulated sensors
- Kafka + Zookeeper
- Kafka UI
- InfluxDB
- Grafana
- Kafka → InfluxDB consumer

---

## ⚙️ Configuration

### `.env` (sample)

Each sensor service can be configured via its own `.env` file:

```env
KAFKA_BROKER=kafka:9092
KAFKA_TOPIC=iot-sensors
SENSOR_TYPE=temperature
SENSOR_ID=sensor-1
SENSOR_INTERVAL=2s
```

Multiple `.env` files can simulate different sensors, each as a separate Docker Compose service.

---

## 📊 Grafana Dashboard

- Access: [http://localhost:3000](http://localhost:3000)
- Login: `admin` / `admin`
- Add InfluxDB data source:
  - URL: `http://influxdb:8086`
  - Database: `sensors`
- Create panels using queries like:

```sql
SELECT mean("value") FROM "sensor_data"
WHERE "type" = 'temperature' AND $timeFilter
GROUP BY time($__interval)
```

---

## 🔔 Alerts

You can configure alerts in Grafana based on thresholds (e.g., temperature > 30°C) and route them to email, Slack, or webhook via **Alerting → Notification Policies**.

---

## 📁 Project Structure

```
.
├── docker-compose.yml
├── .env (per sensor)
├── sensors/           # Go code for sensor simulator
├── influx-writer/     # Go code for Kafka-to-Influx consumer
├── Dockerfile         # Multi-stage build for Go apps
├── init-sensors.iql   # Creates DB if needed
└── docs/
    └── system-diagram.png
```

---

## 🛠 Technologies Used

- Go 1.21+
- Docker & Docker Compose
- Kafka + Zookeeper
- InfluxDB 1.8
- Grafana
- Kafka UI

---

## 👨‍💻 Author

Built by [Majid Khorashadi](https://github.com/your-github)

---

## 📝 License

MIT License. See `LICENSE` file for details.