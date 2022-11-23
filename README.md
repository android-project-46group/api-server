# API server

[![License](https://img.shields.io/badge/license-MIT-blue)](./LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4)](https://pkg.go.dev/github.com/android-project-46group/api-server)
[![Go Report Card](https://goreportcard.com/badge/android-project-46group/api-server)](http://goreportcard.com/report/android-project-46group/api-server)

[![](https://img.shields.io/badge/android-555.svg?logo=kotlin)](https://github.com/android-project-46group/android)
[![](https://img.shields.io/badge/server-555.svg?logo=go)](https://github.com/android-project-46group/api-server)
[![](https://img.shields.io/badge/crawler-555.svg?logo=python)](https://github.com/android-project-46group/api)
[![](https://img.shields.io/badge/ios-555.svg?logo=swift)](https://github.com/android-project-46group/ios)

API server with ubuntu 20.04 on raspberry pi 4

## Tech Stack

-   ubuntu: 20.04
-   apache2: 2.4.41 (Ubuntu)
-   golang: 1.16 linux/arm64
-   postgres: 12.8 (Ubuntu 12.8-0ubuntu0.20.04.1)

## System Architecture

![System Architecture](./docs/system_architecture.png)

## Database design

![Database design](./docs/er.png)

## API design

[OpenAPI](./docs/openapi.yaml)

### DB Access (API)

Using test API key

-   https://kokoichi0206.mydns.jp/api/v1/members?gn=nogizaka&key=e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
-   https://kokoichi0206.mydns.jp/api/v1/songs?gn=nogizaka&key=e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
-   https://kokoichi0206.mydns.jp/api/v1/positions?title=%E3%81%A3%E3%81%A6%E3%81%8B&key=e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
-   https://kokoichi0206.mydns.jp/api/v1/formations?gn=hinatazaka&key=e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
-   https://kokoichi0206.mydns.jp/api/v1/blogs?gn=nogizaka&key=e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855

### Face Images

-   https://kokoichi0206.mydns.jp/imgs/nogi/iwamotorenka.jpeg
-   https://kokoichi0206.mydns.jp/imgs/sakura/fujiyoshikarin.jpeg
-   https://kokoichi0206.mydns.jp/imgs/hinata/matsudakonoka.jpeg

### Images of Latest Blogs

-   https://kokoichi0206.mydns.jp/imgs/blog/nogi/endousakura.jpeg
-   https://kokoichi0206.mydns.jp/imgs/blog/sakura/fujiyoshikarin.jpeg
-   https://kokoichi0206.mydns.jp/imgs/blog/hinata/matsudakonoka.jpeg

## Getting Started

```sh
# Setup

## ----- Migration for sqlboiler -----
go install github.com/volatiletech/sqlboiler/v4@v4.11.0
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

## version が異なると大量の差分が出る可能性があるので、なるべく揃えたい
sqlboiler --version
SQLBoiler v4.13.0
```

## License

This repository is under [MIT License](./LICENSE).
