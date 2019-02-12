package main

import (
	"fmt"
	"test/retiever/mock"
	"test/retiever/real"
	"time"
)

// 使用者规定接口有get方法
type Retiever interface {
	// 不用加func关键字
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retiever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string {
			"name" : "ccmouse",
			"course" : "golang",
		})
}

type RetieverPoster interface {
	Retiever
	Poster
}

func session(s RetieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "anther fake imooc.com",
	})

	return s.Get(url)
}

func main() {
	var r Retiever

	retiever := mock.Retiever{"this is fake imooc.com"}

	r = &retiever
	inspect(r)

	r = &real.Retiever{
		UserAgent: "macbook 2018",
		TimeOut: time.Minute,
	}
	inspect(r)

	// Type assertion
	realRetiever := r.(*real.Retiever)
	fmt.Println(realRetiever.TimeOut)

	fmt.Println("try a session")
	fmt.Println(session(&retiever))
}

func inspect(r Retiever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T, %v\n", r, r)
	fmt.Print(" > Type Switch:")
	// 判断接口类型
	switch v := r.(type) {
	case *mock.Retiever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retiever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println("Inspecting", r)
}
