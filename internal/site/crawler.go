package site

import (
	"context"
	"github.com/Gssssssssy/ns-onsale/internal/schedule"
	"github.com/gocolly/colly"
)

const DefaultRetryTimes uint = 3

type Collector interface {
	Inquiry(ctx context.Context, task schedule.Task)
}

func MakeColly() *colly.Collector {
	c := colly.NewCollector(colly.AllowURLRevisit())
	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
		request.Headers.Set("Accept-Encoding", "gzip, deflate, br")
		request.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	})
	return c
}
