package persist

import (
	"context"
	"encoding/json"
	"test/crawier/engine"
	"test/crawier/model"
	"testing"

	"github.com/olivere/elastic"
)

func TestSave(t *testing.T) {
	expect := engine.Item{
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

	// TODO: try to start up elastic search
	// here using docker go client
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	// save item
	err = save(client, "dating_test", expect)
	if err != nil {
		panic(err)
	}

	// get id
	resp, err := client.Get().
		Index("dating_test").
		Type(expect.Type).
		Id(expect.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", *resp.Source)
	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FormJsonObj(actual.Payload)

	actual.Payload = actualProfile

	// verify result
	if actual != expect {
		t.Errorf("got %v; expect %v", actual, expect)
	}
}
