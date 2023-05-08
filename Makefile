deploy:
	helm --kubeconfig ~/.kube/config --kube-context docker-desktop upgrade \
		--install dragonfly \
		--values=./deploy/helm/dragonfly/values.yaml \
		--values=./deploy/helm/dragonfly/values_$(env).yaml \
		./deploy/helm/dragonfly

.PHONY: deploy