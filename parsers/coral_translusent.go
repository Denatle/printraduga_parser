package parsers

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/chromedp/chromedp"

	shared "printraduga_parser/shared"
)

type CoralTranslusentParser struct {
}

func (p CoralTranslusentParser) Parse() shared.CostInfo {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	link := "https://www.coral-print.ru/pechat-nakleek/"

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(link),
		chromedp.WaitVisible(".wrap"),
		// chromedp.Sleep(time.Second*1),
		// chromedp.ScrollIntoView(".holst_list"),
		chromedp.Text(".vertical-align-middle > .price", &res, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}
	intVar, err := strconv.Atoi(strings.Replace(res, " руб", "", -1))
	if err != nil {
		log.Fatal(err)
	}

	return shared.CostInfo{
		Name:       "Coral translusent",
		Cost:       intVar,
		ParserType: 0,
		Link:       link,
	}
}
