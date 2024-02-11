deploy:
	helm --kube-context $(cluster_name) upgrade \
		--install dragonfly \
		--values=./deploy/helm/dragonfly/values.yaml \
		--values=./deploy/helm/dragonfly/values_$(env).yaml \
		./deploy/helm/dragonfly

.PHONY: deploy