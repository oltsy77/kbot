# FROM quay.io/projectquay/golang:1.21 AS builder
# FROM --platform=$BUILDPLATFORM quay.io/projectquay/golang:1.21 AS builder

# ARG TARGETOS
# ARG TARGETARCH

# WORKDIR /go/src/app
# COPY . .
# RUN make build

# FROM scratch
# WORKDIR /
# COPY --from=builder /go/src/app/kbot .
# COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# ENTRYPOINT ["./kbot", "start"]
FROM --platform=$BUILDPLATFORM quay.io/projectquay/golang:1.21 AS builder

ARG TARGETOS
ARG TARGETARCH

WORKDIR /go/src/app
COPY . .

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} make test

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} make build

FROM alpine
WORKDIR /
COPY --from=builder /go/src/app/kbot .
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["./kbot", "start"]
