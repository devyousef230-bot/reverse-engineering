# Peperouny

**A focused, local-first NPVT / NPVS file workbench.**

Peperouny is an independently branded web tool based on the decoding implementation in [Pantegnos](https://github.com/FrontierTM/Pantegnos). It intentionally focuses on only two formats: `.npvt` and `.npvs`.

## Scope

- Inspect and decode supported `.npvt` exports.
- Inspect and decode supported `.npvs` containers, including open exports and supported encrypted generations.
- Run the engine in the browser with Go WebAssembly.
- Keep selected files on the user's device; no file upload endpoint is part of this app.
- Persian / English UI and light / dark themes.

This is a beta research utility. It does not promise that every historical or future variant is supported. Recipient-encrypted NPVS exports may require recipient private-key material and are not automatically recoverable. A passphrase-protected file requires its actual passphrase.

## Run locally

Install Go 1.26.3 or newer:

```sh
go test ./...
GOOS=js GOARCH=wasm go build -o web/peperouny.wasm ./cmd/wasm
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" web/wasm_exec.js
python -m http.server 8000 -d web
```

Open http://localhost:8000. The browser console should report that the Peperouny engine is ready.

## Privacy

Files are read using the browser File API and passed to the local WebAssembly engine. The app does not transmit file contents. Decrypted output can contain credentials, tokens, hostnames, or other sensitive configuration data. Use only files you own or are authorized to inspect, and avoid sharing recovered output.

## Upstream and license

Peperouny is derived from Pantegnos by FrontierTM. The relevant upstream implementation files and embedded data are retained with their original notices where present. This derivative is distributed under **GNU Affero General Public License v3.0 (AGPL-3.0)**. See [LICENSE](LICENSE) and [NOTICE](NOTICE).

Peperouny is an independent project and is not endorsed by FrontierTM. The Peperouny brand and UI are separate from the upstream project.
