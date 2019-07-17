package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Link contains the anchor link with the text and href
type Link struct {
	text string
	href string
}

func main() {
	url := "https://ryanbensonmedia.com"
	_, err := getPageLinks(url)
	if err != nil {
		log.Fatal(err)
	}
}

// getPageLinks runs the processing for getting links
// @params url {string} where to collect our links
// @return nil
func getPageLinks(url string) ([]Link, error) {
	body, err := getPage(url)
	if err != nil {
		return nil, err
	}

	content, err := getPageContent(body)
	if err != nil {
		return nil, err
	}

	links, err := getLinks(content)
	if err != nil {
		return nil, err
	}
	return links, err
}

// getPage collects the body of a url
// @params: url {string} what URL to retrieve
// @return: io.ReadCloser
func getPage(url string) (io.ReadCloser, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Status code is not valid %v", res.StatusCode)
	}
	return res.Body, nil
}

// getPageContent parses the body to get the page contents
// @params: body {io.ReadCloser} body of the page to parse
// @return: *goquery.Document
func getPageContent(body io.ReadCloser) (*goquery.Document, error) {
	if body == nil {
		return nil, errors.New("Must provide a valid document")
	}
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// getLinks parses the page to get all of the <a> links with an href
// @params: doc {*goquery.Document} parsed version of the body
// @return: []Link
func getLinks(doc *goquery.Document) ([]Link, error) {
	if doc == nil {
		return nil, errors.New("Must provide a valid document")
	}

	var links []Link
	doc.Find("body a").Each(func(index int, linkTag *goquery.Selection) {
		link, exists := linkTag.Attr("href")
		if exists {
			linkText := strings.TrimSpace(linkTag.Text())
			l := Link{text: linkText, href: link}
			links = append(links, l)
		}
	})
	return links, nil
}
