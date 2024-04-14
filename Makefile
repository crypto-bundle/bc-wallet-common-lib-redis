deploy_dragonfly:
	$(if $(and $(env),$(repository)),,$(error 'env' and/or 'repository' is not defined))

	$(eval context=$(or $(context),k0s-dev-cluster))
	$(eval platform=$(or $(platform),linux/amd64))

	helm --kube-context $(context) upgrade \
		--install dragonfly \
		--values=./deploy/helm/dragonfly/values.yaml \
		--values=./deploy/helm/dragonfly/values_$(env).yaml \
		./deploy/helm/dragonfly

destroy_dragonfly:
	$(if $(and $(env),$(repository)),,$(error 'env' and/or 'repository' is not defined))

	$(eval context=$(or $(context),k0s-dev-cluster))
	$(eval platform=$(or $(platform),linux/amd64))

	helm --kube-context $(context) uninstall dragonfly

build_redis_dependencies:
	$(if $(and $(env),$(repository)),,$(error 'env' and/or 'repository' is not defined))

	$(eval context=$(or $(context),k0s-dev-cluster))
	$(eval platform=$(or $(platform),linux/amd64))

	helm --kube-context $(context) dependency build \
		./deploy/helm/redis/bitnami/redis/

deploy_redis:
	$(if $(and $(env),$(repository)),,$(error 'env' and/or 'repository' is not defined))

	$(eval context=$(or $(context),k0s-dev-cluster))
	$(eval platform=$(or $(platform),linux/amd64))

	helm --kube-context $(context) upgrade \
		--install redis \
		--values=./deploy/helm/redis/bitnami/redis/values.yaml \
		--values=./deploy/helm/redis/bitnami/redis/values_$(env).yaml \
		./deploy/helm/redis/bitnami/redis/

destroy_redis:
	$(if $(and $(env),$(repository)),,$(error 'env' and/or 'repository' is not defined))

	$(eval context=$(or $(context),k0s-dev-cluster))
	$(eval platform=$(or $(platform),linux/amd64))

	helm --kube-context $(context) uninstall redis

.PHONY: deploy_dragonfly