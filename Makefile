build-be:
	@echo "Build version: $(LDFLAGS)"
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./dist/service ./cmd/service

## Generate wire code
wire-up:
	@wire gen ./cmd/service

deploy-prod-all:
	make wire-up
	make build-be
	ansible-playbook -i playbooks/inventory.yml playbooks/tasks/oneshield.yaml --extra-vars "target_host_var=inet_cl"