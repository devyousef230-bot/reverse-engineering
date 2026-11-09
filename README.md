# Peperouny

<p align="center">
  <strong>An independent reverse-engineering project for NPVT & NPVS files.</strong>
  <br />
  Built through independent research, format analysis, and implementation.
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Status-Beta-orange" alt="Beta" />
  <img src="https://img.shields.io/badge/Formats-NPVT%20%7C%20NPVS-blue" alt="Supported formats" />
  <img src="https://img.shields.io/badge/Runtime-WebAssembly-purple" alt="WebAssembly" />
</p>

---

## Overview

**Peperouny** is an independent file-analysis and reverse-engineering project focused on understanding and inspecting NPVT and NPVS file formats.

The project was developed through independent investigation and implementation, using format analysis and reverse-engineering techniques rather than relying on another project's implementation.

Its goal is to make working with supported file variants more accessible through a lightweight browser-based interface.

## Features

- **NPVT support** — inspect supported `.npvt` files.
- **NPVS support** — inspect supported `.npvs` files.
- **Browser-based interface** — work with files through a simple web UI.
- **WebAssembly runtime** — run the analysis engine directly in the browser.
- **Local file processing** — files are read locally by the browser.
- **Password handling** — supports password input when required by a supported format variant.
- **Bilingual interface** — English and Persian.
- **Light and dark themes** — choose the appearance you prefer.

## Technology

| Component | Technology |
|---|---|
| Core engine | Go |
| Browser runtime | WebAssembly |
| Frontend | HTML, CSS, JavaScript |
| File handling | Browser File API |

## Getting Started

### Requirements

- Go 1.26.3 or newer
- A modern browser with WebAssembly support
- Python 3 (for the local development server)

### Build and run

Clone the repository:

```bash
git clone https://github.com/devyousef230-bot/reverse-engineering.git
cd reverse-engineering
```

Run tests:

```bash
go test ./...
```

Build the WebAssembly engine:

```bash
GOOS=js GOARCH=wasm go build -o web/peperouny.wasm ./cmd/wasm
```

Copy the Go WebAssembly runtime:

```bash
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" web/wasm_exec.js
```

Start a local server:

```bash
python -m http.server 8000 -d web
```

Open:

```text
http://localhost:8000
```

## Privacy & Security

Peperouny is designed to process selected files locally in the browser.

- File contents are read through the browser's File API.
- The WebAssembly engine performs processing in the browser.
- Decrypted output may contain sensitive information, including credentials or configuration data.

**Important:** Only inspect files you own or are explicitly authorized to analyze. Keep recovered data private and secure.

## Project Status

Peperouny is currently a beta project.

Support for individual file variants may differ. Some encrypted formats may require credentials or cryptographic material that cannot be recovered from the file alone.

## Independence

Peperouny is an independently developed project. Its reverse-engineering work and implementation are separate from other projects.

It is not affiliated with, endorsed by, or presented as a derivative of another project.

## Contributing

Suggestions, bug reports, and technical discussions are welcome.

If you discover an unsupported file variant or an unexpected behavior, please provide a minimal, non-sensitive example when reporting it.

## Disclaimer

This project is provided for educational, research, and authorized file-analysis purposes.

The authors are not responsible for misuse or unauthorized access to data.

---

<p align="center">
  <strong>Peperouny</strong>
  <br />
  Independent research. Independent implementation.
</p>
