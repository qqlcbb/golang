package parser

import (
	"regexp"
	"test/crawier/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" data-v-473e2ba0>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	// -1 代表有匹配的都输出出来
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		// result.Items = append(result.Items, "City " + string(m[2]))

		result.Requests = append(result.Requests, engine.Request{
				Url: 		string(m[1]),
				ParserFunc: ParseCity,
		})
	}

	return result
}
