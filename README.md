# wurl - console client for websocket protocol

[![Documentation](https://godoc.org/github.com/github.com/xakep666/wurl?status.svg)](http://godoc.org/github.com/xakep666/wurl)
[![Go Report Card](https://goreportcard.com/badge/github.com/xakep666/wurl)](https://goreportcard.com/report/github.com/xakep666/wurl)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/github.com/xakep666/wurl/LICENSE)

## Abstract

At the moment we already have quite few websocket clients. Starting from browser addons, ending with console clients.

But I`m not satisfied with either of them. Browser addons requires installed and running browser.
NodeJS-based clients requires node and tons of dependencies.
But most importantly, none of them allows you to specify additional headers for request.

So I decided to write own console websocket client...

## Installation
`go get -u github.com/xakep666/wurl`

`vgo get -u github.com/xakep666/wurl`

Pre-built binary releases available for:
- Linux: x86_64, x86, arm
- Mac OS (darwin): x86_64
- Windows: x86_64, x86

## Current features
- Read text/binary messages from connection and display it
- Ability to set additional headers for connection upgrade request
- Correctly processes ping message (by default responses with pongs message)
- Can periodically send ping message to server (period can be set through flags)

### TODOs for v1
- [x] Document all packages
- [x] Flag to show handshake response
- [x] Store and load options from file
- [x] Warning about binary messages before displaying (cURL-like)
- [x] Ability to specify output
- [x] Option to send message to server before reading
- [x] Good description for all flags/commands
- [x] Proxy support
- [x] Bash autocomplete
- [ ] Package to rpm, deb, for Arch (after release)
