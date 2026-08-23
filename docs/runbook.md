# Production Runbook

Operational guidance for the hosted instance at `helm-mcp.kubedoll.com`.

## Deploy pipeline

`Workers Builds` watches `main`. Every merge runs `npx wrangler deploy`,
which builds the Dockerfile (Go toolchain stamps VCS revision/time into the
binary via `-buildvcs`) and rolls out container instances. Build status
appears as the `Workers Builds: mcp-helm` check on each commit.

## Verify a deployment

The MCP `initialize` response reports the exact running build:

```bash
curl -s -X POST https://helm-mcp.kubedoll.com/mcp \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json, text/event-stream' \
  -d '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2025-06-18","capabilities":{},"clientInfo":{"name":"probe","version":"0"}}}'
```

`serverInfo.version` should read like `vdev (commit: <sha>, date: <iso8601>)`
for CI builds, or `vX.Y.Z (...)` for goreleaser release artifacts.
Compare `<sha>` against `git log origin/main -1`.

Liveness: `GET /healthz` returns 200.

## Rate limiting (recommended one-time dashboard step)

The public endpoint has no auth by design. Add a Cloudflare WAF rate-limit
rule so a single client cannot exhaust container instances:

1. dash.cloudflare.com → kubedoll.com zone → **Security → WAF → Rate limiting rules**
2. Create rule: match `http.host eq "helm-mcp.kubedoll.com"`, count both
   `/mcp` and `/healthz`, threshold ~120 requests / 1 minute / per IP,
   action Block (10 min).
3. Free plan includes one rule; keep it reserved for this host.

## Container egress

Containers run with `enableInternet = true` — required to fetch Helm
repository indexes and charts. The Go server enforces its own egress policy
(`--allowed-hosts`, `--denied-hosts`, private-IP blocking); see
`internal/helm` for defaults before loosening anything.

## Known vulnerability exceptions

Tracked in `.trivyignore`, each with rationale and removal conditions:

| ID | Why | Track |
|----|-----|-------|
| CVE-2026-50163 | oras-go/v2 ≤2.6.1 via Helm v4, no patched v2 release | #95 |
| GO-2026-5932 | review advisory vs x/crypto/openpgp; unreachable (govulncheck enforces), unfixed by design | revisit if we ever import openpgp |

## Rollback

Redeploy a known-good commit: re-run the Workers Builds deploy for that
commit from the dashboard (**Workers & Pages → mcp-helm → Deployments**),
or push a revert to `main` and let CD roll it out. Container rollouts are
gradual; allow several minutes before declaring failure.
