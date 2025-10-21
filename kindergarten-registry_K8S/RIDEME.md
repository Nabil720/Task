
<h1 style="color: #1bf89cff; font-size: 48px; font-weight: bold;">K8s Cluster based hosting</h1>


## Step-by-Step Deployment Guide

### Launch EC2 Instance
1. Create an EC2 instance using **Ubuntu 22.04 (t3.micro)**.  
2.  Connect via SSH:
   ```bash
   ssh -i key.pem ubuntu@<EC2_PUBLIC_IP>
```


## Install Tools

```
# update
sudo apt update -y
sudo apt upgrade -y

# Install jq, curl, unzip 
sudo apt install -y curl unzip jq

# Install AWS CLI v2
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
rm -rf awscliv2.zip aws

# Install kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
rm kubectl

# Install eksctl (the easiest way to create EKS)
curl --silent --location "https://github.com/weaveworks/eksctl/releases/latest/download/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp
sudo mv /tmp/eksctl /usr/local/bin

# (Optional) Install helm if you want to use it later
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash

```
## AWS CLI credentials
```
aws configure
# AWS Access Key ID [None]: <YOUR_ACCESS_KEY>
# AWS Secret Access Key [None]: <YOUR_SECRET_KEY>
# Default region name [None]: us-east-1   
# Default output format [None]: json
```
## EKS Cluster
```
eksctl create cluster \
  --name kindergarten-cluster \
  --region us-east-1 \
  --nodegroup-name ng-standard \
  --node-type t3.small \
  --nodes 2 \
  --nodes-min 1 \
  --nodes-max 3 \
  --managed

# Chack
kubectl get nodes
kubectl get ns
```

## Nginx ingress controller install
```
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm install nginx-ingress ingress-nginx/ingress-nginx \
  --namespace ingress-nginx --create-namespace \
  --set controller.publishService.enabled=true

# chack
kubectl get svc -n ingress-nginx

```

###  Clone the Repository
```
cd ~
git clone https://github.com/Nabil720/Task.git

cd Task/kindergarten-registry_K8S/K8s

# Apply all YAML files
kubectl apply -f .

# Chack 

kubectl get all

```






##  Applicatio architecture

```
                                   +-----------------------+
                                   |    Web Browser        |
                                   |  (React Frontend UI)  |
                                   +-----------+-----------+
                                               |
                                               v
                                   +-----------------------+
                                   |     Nginx Ingress     |
                                   |  (Routing & Proxy)    |
                                   +-----------+-----------+
                                               |
                                               v
                     +-----------------------------------------------+
                     |     EKS Cluster (Kubernetes Environment)     |
                     |                                               |
                     |  +-------------------+   +----------------+  |
                     |  |   React Frontend   |   |   Backend      |  |
                     |  |  (Pod - Static UI) |   |  (Go/Express   |  |
                     |  |  (Nginx Serve)     |   |   API Server)  |  |
                     |  +-------------------+   +----------------+  |
                     |             |                  |            |
                     |             v                  v            |
                     |  +--------------------------+-------------+  |
                     |  |   MongoDB (Pod)           |             |  |
                     |  |   (Persistent Volume)     |             |  |
                     |  +--------------------------+-------------+  |
                     +-----------------------------------------------+
                                               |
                                               v
                                    +------------------------+
                                    |    AWS EBS (Storage)   |
                                    |  (Persistent Volumes)  |
                                    +------------------------+


```




![Website View](./Images/Screenshot%20from%202025-10-14%2011-34-55.png)

# After Added Probe
![Website View](./Images/Screenshot%20from%202025-10-21%2016-14-38.png)
![Cluster View](./Images/Screenshot%20from%202025-10-21%2016-15-31.png)