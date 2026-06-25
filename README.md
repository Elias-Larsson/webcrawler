# Webcrawler

A small concurrent web crawler written in Go. It starts from a seed URL, fetches pages, extracts links, keeps crawling links on the same hostname, respects basic `robots.txt` rules, and prints a summary of every URL it discovered.

The current CLI entry point is configured to crawl [Books to Scrape](https://books.toscrape.com/), a demo site commonly used for scraping practice.

## Features

- Concurrent crawling with a fixed worker pool
- Same-host crawling only
- Relative link resolution
- Fragment removal from URLs
- Basic `robots.txt` support for `Allow` and `Disallow` rules
- Duplicate URL tracking
- Console summary with crawl duration and discovered links

## Project structure

```text
.
├── cmd/
│   └── crawler/
│       └── main.go          # CLI entry point
├── internal/
│   ├── crawler/             # Crawl logic, workers, parsing, robots rules, reporting
│   ├── storage/             # Placeholder for future persistence
│   └── utils/               # Small helper utilities
├── go.mod
├── go.sum
└── README.md
```
## Requirements
Go 1.25 or newer
## Installation

Clone the repository:

`git clone https://github.com/Elias-Larsson/webcrawler.git`
`cd webcrawler`

Download dependencies:

`go mod download`
## Usage

Run the crawler:

`go run ./cmd/crawler`

By default, the crawler starts from:

https://books.toscrape.com/

The seed URL is currently hardcoded in cmd/crawler/main.go:

crawler.RootExec("https://books.toscrape.com/")

To crawl another site, update that value and run the command again.
