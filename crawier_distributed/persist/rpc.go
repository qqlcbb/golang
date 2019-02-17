package persist

import (
	"github.com/olivere/elastic"
	"log"
	"test/crawier/engine"
	"test/crawier/persist"
)

type ItemSaverService struct {
	Index string
	Client *elastic.Client
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v: %v", item, err)
	}
	return err
}
