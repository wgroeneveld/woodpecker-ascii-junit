FROM --platform=${BUILDPLATFORM} docker.io/golang:1.24 AS builder

WORKDIR /src
COPY . .

ARG TARGETOS TARGETARCH

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags '-s -w -extldflags "-static"' -v -a -o woodpecker-ascii-junit .

FROM alpine:3.21

COPY --from=builder /src/woodpecker-ascii-junit /bin/woodpecker-ascii-junit

RUN apk add -q --no-cache ca-certificates

ENTRYPOINT ["/bin/woodpecker-ascii-junit"]
