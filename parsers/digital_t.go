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

	link := "https://digital-printing.ru/prints/stickers/prozrachnye#agree"
	// run task list
	var res string
	err := chromedp.Run(taskCtx,
		chromedp.EmulateViewport(1000, 1200),
		chromedp.Navigate(link),
		chromedp.WaitVisible("#pxpProducCalc > div.pxp-material-selector > div > div > ul > li.material-selector__item.material-types.option-row > span"),
		chromedp.Click("#pxpProducCalc > div.pxp-material-selector > div > div > ul > li:nth-child(2) > ul > li:nth-child(3) > a"),
		// chromedp.WaitEnabled(`//*[@id="pxpProducCalc"]/div[1]/div/div/ul/li[3]/ul/li[5]/a`),
		chromedp.Sleep(time.Millisecond*300),
		chromedp.Click(`//*[@id="pxpProducCalc"]/div[1]/div/div/ul/li[3]/ul/li[5]/a`),
		chromedp.Evaluate(`document.querySelector("#pxpProducCalc > div.pxp-material-selector > div > div > ul > li:nth-child(3) > ul > li:nth-child(5) > a").dispatchEvent(new Event("click"))`, nil),
		chromedp.Click(`//*[@id="pxpProducCalc"]/div[3]/div/div/ul/li[1]/ul/li[2]/span/label`),
		chromedp.SendKeys("#txtQuantity", kb.Backspace+kb.Backspace+"000"),
		chromedp.Sleep(time.Millisecond*700),
		chromedp.Text("#pxpProducCalc > div.pxp-total-price > div > div.totalPriceContainer > div > span", &res, chromedp.NodeVisible),
		// chromedp.Sleep(time.Hour),
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
