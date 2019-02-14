package view

import (
	"os"
	"test/crawier/engine"
	"test/crawier/frontend/model"
	common "test/crawier/model"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")

	out, err := os.Create("template.test.html")

	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url: "http://album.zhenai.com/u/87617540",
		Type: "zhenai",
		Id: "87617540",
		Payload: common.Profile{
			Name: "杨丽颖",
			Marriger: "未婚",
			Age: 25,
			Xingzuo: "射手",
			Height: "158",
			Weight: "47",
			Income: "3千以下",
		},
	}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}