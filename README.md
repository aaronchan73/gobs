# Gobs
Lightweight and minimalist observability for distributed systems.

## Logs
- Found in `logs.go`
- Currently outputs to stdout

## Metrics
- Found in `metrics.go`
- Currently outputs to stdout
- Supports counters, gauges, and histograms

## Traces
- Found in `traces.go`
- Currently outputs to stdout
- Supports spans and traces

## Testing
![arch](https://github.com/aaronchan73/gobs/assets/84891329/9daf11da-8d68-49ce-ae02-8d0815b8f9d1)
- TODO: test communication between pods using Docker/Kubernetes
- TODO: visualize onto Grafana dashboards

## Resources
- https://www.splunk.com/en_us/blog/learn/observability.html
- https://www.datadoghq.com/three-pillars-of-observability/
