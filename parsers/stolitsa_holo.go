package parsers

import (
	"context"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"

	shared "printraduga_parser/shared"
)

type StolitsaHoloParser struct {
}

func (p StolitsaHoloParser) Parse() (shared.ParseResult, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx) // chromedp.WithDebugf(log.Printf),

	defer cancel()

	link := "https://stolitsaprint.ru/pechat-nakleek/golograficheskie/"

	// run task list
	var res string
	err := chromedp.Run(taskCtx,
		chromedp.EmulateViewport(1000, 1200),
		chromedp.Navigate(link),
		chromedp.WaitVisible("#grid-1-1-24 > div"),
		chromedp.SendKeys("#input_text-4", kb.Backspace+kb.Backspace+"300"),
		chromedp.SendKeys("#input_text-4", kb.Backspace+kb.Backspace+"300"),
		chromedp.SendKeys("#input_text-4", kb.Backspace+kb.Backspace+"300"),
		chromedp.Sleep(time.Hour),
	)
	if err != nil {
		return shared.ParseResult{}, err
	}
	intVar, err := strconv.Atoi(res)
	if err != nil {
		return shared.ParseResult{}, err
	}

	return shared.ParseResult{
		ParserType: "Translusent",
		Data: shared.CostData{
			Name: "Stolitsa",
			Cost: intVar,
			Link: link,
		},
	}, nil
}
