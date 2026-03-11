package main

import "web/crawl/internal/crawler"

func main() {
	crawler.RootExec("https://books.toscrape.com/")
}