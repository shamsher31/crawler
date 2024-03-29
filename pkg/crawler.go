package crawler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gocolly/colly"
)

// LinkInfo is a simple struct which collect links from website
type LinkInfo struct {
	StatusCode int
	Links      []string
}

// HealthCheck is used to verify if the service is up and running
func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Run executes crawler and collect links from the pages
func Run(w http.ResponseWriter, r *http.Request) {

	URL := r.URL.Query().Get("url")

	// Validate given URL
	if err := Validate(URL); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// collect links
	links, err := crawlURL(URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(links)
}

// crawlURL creates collector, visit URL and collect
// all the links from given pages
func crawlURL(URL string) (*LinkInfo, error) {

	u, err := url.Parse(URL)
	if err != nil {
		return nil, errors.New("Invalid URL")
	}

	c := colly.NewCollector(
		// Restrict crawling to specific domains
		colly.AllowedDomains(u.Host),

		// MaxDepth is 2, so only the links on the scraped page
		// and links on those pages are visited
		colly.MaxDepth(2),
	)

	c.Limit(&colly.LimitRule{
		// Filter domains affected by this rule
		DomainGlob: ".*" + u.Host + ".*",
		// Set a delay between requests to these domains
		Delay: 1 * time.Second,
		// Add an additional random delay
		RandomDelay: 1 * time.Second,

		// Limit the maximum parallelism to 2
		// This is necessary if the goroutines are dynamically
		// created to control the limit of simultaneous requests.
		Parallelism: 2,
	})

	info := &LinkInfo{Links: make([]string, 0)}

	// Search pages with links and add them to link collector
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link != "" {
			info.Links = append(info.Links, link)
		}
	})

	// extract status code
	c.OnResponse(func(r *colly.Response) {
		info.StatusCode = r.StatusCode
	})

	// collect any error during crawling
	c.OnError(func(r *colly.Response, err error) {
		log.Println("error:", r.StatusCode, err)
		info.StatusCode = r.StatusCode
	})

	// Hit url with above config set
	c.Visit(URL)

	return info, nil
}
