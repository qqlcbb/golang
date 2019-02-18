package main

import (
	"flag"
	"log"
	"net/rpc"
	"strings"
	"test/crawier/engine"
	"test/crawier/scheduler"
	"test/crawier/zhenai/parser"
	"test/crawier_distributed/config"
	itemsaver "test/crawier_distributed/persist/client"
	"test/crawier_distributed/rpcsupport"
	worker "test/crawier_distributed/worker/client"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workHosts = flag.String("worker_hosts", "", "worker hosts(comma separated)")
)

func main() {
	flag.Parse()

	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)

	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workHosts, ","))
	processor := worker.CreateProcessor(pool)

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

func createClientPool(hosts []string) chan *rpc.Client {

	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			log.Printf("connected to %s", h)
			clients = append(clients, client)
		} else {
			log.Printf("Error connected to %s; %v", h, err)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		// 死循环分发client给channel
		for {
			for _, client := range clients {
				out <- client
			}
		}
	} ()
	return out
}
