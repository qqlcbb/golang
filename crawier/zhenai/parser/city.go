package parser

import (
	"regexp"
	"test/crawier/engine"
)

var profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

var cityUrlRe = regexp.MustCompile(`(http://www.zhenai.com/zhenghun/shanghai/[^"]+)`)


func ParseCity(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}
	matches := profileRe.FindAllSubmatch(contents, -1)
	for _, val := range matches {
		name := string(val[2])

		// result.Items = append(result.Items, "User " + name)
		result.Requests = append(result.Requests, engine.Request{
			Url: 		string(val[1]),
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParseProfile(contents, name)
			},
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
