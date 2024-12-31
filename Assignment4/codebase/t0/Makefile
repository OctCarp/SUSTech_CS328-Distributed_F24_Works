.PHONY: create-cluster delete-cluster apply-config update-image show-pods

CLUSTER_NAME := a4t0
DEPLOYMENT_NAME := a4t0
DEPLOYMENT_SERVICE := a4t0-service
OLD_IMAGE := a4-flask:1.0.0
NEW_IMAGE := a4-flask:1.0.1
K8S_CONFIG_FILE := t0.yaml

# Create a new cluster and load the images
create-cluster:
	kind create cluster --name $(CLUSTER_NAME) --config kind-config.yaml
	kind load docker-image $(OLD_IMAGE) --name $(CLUSTER_NAME)
	kind load docker-image $(NEW_IMAGE) --name $(CLUSTER_NAME)

# Delete the cluster
delete-cluster:
	kind delete cluster --name $(CLUSTER_NAME)

apply-config:
	kubectl apply -f $(K8S_CONFIG_FILE)
	kubectl set image deployment/$(DEPLOYMENT_NAME) flask-containers=$(OLD_IMAGE)

# Update the deployment with the new image
update-image:
	kubectl set image deployment/$(DEPLOYMENT_NAME) flask-containers=$(NEW_IMAGE)
	kubectl describe deployment $(DEPLOYMENT_NAME)

# Show the pods status
show-pods:
	kubectl get pods -o wide