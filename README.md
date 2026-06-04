# Worten CLI

Local OpenAPI seed extracted from the working housebuy Worten CLI.
This does not claim to describe the full Worten platform.
It only captures the load-bearing endpoints currently used by the housebuy operator.

Created by [@asantos00](https://github.com/asantos00) (Alexandre Santos).

## Install

The recommended path installs both the `worten-pp-cli` binary and the `pp-worten` agent skill (Claude Code, Codex, Cursor, Gemini CLI, GitHub Copilot, and other agents supported by the upstream [`skills`](https://github.com/vercel-labs/skills) CLI) in one shot:

```bash
npx -y @mvanhorn/printing-press-library install worten
```

For CLI only (no skill):

```bash
npx -y @mvanhorn/printing-press-library install worten --cli-only
```

For skill only — installs the skill into the same agents as the default command above, but skips the CLI binary (use this to update or reinstall just the skill):

```bash
npx -y @mvanhorn/printing-press-library install worten --skill-only
```

To constrain the skill install to one or more specific agents (repeatable — agent names match the [`skills`](https://github.com/vercel-labs/skills) CLI):

```bash
npx -y @mvanhorn/printing-press-library install worten --agent claude-code
npx -y @mvanhorn/printing-press-library install worten --agent claude-code --agent codex
```

### Without Node

The generated install path is category-agnostic until this CLI is published. If `npx` is not available before publish, install Node or use the category-specific Go fallback from the public-library entry after publish.

### Pre-built binary

Download a pre-built binary for your platform from the [latest release](https://github.com/mvanhorn/printing-press-library/releases/tag/worten-current). On macOS, clear the Gatekeeper quarantine: `xattr -d com.apple.quarantine <binary>`. On Unix, mark it executable: `chmod +x <binary>`.

<!-- pp-hermes-install-anchor -->
## Install for Hermes

From the Hermes CLI:

```bash
hermes skills install mvanhorn/printing-press-library/cli-skills/pp-worten --force
```

Inside a Hermes chat session:

```bash
/skills install mvanhorn/printing-press-library/cli-skills/pp-worten --force
```

## Install for OpenClaw

Tell your OpenClaw agent (copy this):

```
Install the pp-worten skill from https://github.com/mvanhorn/printing-press-library/tree/main/cli-skills/pp-worten. The skill defines how its required CLI can be installed.
```

## Use with Claude Desktop

This CLI ships an [MCPB](https://github.com/modelcontextprotocol/mcpb) bundle — Claude Desktop's standard format for one-click MCP extension installs (no JSON config required).

To install:

1. Download the `.mcpb` for your platform from the [latest release](https://github.com/mvanhorn/printing-press-library/releases/tag/worten-current).
2. Double-click the `.mcpb` file. Claude Desktop opens and walks you through the install.

Requires Claude Desktop 1.0.0 or later. Pre-built bundles ship for macOS Apple Silicon (`darwin-arm64`) and Windows (`amd64`, `arm64`); for other platforms, use the manual config below.

<details>
<summary>Manual JSON config (advanced)</summary>

If you can't use the MCPB bundle (older Claude Desktop, unsupported platform), install the MCP binary and configure it manually.


Install the MCP binary from this CLI's published public-library entry or pre-built release.

Add to your Claude Desktop config (`~/Library/Application Support/Claude/claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "worten": {
      "command": "worten-pp-mcp"
    }
  }
}
```

</details>

## Quick Start

### 1. Install

See [Install](#install) above.

### 2. Verify Setup

```bash
worten-pp-cli doctor
```

This checks your configuration.

### 3. Try Your First Command

```bash
worten-pp-cli worten-api get-offer-stock --offer-id 550e8400-e29b-41d4-a716-446655440000 --search-query example-value --radius 42
```

## Usage

Run `worten-pp-cli --help` for the full command reference and flag list.

## Commands

### worten-api

Manage worten api

- **`worten-pp-cli worten-api get-offer-stock`** - Fetch nearby store pickup stock for an offer.
- **`worten-pp-cli worten-api get-product-details`** - Fetch product details by Worten product identifier.
- **`worten-pp-cli worten-api get-search-suggestions`** - Fetch search suggestions for a text query.
- **`worten-pp-cli worten-api get-technical-specifications`** - Fetch technical specifications for a Worten product.
- **`worten-pp-cli worten-api search-products`** - Search Worten products by query and context.


## Output Formats

```bash
# Human-readable table (default in terminal, JSON when piped)
worten-pp-cli worten-api get-offer-stock --offer-id 550e8400-e29b-41d4-a716-446655440000 --search-query example-value --radius 42

# JSON for scripting and agents
worten-pp-cli worten-api get-offer-stock --offer-id 550e8400-e29b-41d4-a716-446655440000 --search-query example-value --radius 42 --json

# Filter to specific fields
worten-pp-cli worten-api get-offer-stock --offer-id 550e8400-e29b-41d4-a716-446655440000 --search-query example-value --radius 42 --json --select id,name,status

# Dry run — show the request without sending
worten-pp-cli worten-api get-offer-stock --offer-id 550e8400-e29b-41d4-a716-446655440000 --search-query example-value --radius 42 --dry-run

# Agent mode — JSON + compact + no prompts in one flag
worten-pp-cli worten-api get-offer-stock --offer-id 550e8400-e29b-41d4-a716-446655440000 --search-query example-value --radius 42 --agent
```

## Agent Usage

This CLI is designed for AI agent consumption:

- **Non-interactive** - never prompts, every input is a flag
- **Pipeable** - `--json` output to stdout, errors to stderr
- **Filterable** - `--select id,name` returns only fields you need
- **Previewable** - `--dry-run` shows the request without sending
- **Explicit retries** - add `--idempotent` to create retries when a no-op success is acceptable
- **Confirmable** - `--yes` for explicit confirmation of destructive actions
- **Piped input** - write commands can accept structured input when their help lists `--stdin`
- **Offline-friendly** - sync/search commands can use the local SQLite store when available
- **Agent-safe by default** - no colors or formatting unless `--human-friendly` is set

Exit codes: `0` success, `2` usage error, `3` not found, `5` API error, `7` rate limited, `10` config error.

## Health Check

```bash
worten-pp-cli doctor
```

Verifies configuration and connectivity to the API.

## Configuration

Config file: `~/.config/worten-reverse-engineered-pp-cli/config.toml`

Static request headers can be configured under `headers`; per-command header overrides take precedence.

## Troubleshooting
**Not found errors (exit code 3)**
- Check the resource ID is correct
- Run the `list` command to see available items

---

Generated by [CLI Printing Press](https://github.com/mvanhorn/cli-printing-press)
