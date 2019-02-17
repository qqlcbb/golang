package persist

import (
	"context"
	"errors"
	"log"
	"test/crawier/engine"

	"github.com/olivere/elastic"
)

func ItemService(index string) (chan engine.Item, error) {
	// must turn off sniff in docker
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Sever: got item #%d: %v:", itemCount, item)
			itemCount ++
			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item Server: error saving item %v: %v", item, err)
			}
		}
	} ()
	return out, nil
}

func Save(client *elastic.Client,index string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
