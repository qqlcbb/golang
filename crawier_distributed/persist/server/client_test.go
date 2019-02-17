package main

import (
	"test/crawier/engine"
	"test/crawier/model"
	"test/crawier_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	// start ItemSaverServer
	go ServerRpc(host, "test1")
	time.Sleep(time.Second)

	// start SaverClient
	client , err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// Call save
	item := engine.Item{
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

	var result string
	err = client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
