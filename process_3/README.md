
# Deployment of Go-App on Amazon Elastic Kubernetes Service (EKS)

## Step 1 : Prerequisites

Before setting up and running this project, ensure you have the following tools installed:

### 1. `kubectl` – Kubernetes Command Line Tool
`kubectl` is a command line tool for working with Kubernetes clusters. Follow the instructions to install or update `kubectl` based on your operating system.

#### Installation
You can install `kubectl` using the following methods:

- **macOS**:  
  ```bash
  brew install kubectl
  ```
  
- **Linux**:  
  ```bash
  curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
  chmod +x ./kubectl
  sudo mv ./kubectl /usr/local/bin/kubectl
  ```

- **Windows**:  
  You can use `choco`:
  ```powershell
  choco install kubernetes-cli
  ```

For detailed instructions, refer to the official Kubernetes documentation:  
[Installing or updating kubectl](https://kubernetes.io/docs/tasks/tools/).

### 2. `eksctl` – EKS Command Line Tool
`eksctl` is a command line tool that simplifies creating and managing Amazon EKS clusters.

#### Installation
- **macOS/Linux**:  
  ```bash
  brew tap weaveworks/tap
  brew install weaveworks/tap/eksctl
  ```

- **Windows**:  
  ```powershell
  choco install eksctl
  ```

You can find more installation methods in the official `eksctl` documentation:  
[Installing or updating eksctl](https://eksctl.io/introduction/#installation).

### 3. `AWS CLI` – Amazon Web Services Command Line Interface
The AWS CLI is a tool for interacting with AWS services, including Amazon EKS. You’ll need to configure it to interact with your AWS account.

#### Installation
Follow the instructions to install the AWS CLI:
- **macOS/Linux**:  
  ```bash
  curl "https://awscli.amazonaws.com/AWSCLIV2.pkg" -o "AWSCLIV2.pkg"
  sudo installer -pkg AWSCLIV2.pkg -target /
  ```

- **Windows**:  
  Download the AWS CLI MSI installer from:  
  [AWS CLI Installer for Windows](https://awscli.amazonaws.com/AWSCLIV2.msi)

After installing, verify the installation:
```bash
aws --version
```

#### Configuration
After installation, configure the AWS CLI with your credentials:
```bash
aws configure
```

For detailed steps, refer to the AWS CLI User Guide:  
[Installing, updating, and uninstalling the AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html).


#Step 2: Install EKS



## Install using Fargate

AWS Fargate is a serverless, pay-as-you-go compute engine that lets you focus on building applications without managing servers.
```
eksctl create cluster --name $cluster_name --region us-east-1 --fargate
```

##  IAM OIDC provider configuration

IAM OIDC identity providers are entities in IAM that describe an external identity provider (IdP) service that supports the OpenID Connect (OIDC) standard, such as Google or Salesforce. 


```
eksctl utils associate-iam-oidc-provider --cluster $cluster_name --approve
```




# Configuring Load Balancer


Create IAM Policy

```
aws iam create-policy \
    --policy-name AWSLoadBalancerControllerIAMPolicy \
    --policy-document file://iam_policy.json
```

Create IAM Role

```
eksctl create iamserviceaccount \
  --cluster=<your-cluster-name> \
  --namespace=kube-system \
  --name=aws-load-balancer-controller \
  --role-name AmazonEKSLoadBalancerControllerRole \
  --attach-policy-arn=arn:aws:iam::<your-aws-account-id>:policy/AWSLoadBalancerControllerIAMPolicy \
  --approve
```

## Deploying ALB controller

Add and update helm repo

```
helm repo add eks https://aws.github.io/eks-charts
```
```
helm repo update eks
```

Install

```
helm install aws-load-balancer-controller eks/aws-load-balancer-controller \            
  -n kube-system \
  --set clusterName=<your-cluster-name> \
  --set serviceAccount.create=false \
  --set serviceAccount.name=aws-load-balancer-controller \
  --set region=<region> \
  --set vpcId=<your-vpc-id>
```

Verify that the deployments are running.

```
kubectl get deployment -n kube-system aws-load-balancer-controller
```

Access the url in browser or use curl  to verify app is running sucessfully .





