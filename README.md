# Gobs
Lightweight and minimalist observability for distributed systems.

## Library
- Found in `/lib`
### Logs
- Found in `logs.go`
- Outputs to collector via HTTP

### Metrics
- Found in `metrics.go`
- Outputs to collector via HTTP
- Supports counters, gauges, and histograms

### Traces
- Found in `traces.go`
- Outputs to collector via HTTP
- Supports spans and traces

## Collector
- HTTP server for data ingestion from clients

## Testing
![arch](./assets/arch.png)
- TODO: replace output to use HTTP/gRPC
- TODO: test communication between pods using Docker/Kubernetes
- TODO: visualize onto Grafana dashboards

## Resources
- https://www.splunk.com/en_us/blog/learn/observability.html
- https://www.datadoghq.com/three-pillars-of-observability/
