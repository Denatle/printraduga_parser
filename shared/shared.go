package shared

import (
	"context"
	"github.com/chromedp/chromedp"
	"time"
)

type Parser interface {
	Parse() CostInfo
}

type CostInfo struct {
	Name       string
	ParserType int
	Cost       int
	Link       string
}

const (
	translusent = iota // 0
	holo        = iota // 1
)

func RunWithTimeOut(ctx *context.Context, timeout time.Duration, tasks chromedp.Tasks) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		timeoutContext, cancel := context.WithTimeout(ctx, timeout*time.Second)
		defer cancel()
		return tasks.Do(timeoutContext)
	}
}
