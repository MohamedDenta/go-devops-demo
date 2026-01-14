IMAGE=ghcr.io/mohameddenta/go-devops-demo

.PHONY: deploy

deploy:
	@TAG=$$(git rev-parse --short HEAD); \
	kubectl kustomize k8s/base/ | \
	sed "s|$(IMAGE):[a-z0-9]*|$(IMAGE):$$TAG|" | \
	kubectl apply -f -; \
	kubectl rollout status deployment go-devops-demo
