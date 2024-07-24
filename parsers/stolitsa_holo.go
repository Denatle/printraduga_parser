package parsers

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"strconv"
	"time"

	shared "printraduga_parser/shared"
)

type StolitsaHoloParser struct {
}

func (p StolitsaHoloParser) Parse() shared.ParseResult {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	defer cancel()

	link := "https://stolitsaprint.ru/pechat-nakleek/golograficheskie/"

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(link),
		// chromedp.Sleep(time.Second*2),
		chromedp.WaitVisible(".uCalc_276043"),
		chromedp.SendKeys("#input_text-4", "300"),
		chromedp.SetAttributeValue("#selector-16", "selectedIndex", "4"),
		// chromedp.Evaluate()
		chromedp.SetAttributeValue("#selector-17", "selectedIndex", "1"),
		chromedp.SetAttributeValue("#selector-18", "selectedIndex", "1"),
		chromedp.Sleep(time.Second*2),
		chromedp.Text(".js-result-sum-value", &res),
	)
	if err != nil {
		log.Fatal(err)
	}
	intVar, err := strconv.Atoi(res)
	if err != nil {
		log.Fatal(err)
	}

	return shared.ParseResult{
		ParserType: "Translusent",
		Data: shared.CostData{
			Name: "Stolitsa",
			Cost: intVar,
			Link: link,
		},
	}
}
