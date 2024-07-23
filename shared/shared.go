package shared

import (
	"context"
	"github.com/chromedp/chromedp"
	"time"
)

type Parser interface {
	Parse() (CostInfo, error)
}

type CostInfo struct {
	Name       string
	ParserType ParserType
	Cost       int
	Link       string
}
type ParserType int

const (
	Translusent ParserType = iota // 0
	Holo        ParserType = iota // 1
)

type ExcelWriter interface {
	Write(filepath string, data []CostInfo) error
}

func RunWithTimeOut(ctx *context.Context, timeout time.Duration, tasks chromedp.Tasks) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		timeoutContext, cancel := context.WithTimeout(ctx, timeout*time.Second)
		defer cancel()
		return tasks.Do(timeoutContext)
	}
}
