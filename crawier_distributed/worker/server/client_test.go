package main

import (
	"fmt"
	"test/crawier_distributed/config"
	"test/crawier_distributed/rpcsupport"
	"test/crawier_distributed/worker"
	"testing"
	"time"
)

func TestCrawService(t *testing.T) {
	const host = ":9001"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/87617540",
		Parser: worker.SerializedParser{
			Name: config.ParserProfile,
			Args: "杨丽颖",
		},
	}

	var result worker.ParserResult

	err = client.Call(config.CrawlServiceRpc, req, &result)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
