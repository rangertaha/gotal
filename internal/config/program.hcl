# Example program configurations for HXE

program "web-server-nginx" {
  id          = 1
  description = "Nginx web server"
  exec        = "nginx -g 'daemon off;'"
  pwd   = "/var/www"
  user        = "www-data"
  group       = "www-data"
  autostart   = true
  enabled     = true
  retries     = 3
}

program "api-server-go" {
  id          = 2
  description = "Go API server"
  exec        = "go run main.go"
  pwd   = "/opt/api"
  user        = "api"
  group       = "api"
  autostart   = true
  enabled     = true
  retries     = 5
}

program "database-postgres" {
  id          = 3
  description = "PostgreSQL database"
  exec        = "postgres -D /var/lib/postgresql/data"
  pwd   = "/var/lib/postgresql"
  user        = "postgres"
  group       = "postgres"
  autostart   = true
  enabled     = true
  retries     = 3
}

program "monitoring-prometheus" {
  id          = 4
  description = "Prometheus monitoring"
  exec        = "prometheus --config.file=/etc/prometheus/prometheus.yml"
  pwd   = "/opt/prometheus"
  user        = "prometheus"
  group       = "prometheus"
  autostart   = false
  enabled     = true
  retries     = 2
} 
