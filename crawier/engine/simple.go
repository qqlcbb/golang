package engine

import (
	"log"
	"test/crawier/fetcher"
)

type SimpleEngine struct {

}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		ParseResult, err := worker(r)
		if err != nil {
			continue
		}

		// 加三个点代表把slice里面的内容展开一个个加进去
		requests = append(requests, ParseResult.Requests...)
		for _, item := range ParseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	// log.Printf("Fetching Url %s", r.Url)
	// 获取url内容
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fecher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	// 通过parser提取内容
	return r.ParserFunc(body), nil
}
