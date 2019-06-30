GOPATH:=$(shell go env GOPATH)
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
DKRCMD=docker
DKRBUILD=$(DKRCMD) build
PORT=8080
APP=crawler

clean:
	# remove old binary
	rm -f ./bin/$(APP)
build:
	# building go binary for server:
	$(GOBUILD) -o bin/$(APP) cmd/server/main.go
run:
	#running binary
	env PORT=$(PORT) ./bin/$(APP)
docker-build:
	# building go binary for the docker container:
	env PORT=$(PORT) GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/$(APP) cmd/server/main.go
	# building docker container
	$(DKRBUILD) -t $(APP) .
docker-run:
	# running docker container
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		$(APP)
test:
	# running tests:
	$(GOTEST) ./... -v
