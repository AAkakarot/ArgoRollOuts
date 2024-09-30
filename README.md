# ArgoRollOuts

Welcome to the **ArgoRollOuts** repository! This project focuses on implementing and managing progressive delivery strategies using [Argo Rollouts](https://argoproj.github.io/argo-rollouts/), a Kubernetes controller that provides advanced deployment capabilities like **Blue/Green Deployments**, **Canary Releases**, and **Automated Rollbacks**.

## Overview

Argo Rollouts extends the Kubernetes Deployment resource and provides additional features for controlling how updates are rolled out. With Argo Rollouts, you can enhance your Kubernetes application deployments by implementing safer, more reliable, and gradual deployment strategies.

### Key Features:
- **Blue/Green Deployments**: Safely switch between two environments for rapid rollbacks and easy updates.
- **Canary Releases**: Gradually shift traffic to the new version, monitor performance, and roll back if needed.
- **Traffic Shifting**: Integration with service mesh tools like Istio, Linkerd, and NGINX for fine-grained traffic management.
- **Automated Rollbacks**: Automatically revert to the previous version if health checks fail during a rollout.
- **Progressive Delivery**: Implement advanced strategies to release features to users in stages, minimizing risk.

## Getting Started

### Prerequisites

- **Kubernetes Cluster**: Ensure you have a running Kubernetes cluster (minikube, GKE, EKS, etc.).
- **kubectl**: The Kubernetes command-line tool must be installed and configured to connect to your cluster.
- **Argo Rollouts**: Install the Argo Rollouts controller by following the [official installation guide](https://argoproj.github.io/argo-rollouts/installation/).

### Installation

1. Install Argo Rollouts in your Kubernetes cluster:

   ```bash
   kubectl create namespace argo-rollouts
   kubectl apply -n argo-rollouts -f https://github.com/argoproj/argo-rollouts/releases/latest/download/install.yaml


2. Verify the installation by checking the Argo Rollouts controller pod:Verify the installation by checking the Argo Rollouts controller pod:

bash
--------- kubectl get pods -n argo-rollouts


Usage
To start using Argo Rollouts, define a rollout resource in your Kubernetes manifests. Here's an example of a basic canary rollout:

yaml: 

 apiVersion: argoproj.io/v1alpha1
kind: Rollout
metadata:
  name: example-rollout
spec:
  replicas: 3
  strategy:
    canary:
      steps:
        - setWeight: 20
        - pause: {duration: 60s}
        - setWeight: 50
        - pause: {duration: 60s}
  template:
    metadata:
      labels:
        app: example
    spec:
      containers:
      - name: example
        image: nginx:1.19.10
Apply this rollout to your cluster:

bash
---------- kubectl apply -f rollout.yaml

Monitor the progress using the Argo Rollouts kubectl plugin:

bash
 kubectl argo rollouts get rollout example-rollout
 
Contributions to ArgoRollOuts are welcome! Feel free to submit pull requests, open issues, or suggest enhancements.

For more information, visit the official Argo Rollouts documentation.

Fin!

