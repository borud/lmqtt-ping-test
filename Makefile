all: test lint vet build

build: server client

server:
	@cd cmd/$@ && go build -o ../../bin/$@$(EXTENSIION) --trimpath -tags osusergo,netgo -ldflags="$(LDFLAGS) -s -w"
	@cd cmd/$@ && GOARCH=arm GOOS=linux go build -o ../../bin/$@-arm --trimpath -tags osusergo,netgo -ldflags="$(LDFLAGS) -s -w"
	@cd cmd/$@ && GOARCH=amd64 GOOS=linux go build -o ../../bin/$@-linux --trimpath -tags osusergo,netgo -ldflags="$(LDFLAGS) -s -w"

client:
	@cd cmd/$@ && go build -o ../../bin/$@$(EXTENSIION) --trimpath -tags osusergo,netgo -ldflags="$(LDFLAGS) -s -w"
	@cd cmd/$@ && GOARCH=arm GOOS=linux go build -o ../../bin/$@-arm --trimpath -tags osusergo,netgo -ldflags="$(LDFLAGS) -s -w"
	@cd cmd/$@ && GOARCH=amd64 GOOS=linux go build -o ../../bin/$@-linux --trimpath -tags osusergo,netgo -ldflags="$(LDFLAGS) -s -w"

test:
	@echo "*** $@"
	@go test ./...

race:
	@echo "*** $@"
	@go test -race ./...

vet:
	@echo "*** $@"
	@go vet ./...

lint:
	@echo "*** $@"
	@revive ./... 

clean:
	@rm -rf bin