package engine

import (
	"log"
	"test/crawier/fetcher"
)

func worker(r Request) (ParseResult, error) {
	// log.Printf("Fetching Url %s", r.Url)
	// 获取url内容
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fecher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	// 通过parser提取内容
	return r.ParserFunc(body, r.Url), nil
}
