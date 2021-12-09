package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Post struct {
	Userid int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	// links := []string{"http://google.com", "http://facebook.com", "http://golang.org", "http://amazon.com"}

	// c := make(chan string)

	// for _, link := range links {
	// 	go checkLink(link, c)
	// }

	// for l := range c {
	// 	go func(l string) {
	// 		time.Sleep(5 * time.Second)
	// 		checkLink(l, c)
	// 	}(l)
	// }
	c := make(chan []Post)
	go getPosts(c)
	fmt.Println("hehehe")
	posts := <-c

	for _, post := range posts {
		fmt.Println(post)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

func getPosts(c chan []Post) {
	url := "https://jsonplaceholder.typicode.com/posts"

	res, err := http.Get(url)

	body, err := io.ReadAll(res.Body)

	var posts []Post

	json.Unmarshal(body, &posts)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	c <- posts

}
