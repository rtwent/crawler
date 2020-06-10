package main

import (
	"fmt"
	"regexp"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.MaxDepth(10),
		colly.URLFilters(
			regexp.MustCompile(`(?m)(.+)/filters/prodazha/(.+)`),
		),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36 Rostyslav"),
		// Visit only domains
		colly.AllowedDomains("www.atlanta.ua", "atlanta.ua"),
	)

	founded := 1
	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		re := regexp.MustCompile(`(?m)(.+)/object/(.+)/(\d+)`)
		if re.Match([]byte(link)) {
			// Print link
			//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
			founded++
			fmt.Printf("Link found: %s, total %d\n", link, founded)
		}

		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		//fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping
	c.Visit("https://www.atlanta.ua/odessa/filters/prodazha/kvartiry")
}
