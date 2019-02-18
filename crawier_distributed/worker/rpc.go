package worker

import (
	"test/crawier/engine"
)

type CrawlService struct {}

func (CrawlService) Process(req Request, result *ParserResult) error {

	// 反序列化
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}

	// 序列化
	*result = SerializeResult(engineResult)

	return nil
}
