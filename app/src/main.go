package main

import (
	"encoding/json"

	"io"

	"log"

	"net/http"

	"os"

	"fmt"

	"github.com/gocolly/colly"

	teststring "crawler/teststring"
)

func main() {
	fmt.Println("Start server... 2")
	http.HandleFunc("/", ExampleHandler)
	http.HandleFunc("/api/crawler/getRussianText", GetRussianTextHandler)
	http.HandleFunc("/api/crawler/visit", visit)
	http.HandleFunc("/showString", ShowStringHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("** Service Started on Port " + port + " **")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// RussianText struct
type RussianText struct {
	PageURL    string
	RefererURL string
	Text       string
}

// ResultRussianText struct
type ResultRussianText struct {
	Status     uint16
	Action     string
	Attributes []RussianText
}

// ExampleHandler Some comment
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	m := RussianText{"pageURL", "refererURL", "text 3"}
	result, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	} else {
		io.WriteString(w, string(result))
	}
}

// GetRussianTextHandler Some comment
func GetRussianTextHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	m := RussianText{"sadasd", "asdasdasd", "asdasdasd"}
	am := ResultRussianText{200, "GetRussianText", []RussianText{m}}
	result, err := json.Marshal(am)
	if err != nil {
		log.Fatal(err)
	} else {
		io.WriteString(w, string(result))
	}
}

func visit(w http.ResponseWriter, r *http.Request) {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("bookingcamps.com"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		go c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://bookingcamps.com")
}

func ShowStringHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	m := teststring.ShowString()
	result, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	} else {
		io.WriteString(w, string(result))
	}
}
