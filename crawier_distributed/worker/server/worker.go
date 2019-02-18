package main

import (
	"fmt"
	"log"
	"test/crawier_distributed/config"
	"test/crawier_distributed/rpcsupport"
	"test/crawier_distributed/worker"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.WorkerPort0), worker.CrawlService{}))
}