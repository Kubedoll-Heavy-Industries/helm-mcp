# Changelog

## [0.1.3](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/compare/v0.1.2...v0.1.3) (2026-03-18)


### Bug Fixes

* **ci:** Bump aquasecurity/trivy-action from 0.34.1 to 0.35.0 ([ff4ec70](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/ff4ec70697e4b634a65816c8869de2c880a398ad))
* **ci:** Bump aquasecurity/trivy-action from 0.34.1 to 0.35.0 ([525ad12](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/525ad1249b33fbb6c1d7dadd0ecf28e7e41cb51f))
* **ci:** Bump docker/bake-action from 6 to 7 ([598f752](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/598f752c6dea035594a173e82c81cb9ebe519263))
* **ci:** Bump docker/bake-action from 6 to 7 ([bdd3c2f](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/bdd3c2f90cea820268e3e6d933f0931541eb4cab))
* **ci:** Bump docker/login-action from 3 to 4 ([d8e5197](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/d8e5197a4dfcf785601b79bdf4b76b0eec60bdfe))
* **ci:** Bump docker/login-action from 3 to 4 ([73be51f](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/73be51f8a103d14f394059be946cde13bc9ecaad))
* **ci:** Bump docker/setup-buildx-action from 3 to 4 ([c52ce31](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/c52ce31690feb90be4680bf68f8808ed59bcc04a))
* **ci:** Bump docker/setup-buildx-action from 3 to 4 ([141ba93](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/141ba9381e2d9ebbb9a4b08f32950d464e9fa52f))
* **ci:** Bump docker/setup-qemu-action from 3 to 4 ([2775ff2](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/2775ff26dadeb6c2ac0ee3463ed5c9be576b9fea))
* **ci:** Bump docker/setup-qemu-action from 3 to 4 ([3145c73](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/3145c73e9f90db06635bcbdfaed173197e50798e))
* **deps:** Bump github.com/modelcontextprotocol/go-sdk ([3f9f515](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/3f9f5159e0049013da64f05aeac341b5cfd65839))
* **deps:** Bump github.com/modelcontextprotocol/go-sdk from 1.4.0 to 1.4.1 ([b99038c](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/b99038cf6adf3f4079ae08b59e3a69cd120e9faf))
* **deps:** bump Go from 1.25 to 1.26 to fix stdlib vulnerabilities ([c219a5b](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/c219a5b13107cb4f2ffe9b80dc4251a6238eeb4b))
* **deps:** bump Go to 1.26.1 for stdlib security patches ([71215ff](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/71215ffa4acbc49e7f9ce26ffd5bda5df05f3dc0))
* **deps:** bump Go to 1.26.1 for stdlib security patches ([762df92](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/762df92316197c7280df1b224aca24b99c68fdc9))
* **deps:** Bump helm.sh/helm/v4 from 4.1.1 to 4.1.3 ([90e195f](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/90e195f4d45f071ca7f1f41e68628ff639ebf973))
* **deps:** Bump helm.sh/helm/v4 from 4.1.1 to 4.1.3 ([5f5db23](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/5f5db2309e442b8effd6aa6304dcde6e00ab58f3))
* **deps:** Bump vitest from 4.0.18 to 4.1.0 in /e2e-ts ([4c3b25a](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/4c3b25a53441fe6db519fefa65e3e995f74b3606))
* **deps:** Bump vitest from 4.0.18 to 4.1.0 in /e2e-ts ([7d961da](https://github.com/Kubedoll-Heavy-Industries/helm-mcp/commit/7d961dad96f696295f50b25ff112a3d5276b5bf4))
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
