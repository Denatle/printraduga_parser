package shared

import (
	"context"
	"github.com/chromedp/chromedp"
	"time"
)

type Parser interface {
	Parse() (ParseResult, error)
}

type ParseResult struct {
	ParserType string
	Data       CostData
}

type ExcelWriter interface {
	Write(filepath string, data map[string][]CostData) error
}

type TableData struct {
	TableName string
	Data      CostData
}

type CostData struct {
	Name string
	Cost int
	Link string
}

func RunWithTimeOut(ctx *context.Context, timeout time.Duration, tasks chromedp.Tasks) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		timeoutContext, cancel := context.WithTimeout(ctx, timeout*time.Second)
		defer cancel()
		return tasks.Do(timeoutContext)
	}
}
