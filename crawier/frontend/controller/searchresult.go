package controller

import (
	"context"
	"github.com/olivere/elastic"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"test/crawier/engine"
	"test/crawier/frontend/model"
	"test/crawier/frontend/view"
)

type SearchResultHandle struct {
	View view.SearchResultView
	Client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandle {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return SearchResultHandle{
		View: view.CreateSearchResultView(template),
		Client: client,
	}
}

// localhost:8888/search?q=天蝎&from=20
func (h SearchResultHandle) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err:= strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = h.View.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandle) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	resp, err := h.Client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(q)).
		From(from).
		Do(context.Background())
	if err != nil {
		return result, err
	}
	result.Hits = resp.TotalHits()
	result.Start = from

	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	return result, nil
}
