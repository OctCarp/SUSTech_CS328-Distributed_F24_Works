.PHONY: create-cluster delete-cluster list-node-labels list-node-taints apply-old apply-new
.PHONY: show-pods show-describe scale scale-1 scale-2 scale-3 scale-4 scale-5 show-pods

CLUSTER_NAME := a4t1
DEPLOYMENT_NAME := a4t1-deployment
K8S_CONFIG_FILE_OLD := t1-old.yaml
K8S_CONFIG_FILE_NEW := t1-new.yaml

create-cluster:
	kind create cluster --name $(CLUSTER_NAME) --config kind-config.yaml

delete-cluster:
	kind delete cluster --name $(CLUSTER_NAME)

list-node-labels:
	# kubectl get nodes --show-labels
	kubectl get nodes -o jsonpath='{range .items[*]}{.metadata.name}{"\t"}{.metadata.labels}{"\n"}{end}'

list-node-taints:
	kubectl get nodes -o jsonpath='{range .items[*]}{.metadata.name}{"\t"}{.spec.taints}{"\n"}{end}'

apply-old:
	kubectl apply -f $(K8S_CONFIG_FILE_OLD)

apply-new:
	kubectl apply -f $(K8S_CONFIG_FILE_NEW)

show-describe:
	kubectl describe deployment $(DEPLOYMENT_NAME)

show-pods:
	kubectl get pods -o wide

scale:
	kubectl scale deployment $(DEPLOYMENT_NAME) --replicas=$(REPLICAS)

scale-1: REPLICAS=1
scale-1: scale

scale-2: REPLICAS=2
scale-2: scale

scale-3: REPLICAS=3
scale-3: scale

scale-4: REPLICAS=4
scale-4: scale

scale-5: REPLICAS=5
scale-5: scale

