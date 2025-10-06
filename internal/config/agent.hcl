// Hxe Configuration
//

// Hxe API Server
server "hxe" {
  host = "0.0.0.0"
  port = 3143
} 
service "programs" {
  directory = "programs"
}

// Timeseries Database: (Optional) Timeseries database client connection
// A timeseries database used to store historical metrics
// By default, InfluxDB is used. In the future, other options may be supported
# service "metrics" {
#   name = "influxdb"
#   url = "http://localhost:8086"
#   token = "influxdb"
# }

