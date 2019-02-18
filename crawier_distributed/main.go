package main

import (
	"fmt"
	"test/crawier/engine"
	"test/crawier/scheduler"
	"test/crawier/zhenai/parser"
	"test/crawier_distributed/config"
	itemsaver "test/crawier_distributed/persist/client"
	worker "test/crawier_distributed/worker/client"
)

func main() {
	itemChan, err := itemsaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))

	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheducler: &scheduler.QueuedScheduler{},
		WorkCount: 	100,
		ItemChan: 	itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParserCityList),
	})
}
