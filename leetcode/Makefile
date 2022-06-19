Q?=
create:
	go run cmd/main.go "$(Q)"

test:
	go test $(shell go list ./... | grep $(shell printf "%04d" $(Q))) -v