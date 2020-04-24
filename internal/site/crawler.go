package site

import (
	"github.com/gocolly/colly"
)

type Collector interface{}

func MakeColly() *colly.Collector {
	c := colly.NewCollector(colly.AllowURLRevisit())
	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	})
	return c
}
