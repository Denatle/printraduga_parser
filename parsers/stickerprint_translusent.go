package parsers

import (
	"context"
	"strconv"
	"strings"

	"github.com/chromedp/chromedp"

	shared "printraduga_parser/shared"
)

type StickerPrintTranslusentParser struct {
}

func (p StickerPrintTranslusentParser) Parse() (shared.ParseResult, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", true))

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx) // chromedp.WithDebugf(log.Printf),

	defer cancel()
	// ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	// defer cancel()

	link := "https://stickerprint.ru/"
	// run task list
	var res string
	err := chromedp.Run(taskCtx,
		chromedp.EmulateViewport(1000, 1200),
		chromedp.Navigate(link),
		chromedp.WaitVisible("#price > div"),
		chromedp.ScrollIntoView("#forms_calc > div.dop_uslug > label:nth-child(1)"),
		chromedp.Click("#forms_calc > div.dop_uslug > label:nth-child(1)"),
		chromedp.Click("#forms_calc > div.dop_uslug > label:nth-child(3)"),
		chromedp.SetValue("#additional_option", "1280|1760|2240|2640|3120|3040|6000"),
		chromedp.SetValue("#quantity", "300"),
		chromedp.Text("#result > span", &res, chromedp.NodeVisible),
		// chromedp.Sleep(time.Hour),
	)
	if err != nil {
		return shared.ParseResult{}, err
	}
	trimmedString := strings.Replace(res, "\u00a0", "", -1)
	intVar, err := strconv.Atoi(trimmedString[:len(trimmedString)-6])
	if err != nil {
		return shared.ParseResult{}, err
	}

	return shared.ParseResult{
		ParserType: "Translusent",
		Data: shared.CostData{
			Name: "StickerPrint",
			Cost: intVar,
			Link: link,
		},
	}, nil
}
