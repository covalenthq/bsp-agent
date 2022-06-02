# Metrics collection and reporting

bsp-agent enabled metrics collection via `--metrics` flag. A metrics server can be enabled (by --metrics.addr and --metrics.port). The metrics are served in two formats:
- `/debug/metrics`: json representation of expvars and go-metrics
- `/debug/metrics/prometheus`: same metrics as above but in prometheus format


Monitoring can be setup (for example) by plugging the endpoint serving in prometheus-format into influxdb, which is plugged into grafana.s