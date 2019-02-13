package parser

import (
	"regexp"
	"test/crawier/engine"
)

var profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

var cityUrlRe = regexp.MustCompile(`(http://www.zhenai.com/zhenghun/shanghai/[^"]+)`)


func ParseCity(contents []byte, _ string) engine.ParseResult {
	result := engine.ParseResult{}
	matches := profileRe.FindAllSubmatch(contents, -1)
	for _, val := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: 		string(val[1]),
			ParserFunc: ProfileParse(string(val[2])),
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, val := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: 		string(val[1]),
			ParserFunc: ParseCity,
		})
	}


	return result
}
