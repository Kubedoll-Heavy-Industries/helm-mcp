ARG GO_VERSION=1.26.6

FROM golang:${GO_VERSION}-trixie@sha256:b75d466dd608587fd66cca705a307ba65b889827d06ad61d6a75f0482b51b7c7 AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY . .

ARG VERSION=dev
ARG COMMIT=none
ARG DATE=unknown
ARG SOURCE=https://github.com/kubedoll-heavy-industries/helm-mcp
ARG TARGETARCH=amd64

RUN --mount=type=cache,target=/root/.cache/go-build \
  CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} GOMAXPROCS=2 \
  go build -p=1 -trimpath -ldflags "-s -w -X main.version=${VERSION} -X main.commit=${COMMIT} -X main.date=${DATE}" \
  -o /out/mcp-helm ./cmd/mcp-helm

FROM gcr.io/distroless/static-debian12:nonroot@sha256:afa5c872c891853ca7fcf1f12c3edb23f7eeef36189728842dd51042ff57f7ab AS runtime

ARG VERSION=dev
ARG COMMIT=none
ARG DATE=unknown
ARG SOURCE=https://github.com/kubedoll-heavy-industries/helm-mcp

LABEL org.opencontainers.image.title="mcp-helm" \
  org.opencontainers.image.description="MCP server for interacting with Helm repositories and charts" \
  org.opencontainers.image.licenses="MIT" \
  org.opencontainers.image.source=$SOURCE \
  org.opencontainers.image.version=$VERSION \
  org.opencontainers.image.revision=$COMMIT \
  org.opencontainers.image.created=$DATE \
  io.modelcontextprotocol.server.name="io.github.kubedoll-heavy-industries/helm-mcp"

EXPOSE 8012

COPY --from=build /out/mcp-helm /mcp-helm

USER nonroot

HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 CMD ["/mcp-helm", "healthcheck"]

ENTRYPOINT ["/mcp-helm"]
CMD ["--listen=:8012", "--transport=http"]

FROM alpine:3.24.1@sha256:28bd5fe8b56d1bd048e5babf5b10710ebe0bae67db86916198a6eec434943f8b AS debug

RUN apk add --no-cache ca-certificates tzdata curl

EXPOSE 8012

COPY --from=build /out/mcp-helm /mcp-helm

USER nobody

HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 CMD ["/mcp-helm", "healthcheck"]

ENTRYPOINT ["/mcp-helm"]
CMD ["--listen=:8012", "--transport=http"]
