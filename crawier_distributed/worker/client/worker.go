package client

import (
	"fmt"
	"test/crawier/engine"
	"test/crawier_distributed/config"
	"test/crawier_distributed/rpcsupport"
	"test/crawier_distributed/worker"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}
	return func(req engine.Request) (engine.ParseResult, error) {
		workerReq := worker.SerializeRequest(req)
		var result worker.ParserResult

		err = client.Call(config.CrawlServiceRpc, workerReq, &result)
		if err != nil {
			return engine.ParseResult{}, err
		}

		return worker.DeserializeResult(result), nil
	}, nil
}
