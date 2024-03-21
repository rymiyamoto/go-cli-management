batch_container = batch

.PHONY: dev
dev:
	docker compose up -d

.PHONY: init-batch
init-batch: dev
	docker compose exec $(batch_container) cobra-cli init

.PHONY: add-batch
add-batch: dev
	@$(eval batch_file := ${name}.go)
	@$(if $(name),, $(error name is not defined))
	@$(eval batch_file_exists := $(shell ls . | grep ${batch_file}))
	@$(if $(batch_file_exists), $(error $(name) is already exists))
	docker compose exec $(batch_container) cobra-cli add $(name) --config ./cobra.yml

.PHONY: run-batch
run-batch: dev
	docker compose exec $(batch_container) go run ./main.go $(line)
