package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("cityList_test_data.html")

	if err != nil {
		panic(nil)
	}
	result := ParseCityList(contents)
	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCitys := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d request, but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if url != result.Requests[i].Url {
			t.Errorf("result should have %s url, but had %s", result.Requests[i].Url, url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d request, but had %d", resultSize, len(result.Items))
	}

	for i, item := range expectedCitys {
		if item != result.Items[i] {
			t.Errorf("result should have %s item, but had %s", result.Items[i], item)
		}
	}
}
