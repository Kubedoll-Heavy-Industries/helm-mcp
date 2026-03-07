# Changelog

## [0.1.3](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/compare/v0.1.2...v0.1.3) (2026-03-07)


### Bug Fixes

* **deps:** bump Go from 1.25 to 1.26 to fix stdlib vulnerabilities ([c219a5b](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/c219a5b13107cb4f2ffe9b80dc4251a6238eeb4b))
* **deps:** bump Go to 1.26.1 for stdlib security patches ([71215ff](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/71215ffa4acbc49e7f9ce26ffd5bda5df05f3dc0))
* **deps:** bump Go to 1.26.1 for stdlib security patches ([762df92](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/762df92316197c7280df1b224aca24b99c68fdc9))
* **docker:** add USER directive for non-root container execution ([03cb117](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/03cb117b7c701a48c9af35f6c667765b43bd254a))
* **docker:** use built-in nobody user in debug stage ([033bcbf](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/033bcbfe908e6c7fe9cd1298531ceaa20734b95e))
* **handler:** strip inline comments from YAML keys and return raw YAML as text ([47db3ba](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/47db3bad089ca1de6817463cda6d7bf8b655f945))
* **server:** enable stateless HTTP mode to prevent stale sessions ([f2c8298](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/f2c82983b029c68d2cf3a4c51787234026042083))
* update repo references from mcp-helm to helm-mcp ([60eb10f](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/60eb10fc109adc14fef5b3d32df4ece62100f02a))
* v0.1.3 — repo rename, stateless HTTP, clean YAML output ([116b64f](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/116b64f122b207200dc0f7a47ba6ee300efa2f35))

## [0.1.2](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/compare/v0.1.1...v0.1.2) (2026-03-01)


### Bug Fixes

* **deploy:** fix Container import and update cloudflare deps ([#9](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/issues/9)) ([9bb9590](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/9bb95901443ef3dac7da5c7f573b87636919fd31))
* **server:** return JSON 404 for unmatched routes ([#11](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/issues/11)) ([b4cddb7](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/b4cddb7284fc763dd0a01ba5d013688f57abe930))

## [0.1.1](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/compare/v0.1.0...v0.1.1) (2026-02-28)


### Bug Fixes

* **ci:** reset release-please state to v0.1.0 ([#5](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/issues/5)) ([198775d](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/198775d177a925bb81631f5035fc15442d3ceadc))
* **deps:** upgrade go.opentelemetry.io/otel/sdk to v1.40.0 ([5387da2](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/5387da2eacdcd3e14e01b9012ddc56a9cf5ef22b))
* **deps:** upgrade otel SDK to fix GHSA-9h8m-3fm2-qjrq ([258b27a](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/258b27a3158d76e3bafdb659e6f89dac8b431f6f))
* **docker:** bump alpine from 3.23.0 to 3.23.3 ([#1](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/issues/1)) ([b2a84fe](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/b2a84fe3dc91afe442ed5cc1146abbc1a6632331))

## 0.1.0 (2026-02-28)

Initial release as a clean-history fork.

### Features

* MCP server providing read-only Helm chart tools for LLM agents
* `search_charts` — search for charts in HTTP/HTTPS repositories
* `get_versions` — list available chart versions (HTTP/HTTPS and OCI registries)
* `get_values` — retrieve values.yaml with depth-limited collapsing, path extraction, and optional JSON schema
* `get_dependencies` — list chart dependencies
* `get_notes` — retrieve NOTES.txt post-install instructions
* OCI registry support (`oci://`) for all chart-specific tools
* SSRF protection and URL validation for both HTTP and OCI endpoints
* Streamable HTTP and stdio transports
* Docker image published to GHCR
