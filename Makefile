OUTPUT_DIR=bin
BINARY=gcp-paas

.PHONY: init, clean, test, run, generate, build, deploy, undeploy

init:
	go fmt ./...
	go mod tidy

clean:
	rm -fr ./${OUTPUT_DIR}
	rm -fr ./api/gen

run:
	go run cli/main.go

generate: 
	cd api; buf generate
	make init

build: 
	go build -o ./${OUTPUT_DIR}/${BINARY} ./cli/main.go

deploy:

undeploy: