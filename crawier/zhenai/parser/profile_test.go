package parser

import (
	"io/ioutil"
	"test/crawier/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "杨丽颖")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element;but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expectProfield := model.Profile{
		Name: "杨丽颖",
		Marriger: "未婚",
		Age: 25,
		Xingzuo: "射手",
		Height: "158",
		Weight: "47",
		Income: "3千以下",
	}

	if profile != expectProfield {
		t.Errorf("expected %v;but was %v", expectProfield, profile)
	}
}
