FROM --platform=${BUILDPLATFORM} docker.io/golang:1.24 AS builder

WORKDIR /src
COPY . .

ARG TARGETOS TARGETARCH

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags '-s -w -extldflags "-static"' -v -a -o app-entrypoint .

FROM alpine:3.21

COPY --from=builder /src/app-entrypoint /bin/app-entrypoint

RUN apk add -q --no-cache ca-certificates

ENTRYPOINT ["/bin/app-entrypoint"]
