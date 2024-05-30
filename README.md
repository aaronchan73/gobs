# Gobs
Lightweight and minimalist observability for distributed systems.

## Library (`/lib`)
APIs that cover the three pillars of observability!

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

## Collector (`/collector`)
HTTP server for data ingestion from clients.

### Running on Docker
1. Build container: `docker build -f Dockerfile.collector -t gobs-collector .`
2. Run container: `docker run -d -p 8080:8080 gobs-collector`

### Endpoints
- `/logs` - update Logs from given JSON
- `/counters` - update Counters from given JSON
- `/gauges` - update Gauges from given JSON
- `/histograms` - update Histograms from given JSON
- `/traces` - update Traces from given JSON

## Testing (`/main`)
Manual tests done inside simulated environment (K8s) and using Postman.

### Architecture
![arch](./assets/arch.png)

### Running on Kubernetes
1. Install minikube and Docker
2. Load containers into minikube \
   a. Start minikube: `minikube start` \
   b. Load images: `eval $(minikube docker-env)`
4. Build containers for collector and main \
   a. Build collector's container: `docker build -f Dockerfile.collector -t gobs-collector .` \
   b. Build main's container: `docker build -f Dockerfile.main -t gobs-main .` \
   c. Verify using `minikube ssh` and `docker images`
5. Apply deployment and service files \
   a. Apply collector deployment: `kubectl apply -f k8s/collector-deployment.yaml` \
   b. Apply collector service: `kubectl apply -f k8s/collector-service.yaml` \
   c. Apply main deployment: `kubectl apply -f k8s/main-deployment.yaml` \
   d. Apply main service: `kubectl apply -f k8s/main-service.yaml` \
   e. Apply Grafana deployment: `kubectl apply -f k8s/grafana-deployment.yaml` \
   f. Apply Grafana service: `kubectl apply -f k8s/grafana-service.yaml`
7. View Grafana dashboards \
   a. Port-forward dashboards: `kubectl port-forward deployment/grafana 3000:3000`

## Resources
- https://www.splunk.com/en_us/blog/learn/observability.html
- https://www.datadoghq.com/three-pillars-of-observability/
