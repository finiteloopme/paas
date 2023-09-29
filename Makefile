OUTPUT_DIR=bin
BINARY=gcp-paas

.PHONY: init, clean, test, run, build, deploy, undeploy

init:
	go fmt ./...
	go mod tidy

clean:
	rm -fr ./${OUTPUT_DIR}

run:
	go run cli/main.go

build:
	go build -o ./${OUTPUT_DIR}/${BINARY} ./...

deploy:

undeploy: