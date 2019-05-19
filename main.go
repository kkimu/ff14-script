package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	baseUrl := os.Getenv("BASE_URL")

	flag.Parse()
	id := flag.Arg(0)

	url := baseUrl + id
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	name := getName(doc)
	levels := getLevels(doc)

	fmt.Println(name)
	fmt.Println(levels)
}

func getLevels(doc *goquery.Document) []int {
	var levels []int
	doc.Find("div.character__profile__detail").Each(func(i int, doc2 *goquery.Selection) {
		doc2.Find("div.character__level").Each(func(i int, doc3 *goquery.Selection) {
			doc3.Find("li").Each(func(i int, doc4 *goquery.Selection) {
				level, err := strconv.Atoi(doc4.Text())
				if err != nil {
					level = 0
				}
				levels = append(levels, level)
			})
		})
	})
	return levels
}

func getName(doc *goquery.Document) string {
	name := doc.Find("p.frame__chara__name").Text()
	return name
}
