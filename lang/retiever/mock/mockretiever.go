package mock

import "fmt"

type Retiever struct {
	Contents string
}

func (r *Retiever) String() string {
	return fmt.Sprintf("Retiever: {contents=}%s", r.Contents)
}

func (r *Retiever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

func (r *Retiever) Get(url string) string {
	return r.Contents
}

