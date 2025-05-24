APP := $(shell basename $(shell git remote get-url origin))
REGISTRY := ghcr.io/oltsy77
VERSION := $(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)

TARGETOS ?= linux
TARGETARCH ?= amd64

format:
	gofmt -s -w ./

get:
	go get ./...

lint:
	go vet ./...

test:
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go test -v ./...

build: format get
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o kbot -ldflags "-X=github.com/oltsy77/kbot/cmd.appVersion=${VERSION}"

image: build
		docker build \
		--build-arg TARGETOS=$(TARGETOS) \
		--build-arg TARGETARCH=$(TARGETARCH) \
		-t $(REGISTRY)/$(APP):$(VERSION)-$(TARGETOS)-$(TARGETARCH) .

push:
	docker push ${REGISTRY}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}

linux:
	$(MAKE) TARGETOS=linux TARGETARCH=amd64 image

arm:
	$(MAKE) TARGETOS=linux TARGETARCH=arm64 image

macos:
	$(MAKE) TARGETOS=darwin TARGETARCH=amd64 image

windows:
	$(MAKE) TARGETOS=windows TARGETARCH=amd64 image

clean:
	rm -f kbot
	docker rmi ${REGISTRY}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH} || true
