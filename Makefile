build:
	@rm -rf ./bin
	@go build -o bin/jat ./cmd/jatcli/main.go
