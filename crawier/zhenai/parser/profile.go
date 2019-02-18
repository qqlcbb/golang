package parser

import (
	"regexp"
	"strconv"
	"test/crawier/engine"
	"test/crawier/model"
	"test/crawier_distributed/config"
)
// 名字
var nameRe = regexp.MustCompile(`<h1 class="nickName" data-v-5b109fc3>([^<]+)</h1>`)
// 年龄
var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)岁</div>`)
// 婚姻
var marrageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^<]+)</div>`)
// 星座
var XingzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^<]+)座`)
// 身高
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)cm</div>`)
// 体重
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)kg</div>`)
// 月收入
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>月收入:([^<]+)</div>`)

var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func parseProfile(contents []byte, url string, name string) engine.ParseResult {
	profile := model.Profile{}
	// 年龄
	age, _ := strconv.Atoi(extractString(contents, ageRe))
	if age != 0 {
		profile.Age = age
	}
	// profile.Name = extractString(contents, nameRe)
	// 使用上一级的名称
	profile.Name = name
	profile.Marriger = extractString(contents, marrageRe)
	profile.Xingzuo = extractString(contents, XingzuoRe)
	profile.Height = extractString(contents, heightRe)
	profile.Weight = extractString(contents, weightRe)
	profile.Income = extractString(contents, incomeRe)

	result := engine.ParseResult{
		Items: []engine.Item {
			{
				Url: url,
				Type: "zhenai",
				Id: extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	matches := re.FindSubmatch(contents)
	if len(matches) >= 2 {
		return string(matches[1])
	}
	return ""
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return config.ParserProfile, p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}

