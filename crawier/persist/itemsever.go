package persist

import (
	"context"
	"log"

	"github.com/olivere/elastic"
)

func ItemService() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Sever: got item #%d: %v:", itemCount, item)
			itemCount ++
			_, err := save(item)
			if err != nil {
				log.Printf("Item Server: error saving item %v: %v", item, err)
			}
		}
	} ()
	return out
}

func save(item interface{}) (id string, err error){
	// must turn off sniff in docker
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return "", err
	}
	resp, err := client.Index().
		Index("dating_profile").
		Type("zhenai").BodyJson(item).
		Do(context.Background())
	if err != nil {
		return "", err
	}

	return resp.Id, nil
}
