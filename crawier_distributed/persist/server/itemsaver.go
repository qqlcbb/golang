package main

import (
	"flag"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"test/crawier_distributed/config"
	"test/crawier_distributed/persist"
	"test/crawier_distributed/rpcsupport"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(ServerRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
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
