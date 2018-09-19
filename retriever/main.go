package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		from map[string]string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func post(poster Poster) {
	poster.Post("http://www.imooc.com",
		map[string]string{
			"name":   "zhangsan",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "http://www.imooc.com"

func session(s RetrieverPoster) string {

	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"This is me"}
	r = &retriever

	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}

	inspect(r)

	//type assertion
	if realRetriever, ok := r.(*real.Retriever); ok {
		fmt.Println(realRetriever.TimeOut)
	} else {
		fmt.Println("not a real retriever")
	}

	fmt.Println(session(&retriever))
	//fmt.Println(download(r))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
