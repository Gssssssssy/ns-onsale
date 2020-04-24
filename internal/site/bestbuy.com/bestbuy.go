package bestbuy_com

import (
	"context"
	"encoding/json"
	"github.com/Gssssssssy/ns-onsale/internal/schedule"
	"github.com/Gssssssssy/ns-onsale/internal/site"
	"github.com/gocolly/colly"
	"github.com/avast/retry-go"
)

// 链接
const (
	InquiryURL              = ""
	ItemBlackDetailURL      = ""
	ItemBlueAndRedDetailURL = ""
)

// 货号
const (
	SkuIDBlack      = 6364253
	SkuIDBlueAndRed = 6364255
)

// 响应报文
type response struct {
}

func Inquiry(ctx context.Context, task schedule.Task) {
	var (
		c             = site.MakeColly()
		decoded       = response{}
		err, errOnRes error
	)
	c.OnResponse(func(r *colly.Response) {
		errOnRes = json.Unmarshal(r.Body, &decoded)
	})
	err = retry.Do(func() error {
		return c.Visit(InquiryURL)
	})
	if err != nil {
		return
	}
	if errOnRes != nil {
		return
	}
}
