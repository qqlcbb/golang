package worker

import (
	"errors"
	"fmt"
	"log"
	"test/crawier/engine"
	"test/crawier/zhenai/parser"
	"test/crawier_distributed/config"
)

type Request struct {
	Url string
	Parser SerializedParser
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()

	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParserResult {
	result := ParserResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}

	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url: r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(r ParserResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}

	for _, request := range r.Requests {
		engineParser, err := deserializeParser(request.Parser)
		if err != nil {
			log.Printf("error deserializing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engine.Request{
			Url: request.Url,
			Parser: engineParser,
		})
	}

	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParserCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParserCityList), nil
	case config.ParserCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParserCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParserProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		} else {
			return nil, fmt.Errorf("invalid arg: %v", p.Args)
		}
	default:
		fmt.Println(p.Name)
		return nil, errors.New("unknown parser name")
	}
}

type ParserResult struct {
	Items []engine.Item
	Requests []Request
}

type SerializedParser struct {
	Name string
	Args interface{}
}