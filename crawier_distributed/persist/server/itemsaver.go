package main

import (
	"fmt"
	"github.com/olivere/elastic"
	"test/crawier_distributed/config"
	"test/crawier_distributed/persist"
	"test/crawier_distributed/rpcsupport"
)

func main() {
	ServerRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex)
}

func ServerRpc(host string, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index: index,
	})
}
