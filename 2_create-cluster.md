
# Create Kubernetes cluster

`⚠️ WARN` Here we assume that local environment is ready:
	- all accounts created and properly configured
	- all tools installed and properly configured

## CIVO

```bash
# Set region
civo region current LON1

# Create cluster
civo kubernetes create wt21 \
	--size=g3.k3s.large \
	--nodes=3 \
	--region=LON1 \
	--version=v1.20.0+k3s1 \
	--wait

# Check cluster
civo kubernetes list
civo kubernetes show wt21

# Connect kubectl to new cluster
civo kubernetes config wt21 --save --merge

# Check kubectl
kubectl config use-context wt21
kubectl get nodes
```

## GCP

```bash
# Set project
export PROJECT_ID=...
gcloud config set project $PROJECT_ID

# Create cluster
gcloud beta container --project "$PROJECT_ID" clusters create "wt21" --zone "europe-west4-a" \
	--release-channel "None" --cluster-version "1.20.8-gke.2100" \
	--node-locations "europe-west4-a" \
	--num-nodes "3" --machine-type "e2-custom-8-16384" --image-type "COS_CONTAINERD" --disk-type "pd-standard" --disk-size "100" \
	--no-enable-autoupgrade --enable-autorepair --max-surge-upgrade 1 --max-unavailable-upgrade 0 --enable-shielded-nodes \
	--max-pods-per-node "110" --default-max-pods-per-node "110" \
	--network "projects/$PROJECT_ID/global/networks/default" \
	--subnetwork "projects/$PROJECT_ID/regions/europe-west4/subnetworks/default" \
	--enable-ip-alias --no-enable-basic-auth --no-enable-intra-node-visibility --no-enable-master-authorized-networks \
	--addons HorizontalPodAutoscaling,HttpLoadBalancing,GcePersistentDiskCsiDriver \
	--scopes "https://www.googleapis.com/auth/devstorage.read_only","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" \
	--metadata disable-legacy-endpoints=true

# Connect kubectl to new cluster
gcloud container clusters get-credentials "wt21" --zone "europe-west4-a" --project "$PROJECT_ID"

# Check kubectl
kubectl config use-context wt21
kubectl get nodes
```

## LOCAL (MINIKUBE)

Please choose the best driver based on your local environment.

```bash
# Start Minikube with Docker driver
minikube start --driver=docker \
	--cpus=8 --memory=12288 --disk-size=50g \
	--kubernetes-version=v1.20.8

# Start Minikube with Hyperkit driver
minikube start --driver=hyperkit \
	--cpus=8 --memory=12288 --disk-size=50g \
	--kubernetes-version=v1.20.8

# Start Minikube with VirtualBox driver
minikube start --driver=virtualbox \
	--cpus=8 --memory=12288 --disk-size=50g \
	--kubernetes-version=v1.20.8

# Check kubectl
kubectl get nodes
```
