package main

import (
	"fmt"
	"test/crawier/engine"
	"test/crawier/scheduler"
	"test/crawier/zhenai/parser"
	"test/crawier_distributed/config"
	"test/crawier_distributed/persist/client"
)

func main() {
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))

	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheducler: &scheduler.QueuedScheduler{},
		WorkCount: 	100,
		ItemChan: 	itemChan,
	}

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
