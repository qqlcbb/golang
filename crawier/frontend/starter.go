package main

import (
	"net/http"
	"test/crawier/frontend/controller"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("crawier/frontend/view")))
	http.Handle("/search", controller.CreateSearchResultHandler("crawier/frontend/view/template.html"))
	err := http.ListenAndServe(":9191", nil)
	if err != nil {
		panic(err)
	}
}
