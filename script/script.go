package script

import (
	"log"

	"github.com/PyMarcus/go_webscraping/model"
	"github.com/gocolly/colly"
)

func getCollector(url string) *colly.Collector {
	return colly.NewCollector(
		colly.AllowedDomains(url),
	)
}

func request(c *colly.Collector){
	c.OnRequest(
		func(r *colly.Request){
			r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
			log.Println("Visiting ", r.URL)
		},
	)
}

func findOnHtml(url string) {
	var texts []model.Text 
	c := getCollector(url)

	request(c)

	c.OnHTML(".quote", func(h *colly.HTMLElement) {
		div := h.DOM 
		text := model.Text{
			Phrase: div.Find(".text").Text(),
			Author: div.Find(".author").Text(),
			Tag:    div.Find(".tags").Text(),
		}
		log.Println(text)
		texts = append(texts, text)
	})

	c.Visit("https://quotes.toscrape.com/random")

	printer(texts)
}

func printer(texts []model.Text){
	for _, text := range texts{
		log.Println(text.Phrase, text.Author, text.Tag)
	}
}

func Run(url string) {
	findOnHtml(url)
}
