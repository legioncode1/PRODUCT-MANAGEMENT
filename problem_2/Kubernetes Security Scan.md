
# Kubernetes Cluster Security Scan with Kubescape

Scanning all resources in your local Kubernetes cluster using Kubescape.

## Prerequisites

1. Minikube
2. Kubescape

## Step-by-Step Guide to Scan All Resources

### Step 1: Install Minikube

```
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-arm64
sudo install minikube-darwin-arm64 /usr/local/bin/minikube
```

### Step 2: Install Kubescape

```
curl -s https://raw.githubusercontent.com/kubescape/kubescape/master/install.sh | /bin/bash
```



### Step 3: Start Minikube

Ensure your Minikube cluster is up and running:

```
minikube start
```

### Step 4: Run the Full Kubescape Scan

Use the following command to scan all resources in your Kubernetes cluster:

```
kubescape scan all --kubeconfig ~/.kube/config --format json --output findings.json
```

### Step 5: Review the Findings

Once the scan is complete, open the `findings.json` file to review a detailed list of the security issues found in your cluster.

```
cat findings.json
```

We can now assess and address any security vulnerabilities that were detected.


