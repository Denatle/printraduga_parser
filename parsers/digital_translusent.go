package parsers

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"

	shared "printraduga_parser/shared"
)

type DigitalTranslusentParser struct {
}

func (p DigitalTranslusentParser) Parse() (shared.ParseResult, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", true))

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx) // chromedp.WithDebugf(log.Printf),

	defer cancel()
	// ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	// defer cancel()

	link := "https://digital-printing.ru/prints/stickers/prozrachnye#agree"
	// run task list
	var res string
	err := chromedp.Run(taskCtx,
		chromedp.EmulateViewport(1000, 1200),
		chromedp.Navigate(link),
		chromedp.WaitVisible("#pxpProducCalc > div.pxp-material-selector > div > div > ul > li.material-selector__item.material-types.option-row > span"),
		chromedp.Click("#pxpProducCalc > div.pxp-material-selector > div > div > ul > li:nth-child(3) > ul > li:nth-child(7) > a"),
		chromedp.Sleep(time.Second*4),
		chromedp.Evaluate(`document.querySelector("#pxpProducCalc > div.pxp-custom-works-selector > div > div > ul > li:nth-child(1) > ul > li:nth-child(2) > span > label").click()`, nil),
		chromedp.Click("#pxpProducCalc > div.pxp-custom-works-selector > div > div > ul > li:nth-child(1) > ul > li.option-item.with-helper.selected > span > label"),
		// chromedp.Click("#pxpProducCalc > div.pxp-custom-works-selector > div > div > ul > li:nth-child(2) > ul > li.option-item.with-helper.zero.selected > span > label"),
		chromedp.SendKeys("#txtQuantity", kb.Backspace+kb.Backspace+"300"),
		chromedp.Sleep(time.Second*2),
		chromedp.Text("#pxpProducCalc > div.pxp-total-price > div > div.totalPriceContainer > div > span", &res, chromedp.NodeVisible),
	)
	if err != nil {
		return shared.ParseResult{}, err
	}
	trimmedString := strings.Replace(res, " ", "", -1)
	intVar, err := strconv.Atoi(trimmedString[:len(trimmedString)-3])
	if err != nil {
		return shared.ParseResult{}, err
	}

	return shared.ParseResult{
		ParserType: "Translusent",
		Data: shared.CostData{
			Name: "Digital",
			Cost: intVar,
			Link: link,
		},
	}, nil
}
