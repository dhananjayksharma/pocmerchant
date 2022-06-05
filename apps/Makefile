.PHONY: all

SERVICE_NAME := pacenow-service

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
	# @export MICROSERVICECAUTHNEWAPI=local
	@echo "[OK] Export var $(MICROSERVICECAUTHNEWAPI)"

.PHONY: run

run:
	@./pacenow-service

	#.PHONY: runasdocker

	#runasdocker:
	#	@echo docker build -t microservice-mediaphoto-service:v1 . && docker run --name services -p 8587:8587 -d microservice-mediaphoto-service:v1

.PHONY: build

build:
	@go build -o .
	@echo "[OK] Server was build!"

.PHONY: test

test:
	@go test -v	
	@echo "[OK] Test and coverage file was created!"

	#@go tool cover -html=cover.out -o cover.html 
	#-coverprofile cover.out ./...

.PHONY: show_coverage

show_coverage:
	#@go tool cover -html=cover.out
	@echo "[OK] Coverage file opened!"
