package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"starbob.com/go/test/morestrings"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {
	p1 := Page{Title: "My HTML", Body: []byte("This is a sample file.")}
	p1.save()
	p2, _ := loadPage(p1.Title)
	fmt.Println(string(p2.Body))

	morestrings.ReverseRunes("babak")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(writer http.ResponseWriter, req *http.Request) {

	writer.Write([]byte("Here's my response for " + (*req.URL).String()))
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}
