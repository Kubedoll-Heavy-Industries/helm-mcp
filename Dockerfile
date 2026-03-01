ARG GO_VERSION=1.25
ARG GO_DIGEST=sha256:fb4095b65a7bb89f039def7e33d7b90095d2c25f34597748758a6f209eead7ff

FROM golang:${GO_VERSION}-trixie@${GO_DIGEST} AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY . .

ARG VERSION=dev
ARG COMMIT=none
ARG DATE=unknown
ARG SOURCE=https://github.com/Kubedoll-Heavy-Industries/mcp-helm
ARG TARGETARCH=amd64

RUN --mount=type=cache,target=/root/.cache/go-build \
  CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} \
  go build -trimpath -ldflags "-s -w -X main.version=${VERSION} -X main.commit=${COMMIT} -X main.date=${DATE}" \
  -o /out/mcp-helm ./cmd/mcp-helm

FROM gcr.io/distroless/static-debian12:nonroot@sha256:5074667eecabac8ac5c5d395100a153a7b4e8426181cca36181cd019530f00c8 AS runtime

ARG VERSION=dev
ARG COMMIT=none
ARG DATE=unknown
ARG SOURCE=https://github.com/Kubedoll-Heavy-Industries/mcp-helm

LABEL org.opencontainers.image.title="mcp-helm" \
  org.opencontainers.image.description="MCP server for interacting with Helm repositories and charts" \
  org.opencontainers.image.licenses="MIT" \
  org.opencontainers.image.source=$SOURCE \
  org.opencontainers.image.version=$VERSION \
  org.opencontainers.image.revision=$COMMIT \
  org.opencontainers.image.created=$DATE

EXPOSE 8012

COPY --from=build /out/mcp-helm /mcp-helm

ENTRYPOINT ["/mcp-helm"]
CMD ["--listen=:8012", "--transport=http"]

FROM alpine:3.23.3@sha256:59855d3dceb3ae53991193bd03301e082b2a7faa56a514b03527ae0ec2ce3a95 AS debug

RUN apk add --no-cache ca-certificates tzdata curl && \
  addgroup -g 65532 -S nonroot && \
  adduser -u 65532 -S nonroot -G nonroot

EXPOSE 8012

COPY --from=build /out/mcp-helm /mcp-helm

HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
  CMD curl -f http://localhost:8012/healthz || exit 1

USER 65532:65532
ENTRYPOINT ["/mcp-helm"]
CMD ["--listen=:8012", "--transport=http"]
