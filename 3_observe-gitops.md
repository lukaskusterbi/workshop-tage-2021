
# Observe with GitOps

## Infrastructure

```bash
cd infrastructure/
```

### 1. ArgoCD for GitOps

```bash
# Deploy ArgoCD to start with GitOps
kustomize build 10_devops/10_argocd/ | kubectl apply -f -

# Deploy ArgoCD defaults
kustomize build 10_devops/11_defaults/ | kubectl apply -f -

# Check pods
watch 'kubectl get pods -n argocd'
```

### 2. ArgoCD UI

```bash
# Leave port-forwarding open, we will use it a lot
kubectl port-forward svc/argocd-server 6443:443 -n argocd

# Open another terminal

# Choose the right command based on your OS
## linux
kubectl get secret argocd-initial-admin-secret -n argocd -o json | jq -r '.data.password' | base64 -d
## macos
kubectl get secret argocd-initial-admin-secret -n argocd -o json | jq -r '.data.password' | base64 -D

# Open ArgoCD UI
open https://localhost:6443/
```

### 3. Namespaces

```bash
# Create required Namespaces
kubectl apply -f 100_gitops/1_namespaces.yaml

# Check in ArgoCD UI

# Check namespaces
kubectl get namespaces
```

### 4. Monitoring

```bash
# Deploy monitoring
kubectl apply -f 100_gitops/2_monitoring.yaml

# Check in ArgoCD UI

# Check pods
watch 'kubectl get pods -n monitoring'
```

### 5. Logging

```bash
# Deploy logging
kubectl apply -f 100_gitops/3_logging.yaml

# Check in ArgoCD UI

# Check pods
watch 'kubectl get pods -n logging'
```

### 6. Tracing

```bash
# Deploy tracing
kubectl apply -f 100_gitops/4_tracing.yaml

# Check in ArgoCD UI

# Check pods
watch 'kubectl get pods -n tracing'
```

### 7. Dashboards

```bash
# Deploy dashboards
kubectl apply -f 100_gitops/5_dashboards.yaml

# Check in ArgoCD UI

# Check pods
watch 'kubectl get pods -n dashboards'
```

### 7. Databases

```bash
# Deploy databases
kubectl apply -f 100_gitops/6_databases.yaml

# Check in ArgoCD UI

# Check pods
watch 'kubectl get pods -n apps'
```

### 8. Brokers

```bash
# Deploy brokers
kubectl apply -f 100_gitops/7_brokers.yaml

# Check in ArgoCD UI

# Check pods
watch 'kubectl get pods -n apps'
```

### 9. Port forwarding

```bash
# Prometheus
kubectl port-forward svc/prometheus 9090 -n monitoring

# Check everyting is properly configured (targets and some metrics like "kube_pod_info")

# Jaeger
kubectl port-forward svc/jaeger-query 16686 -n tracing

# Check is running

# Leave Grafana port-forwarding open, we will use it a lot
kubectl port-forward svc/grafana 3000 -n dashboards

# Check everyting is properly configured (data sources and dashboards)
```

---

## Applications

### 1. Standalone application

```bash
# Deploy
kubectl apply -f 100_gitops/8_apps.yaml

# Check pods
watch 'kubectl get pods -n apps'

# Check logs
kubectl logs -l app=standalone -f -n apps
```

### 2. gRPC applications

```bash
# Edit 100_gitops/8_apps/kustomization.yaml
open 100_gitops/8_apps/kustomization.yaml

# Uncomment line 8

# Deploy
kubectl apply -f 100_gitops/8_apps.yaml

# Check pods
watch 'kubectl get pods -n apps'

# Check logs
kubectl logs -l app=grpc-server -f -n apps
kubectl logs -l app=grpc-client -f -n apps
```

### 3. HTTP applications

```bash
# Edit 100_gitops/8_apps/kustomization.yaml
open 100_gitops/8_apps/kustomization.yaml

# Uncomment line 9

# Deploy
kubectl apply -f 100_gitops/8_apps.yaml

# Check pods
watch 'kubectl get pods -n apps'

# Check logs
kubectl logs -l app=http-server -f -n apps
kubectl logs -l app=http-client -f -n apps

# Port forward
kubectl port-forward svc/http-client 8080 -n apps

# Now MANUAL make some requests using Postman
```

### 4. Broker applications

```bash
# Edit 100_gitops/8_apps/kustomization.yaml
open 100_gitops/8_apps/kustomization.yaml

# Uncomment line 10

# Deploy
kubectl apply -f 100_gitops/8_apps.yaml

# Check pods
watch 'kubectl get pods -n apps'

# Check logs
kubectl logs -l app=kafka-producer -f -n apps
kubectl logs -l app=kafka-consumer -f -n apps
```
