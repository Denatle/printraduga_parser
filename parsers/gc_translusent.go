package parsers

import (
	// "context"
	"github.com/chromedp/chromedp"
	"log"
	"time"

	cu "github.com/Davincible/chromedp-undetected"
	shared "printraduga_parser/shared"
)

type GcTranslusentParser struct {
}

func (p GcTranslusentParser) Parse() shared.CostInfo {
	// ctx, cancel := chromedp.NewContext(context.Background())
	ctx, cancel, err1 := cu.New(cu.NewConfig(
		// Remove this if you want to see a browser window.
		cu.WithHeadless(),

		// If the webelement is not found within 10 seconds, timeout.
		cu.WithTimeout(10*time.Second),
	))
	if err1 != nil {
		panic(err1)
	}
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://gcprint.ru/catalog/nakleyki/stikery-na-prozrachnoy-plenke-s-pechatyu-belym-tsvetom/`),
		chromedp.Sleep(time.Second*2),
		// chromedp.ScrollIntoView(".holst_list"),
		// shared.RunWithTimeOut(&ctx, 15, chromedp.Tasks{
		// chromedp.WaitVisible(".holst"),
		// chromedp.Text(".calc-tabs__item", &res, chromedp.NodeVisible),
		// chromedp.Click(".holst-calc > .holst-calc__item", chromedp.NodeVisible),
		// chromedp.Click("#typeRound", chromedp.NodeVisible),
		// chromedp.Click("#typeRectangle", chromedp.NodeVisible),
		chromedp.Click(".holst-calc__item-head", chromedp.NodeVisible),
		// }),
	)
	if err != nil {
		log.Fatal(err)
	}

	return shared.CostInfo{
		Name:       res,
		Cost:       10,
		ParserType: 0,
	}
}
