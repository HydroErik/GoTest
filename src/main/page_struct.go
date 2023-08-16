package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error  {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

/* Uses underscore to ignore potential error from os.readFile(...)
func loadPage(title string) *Page {
	filename := title+ ".txt"
	body, _ := os.ReadFile(filename)
	return &Page{Title: title, Body: body}
}
*/

//This version is error safe checing error return from os.ReadFile(...)
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "Test", Body: []byte("Hello World")}
	p1.save()
	p2, _ := loadPage("Test")
	fmt.Println(string(p2.Body))

}