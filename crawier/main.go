package main

import (
	"test/crawier/engine"
	"test/crawier/persist"
	"test/crawier/scheduler"
	"test/crawier/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemService("dating_profile")
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

	// e.Run(engine.Request{
	// 	Url: "http://www.zhenai.com/zhenghun/shanghai",
	// 	ParserFunc: parser.ParseCity,
	// })
}

