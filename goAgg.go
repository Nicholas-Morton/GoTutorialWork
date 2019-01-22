package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

func main() {

	var s SitemapIndex
	var n News
	//NM := make(map[string]NewsMap)
	i := 0
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	//s.Locatoins[:1] to remove newlines cause what the fuck

	for _, Location := range s.Locations {
		resp, _ = http.Get(Location)
		bytes, _ = ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		fmt.Printf(n.Titles[i], n.Keywords[i], n.Locations[i])
		i++
		/*
			for idx, _ := range n.Titles {
				NM[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
				fmt.Println()
			}
		*/
	}

	fmt.Println(len(n.Titles))

	//fmt.Printf("Title: %S Keywords: %S Locations: %S", n.Titles[1], n.Keywords[1], n.Locations[1])

	/*
		for idx, data := range NM {
			fmt.Println("\n\n\n", idx)
			fmt.Println("\n", data.Keyword)
			fmt.Println("\n", data.Location)
		}
	*/
}
