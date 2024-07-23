package main

import (
	"log"
	parsers "printraduga_parser/parsers"
	"printraduga_parser/shared"
	"sync"
)

func main() {
	parsers := []shared.Parser{
		parsers.DigitalTranslusentParser{},
		parsers.CoralTranslusentParser{},
	}
	var wg sync.WaitGroup
	var resultMutex sync.Mutex
	var results []shared.CostInfo
	for _, parser := range parsers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Println("Parser in process: %v", parser)
			var result = parser.Parse()
			resultMutex.Lock()
			results = append(results, result)
			resultMutex.Unlock()
		}()
	}
	wg.Wait()

	for _, result := range results {
		log.Println(result)
	}
}
