package client

import (
	"net/rpc"
	"test/crawier/engine"
	"test/crawier_distributed/config"
	"test/crawier_distributed/worker"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {

	return func(req engine.Request) (engine.ParseResult, error) {
		workerReq := worker.SerializeRequest(req)
		var result worker.ParserResult
		// 从channel获取client
		c := <- clientChan
		err := c.Call(config.CrawlServiceRpc, workerReq, &result)
		if err != nil {
			return engine.ParseResult{}, err
		}

		return worker.DeserializeResult(result), nil
	}
}
