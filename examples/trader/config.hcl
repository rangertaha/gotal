# Provider
provider "polygon" {
    name = "provider1"
    api_key = "1234567890"
    api_secret = "1234567890"
    api_passphrase = "1234567890"
}

# Indicator
indicator "macd" {
    short_period = 12
    long_period = 26
    signal_period = 9
}

# Broker
broker "coinbase" { 
    name = "broker1"
    api_key = "1234567890"
    api_secret = "1234567890"
    api_passphrase = "1234567890"
    api_url = "https://api.coinbase.com/v2"
}

# Strategy
strategy "macd" {
    name = "macd01"
}

# Storage
storage "influxdb" {
    name = "influxdb01"
    url = "http://localhost:8086"
    token = "1234567890"
    org = "my_org"
    bucket = "my_bucket"
}
