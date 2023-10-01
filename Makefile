OUTPUT_DIR=bin
BINARY=gcp-paas

.PHONY: init, clean, test, run, generate, compile, build, deploy, undeploy

init:
	go fmt ./...
	go mod tidy
	cd api/infra; buf mod update

clean:
	rm -fr ./${OUTPUT_DIR}
	rm -fr ./api/gen

run:
	go run cli/main.go

generate: init
	cd api/infra; buf generate
	make init

compile: clean, generate
	go build ./...

build: 
	go build -o ./${OUTPUT_DIR}/${BINARY} ./cli/main.go

deploy:

undeploy: