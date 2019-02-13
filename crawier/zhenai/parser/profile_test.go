package parser

import (
	"io/ioutil"
	"test/crawier/engine"
	"test/crawier/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "http://album.zhenai.com/u/87617540","杨丽颖")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element;but was %v", result.Items)
	}

	actual := result.Items[0]

	expectProfield := engine.Item{
		Url: "http://album.zhenai.com/u/87617540",
		Type: "zhenai",
		Id: "87617540",
		Payload: model.Profile{
			Name: "杨丽颖",
			Marriger: "未婚",
			Age: 25,
			Xingzuo: "射手",
			Height: "158",
			Weight: "47",
			Income: "3千以下",
		},
	}

	if actual != expectProfield {
		t.Errorf("expected %v;but was %v", expectProfield, actual)
	}
}
