# Winny API

Backend for **Winny**, the chat assistant on [www.winapps.io](https://www.winapps.io). It answers questions about Mitch Wintrow, his work history and his projects, grounded in markdown docs generated from the site's content layer, and refuses anything off-topic.

Go 1.24 stdlib HTTP server (`cmd/assistant-api`), port `10123` by default. The model runs in Ollama on the home server (2× RTX 5080) as `winny-ai-v1`.

## API
- `GET /api/v1/health`
- `POST /api/v1/chat/stream`: NDJSON stream of `{type: "meta" | "delta" | "done" | "error"}` events. The site widget depends on this shape; don't change it.
- `POST /api/v1/chat/reset`

## How answers are grounded
- `internal/content/**/page.md`: one doc per live site route (title, route, canonical URL, Winny notes, search keywords, links, content snapshot).
- `internal/content/facts.md`: current role, past roles with end dates, flagship projects, languages, contact. `catalog.BuildSystemPrompt` puts it into every prompt, so no role facts are hard-coded in Go.
- `internal/service/chat_service.go` ranks the top four docs per question. The out-of-scope gate passes a question if it hits a doc's Search Keywords, a small static list, or a ranked doc.

## Environment variables (names only)
`PORT`, `CORS_ALLOWED_ORIGINS`, `CONTENT_DIR` (default `./internal/content`), `OLLAMA_HOST`, `OLLAMA_BIN`, `OLLAMA_MODEL` (default `winny-ai-v1`), `OLLAMA_KEEP_ALIVE`, `OLLAMA_CONTEXT_LENGTH`, `OLLAMA_KV_CACHE_TYPE`, `OLLAMA_FLASH_ATTENTION`, `SITE_BASE_URL`, `SITE_SEED_PATHS`, `SITE_REFRESH_INTERVAL`, `SITE_MAX_PAGES`, `SITE_MAX_CONTENT_CHARS`, `SITE_REQUEST_TIMEOUT`, `SITE_USER_AGENT`, `HTTP_READ_TIMEOUT`, `HTTP_WRITE_TIMEOUT`, `HTTP_IDLE_TIMEOUT`.

## Regenerate the knowledge base
The docs are generated, never hand-edited. From the site repo (sibling checkout):

```sh
cd ../winapps_next && npm run export:winny   # writes ../winapps-api/internal/content, deletes stale docs
cd ../winapps-api && go test ./...           # loader + scope-gate tests run over the new docs
```

(`generate_content_docs.py` is gone; the exporter replaces it.)

## Rebuild the model
After changing `Modelfile.winny-ai-v1` (persona + rules only; it matches `catalog.Persona()`):

```sh
ollama create winny-ai-v1 -f Modelfile.winny-ai-v1
```

## Build and test
```sh
go vet ./... && go build ./... && go test ./...
go run ./cmd/assistant-api
```

## Deploy (home server)
```sh
cd /home/mwintrow/winapps/winny-ai/winny-ai
git pull
go build -o assistant-api ./cmd/assistant-api
ollama create winny-ai-v1 -f Modelfile.winny-ai-v1   # only if the Modelfile changed
sudo systemctl restart <winny-service-name>          # TODO(Mitch): confirm the unit name
curl -s https://winny-ai.winapps.dev/api/v1/health
```

Public host: `https://winny-ai.winapps.dev` behind nginx (`X-Accel-Buffering: no` for streaming).
