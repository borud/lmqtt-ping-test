ifeq ($(OS),Windows_NT)
	EXTENSION=.exe
else
	EXTENSION=""
endif

all: test lint vet build

build: server client

server:
	@cd cmd/$@ && go build -o ../../bin/$@$(EXTENSIION) --trimpath -tags osusergo,netgo -ldflags="-s -w"
	@cd cmd/$@ && GOARCH=arm GOOS=linux go build -o ../../bin/$@-arm --trimpath -tags osusergo,netgo -ldflags="-s -w"
	@cd cmd/$@ && GOARCH=amd64 GOOS=linux go build -o ../../bin/$@-linux --trimpath -tags osusergo,netgo -ldflags="-s -w"

client:
	@cd cmd/$@ && go build -o ../../bin/$@$(EXTENSIION) --trimpath -tags osusergo,netgo -ldflags="-s -w"
	@cd cmd/$@ && GOARCH=arm GOOS=linux go build -o ../../bin/$@-arm --trimpath -tags osusergo,netgo -ldflags="-s -w"
	@cd cmd/$@ && GOARCH=amd64 GOOS=linux go build -o ../../bin/$@-linux --trimpath -tags osusergo,netgo -ldflags="-s -w"

test:
	@echo "*** $@"
	@go test ./...

vet:
	@echo "*** $@"
	@go vet ./...

lint:
	@echo "*** $@"
	@golint ./... 

clean:
	@rm -rf bin