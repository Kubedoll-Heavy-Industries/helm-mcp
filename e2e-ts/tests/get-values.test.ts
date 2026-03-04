import { describe, it, expect, inject, beforeAll, afterAll } from "vitest";
import { Client } from "@modelcontextprotocol/sdk/client/index.js";
import { StreamableHTTPClientTransport } from "@modelcontextprotocol/sdk/client/streamableHttp.js";

const baseUrl = inject("baseUrl");

let client: Client;
let transport: StreamableHTTPClientTransport;

beforeAll(async () => {
  transport = new StreamableHTTPClientTransport(new URL(`${baseUrl}/mcp`));
  client = new Client({ name: "e2e-test", version: "1.0.0" });
  await client.connect(transport);
});

afterAll(async () => {
  await client?.close();
  await transport?.close();
});

const REPO_URL = "https://prometheus-community.github.io/helm-charts";

describe("get_values", () => {
  it("returns YAML values and version for alertmanager", async () => {
    const result = await client.callTool({
      name: "get_values",
      arguments: { repository_url: REPO_URL, chart_name: "alertmanager" },
    });

    expect(result.isError).toBeFalsy();
    const content = result.content as Array<{ type: string; text: string }>;
    // Text content is raw YAML, not JSON
    expect(content[0].text).toBeTruthy();
    expect(content[0].text).toMatch(/replicaCount|image|service/);

    // Structured data is in structuredContent
    const structured = result.structuredContent as {
      version: string;
      values: string;
    };
    expect(structured.version).toBeTruthy();
    expect(typeof structured.version).toBe("string");
    expect(structured.values).toBeTruthy();
  });

  it("accepts include_schema parameter", async () => {
    const result = await client.callTool({
      name: "get_values",
      arguments: {
        repository_url: REPO_URL,
        chart_name: "prometheus-pushgateway",
        include_schema: true,
      },
    });

    expect(result.isError).toBeFalsy();
    const content = result.content as Array<{ type: string; text: string }>;
    expect(content[0].text).toBeTruthy();

    const structured = result.structuredContent as {
      version: string;
      values: string;
    };
    expect(structured.version).toBeTruthy();
    expect(structured.values).toBeTruthy();
  });

  it("narrows to subtree when path is specified", async () => {
    const result = await client.callTool({
      name: "get_values",
      arguments: {
        repository_url: REPO_URL,
        chart_name: "alertmanager",
        path: ".image",
      },
    });

    expect(result.isError).toBeFalsy();
    const content = result.content as Array<{ type: string; text: string }>;
    // Text content should contain image-related values
    expect(content[0].text).toMatch(/repository|tag|pullPolicy/);

    const structured = result.structuredContent as {
      values: string;
      path: string;
    };
    expect(structured.values).toBeTruthy();
    expect(structured.path).toBe(".image");
    expect(structured.values).toMatch(/repository|tag|pullPolicy/);
  });
});
