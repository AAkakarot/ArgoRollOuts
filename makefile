# Makefile for automating Docker build and Kubernetes deployment

# Define the image name and tag
IMAGE_NAME = akashkdoc/go-argo-app
TAG = latest

# Docker build and push commands
build:
    docker build --no-cache -t $(IMAGE_NAME):$(TAG) .

push:
    docker push $(IMAGE_NAME):$(TAG)

deploy:
    kubectl apply -f k8s/

# A combined command to build, push, and deploy
all: build push deploy

# A specific command to only deploy without rebuild
deploy-only:
    kubectl apply -f k8s/
