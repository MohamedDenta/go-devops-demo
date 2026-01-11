IMAGE=ghcr.io/mohameddenta/go-devops-demo

.PHONY: deploy

deploy:
	@TAG=$$(git rev-parse --short HEAD); \
	kustomize edit set image $(IMAGE)=$(IMAGE):$$TAG; \
	kubectl apply -k k8s/; \
	kubectl rollout status deployment go-devops-demo
