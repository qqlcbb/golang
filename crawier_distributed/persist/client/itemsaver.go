package client

import (
	"log"
	"test/crawier/engine"
	"test/crawier_distributed/config"
	"test/crawier_distributed/rpcsupport"
)

func ItemSaver(host string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	client , err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	go func() {
		itemCount := 0;
		for {
			item := <- out
			var result string
			itemCount ++
			err = client.Call(config.ItemSaverRpc, item, &result)
			if err != nil || result != "ok" {
				log.Printf("item saver: error saving item %s: %s", item, err)
			}
		}
	} ()
	return out, nil
}
