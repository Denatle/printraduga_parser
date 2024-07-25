package parsers

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"

	shared "printraduga_parser/shared"
)

type CoralTranslusentParser struct {
}

func (p CoralTranslusentParser) Parse() (shared.ParseResult, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", true))

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx) // chromedp.WithDebugf(log.Printf),
	defer cancel()

	link := "https://www.coral-print.ru/pechat-nakleek/"

	// run task list
	var res string
	err := chromedp.Run(taskCtx,
		chromedp.Navigate(link),
		chromedp.WaitVisible(".wrap"),
		chromedp.SetValue("#count", "1000"),
		// chromedp.SendKeys("#count", "1000"+kb.Delete+kb.Delete+kb.Delete),
		chromedp.SetValue("#m", "m3"),
		chromedp.Sleep(time.Millisecond*500),
		chromedp.Text(".vertical-align-middle > .price", &res, chromedp.NodeVisible),
	)
	if err != nil {
		return shared.ParseResult{}, err
	}
	intVar, err := strconv.Atoi(strings.Replace(res, " руб", "", -1))
	if err != nil {
		return shared.ParseResult{}, err
	}

	return shared.ParseResult{
		ParserType: "Translusent",
		Data: shared.CostData{
			Name: "Coral",
			Cost: intVar,
			Link: link,
		},
	}, nil
}
