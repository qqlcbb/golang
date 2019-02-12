package persist

import (
	"context"
	"encoding/json"
	"test/crawier/model"
	"testing"

	"github.com/olivere/elastic"
)

func TestSave(t *testing.T) {
	expect := model.Profile{
		Name: "杨丽颖",
		Marriger: "未婚",
		Age: 25,
		Xingzuo: "射手",
		Height: "158",
		Weight: "47",
		Income: "3千以下",
	}

	id, err := save(expect)
	if err != nil {
		panic(err)
	}
	// TODO: try to start up elastic search
	// here using docker go client
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index("dating_profile").
		Type("zhenai").Id(id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", *resp.Source)
	var actual model.Profile
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	if actual != expect {
		t.Errorf("got %v; expect %v", actual, expect)
	}
}
