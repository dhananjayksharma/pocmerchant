.PHONY: all

SERVICE_NAME := dkgosql-merchant-service-v3

include .envrc

all: exportvar clean test build run
	#run no use as 30Mar2022 runasdocker

.PHONY: clean

clean:
	@go clean
	@-rm -f cover.html cover.out
	@-rm -f pacenow-service
	@echo "[OK] Cover file and old build removed!"

.PHONY: exportvar

exportvar:
	# @export MICROSERVICECDEMONEWAPI=local
	@echo "[OK] Export var $(MICROSERVICECDEMONEWAPI)"

.PHONY: run

run:
	@./dkgosql-merchant-service-v3
	
	#.PHONY: runasdocker

	#runasdocker:
	#	@echo docker build -t microservice-mediaphoto-service:v1 . && docker run --name services -p 8587:8587 -d microservice-mediaphoto-service:v1

.PHONY: build

build:
	@go build -o .
	@echo "[OK] Server was build!"

.PHONY: test

mock-gen:
	mockgen -source pkg/v1/models/merchants/merchant.go -destination pkg/v1/models/mocks/merchant_mocks.go -package=mocks
	mockgen -source internals/adapter/mysql/query/mysqlstore.go -destination pkg/v1/models/mocks/query_mysqlstore_mocks.go -package=mocks

test:
	@go test -v	
	@echo "[OK] Test and coverage file was created!"

	#@go tool cover -html=cover.out -o cover.html 
	#-coverprofile cover.out ./...

.PHONY: show_coverage

show_coverage:
	#@go tool cover -html=cover.out
	@echo "[OK] Coverage file opened!"
