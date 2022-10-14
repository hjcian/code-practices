.PHONY: venv, source

venv:
	@virtualenv -p python3 env

source:
	ln -s env/bin/activate

tidy:
	@go mod tidy

q?=
create:
	@go run cmd/main.go "$(q)"

target?=
test:
	@go test $(shell go list ./leetcode/... | grep $(q)) -run=$(target) -v
