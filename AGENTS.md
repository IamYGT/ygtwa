# YGTWA — Agent Instructions

This is **YGTWA**, a Go-based WhatsApp Web API bridge maintained by YGTWA.
It is based on the whatsmeow library and provides a REST API and MCP interface for WhatsApp.

## Repository Structure

```
src/
  cmd/          → CLI entry points (rest, mcp)
  config/       → Global settings and environment variable mapping
  domains/      → Interface definitions (ports)
  infrastructure/
    whatsapp/   → whatsmeow client, event handlers, device manager
    chatwoot/   → Optional Chatwoot integration
    chatstorage/→ SQLite chat history store
  ui/
    rest/       → Fiber HTTP handlers
    mcp/        → MCP (Model Context Protocol) handlers
    websocket/  → WebSocket event broadcasting
  usecase/      → Business logic
  validations/  → Request validation
  views/        → HTML frontend assets
  pkg/          → Shared utilities
docker/         → Dockerfile (golang.Dockerfile)
docs/           → OpenAPI spec, SDK config, guides
```

## Build

```bash
cd src && go build -o ../bin/ygtwa
```

## Docker Build

```bash
docker build -f docker/golang.Dockerfile -t ygtwa:latest .
```

## Key Conventions

- Multi-device support via `X-Device-Id` header
- All webhook events filtered by `WHATSAPP_WEBHOOK_WHITELIST_EVENTS`
- Chat history stored in SQLite (`storages/chatstorage.db`)
- Device sessions stored in `storages/whatsapp.db`
